package server

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	secretRequired "github.com/danielheath/gin-teeny-security"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jcelliott/lumber"
)

const minutesToUnlock = 10.0

type Site struct {
	PathToData      string
	Css             []byte
	DefaultPage     string
	DefaultPassword string
	Debounce        int
	Diary           bool
	SessionStore    cookie.Store
	SecretCode      string
	AllowInsecure   bool
	Fileuploads     bool
	MaxUploadSize   uint
	Logger          *lumber.ConsoleLogger
	MaxDocumentSize uint // in runes; about a 10mb limit by default
	saveMut         sync.Mutex
}

func (s *Site) defaultLock() string {
	if s.DefaultPassword == "" {
		return ""
	}
	return HashPassword(s.DefaultPassword)
}

var hotTemplateReloading bool
var LogLevel int = lumber.WARN

func Serve(
	filepathToData,
	host,
	port,
	crt_path,
	key_path string,
	TLS bool,
	cssFile string,
	defaultPage string,
	defaultPassword string,
	debounce int,
	diary bool,
	secret string,
	secretCode string,
	allowInsecure bool,
	fileuploads bool,
	maxUploadSize uint,
	maxDocumentSize uint,
	logger *lumber.ConsoleLogger,
) {
	var customCSS []byte
	// collect custom CSS
	if len(cssFile) > 0 {
		var errRead error
		customCSS, errRead = ioutil.ReadFile(cssFile)
		if errRead != nil {
			fmt.Println(errRead)
			return
		}
		fmt.Printf("Loaded CSS file, %d bytes\n", len(customCSS))
	}

	router := Site{
		PathToData:      filepathToData,
		Css:             customCSS,
		DefaultPage:     defaultPage,
		DefaultPassword: defaultPassword,
		Debounce:        debounce,
		Diary:           diary,
		SessionStore:    cookie.NewStore([]byte(secret)),
		SecretCode:      secretCode,
		AllowInsecure:   allowInsecure,
		Fileuploads:     fileuploads,
		MaxUploadSize:   maxUploadSize,
		Logger:          logger,
		MaxDocumentSize: maxDocumentSize,
	}.Router()

	if TLS {
		http.ListenAndServeTLS(host+":"+port, crt_path, key_path, router)
	} else {
		panic(router.Run(host + ":" + port))
	}
}

func (s Site) Router() *gin.Engine {
	if s.Logger == nil {
		s.Logger = lumber.NewConsoleLogger(lumber.TRACE)
	}

	if hotTemplateReloading {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"sniffContentType": s.sniffContentType,
	})

	if hotTemplateReloading {
		router.LoadHTMLGlob("templates/*.tmpl")
	} else {
		router.HTMLRender = s.loadTemplates("index.tmpl")
	}

	router.Use(sessions.Sessions("_session", s.SessionStore))
	if s.SecretCode != "" {
		cfg := &secretRequired.Config{
			Secret: s.SecretCode,
			Path:   "/login/",
			RequireAuth: func(c *gin.Context) bool {
				page := c.Param("page")

				if page == "favicon.ico" || page == "static" || page == "uploads" {
					return false // no auth for these
				}

				return true
			},
		}
		router.Use(cfg.Middleware)
	}

	// router.Use(static.Serve("/static/", static.LocalFile("./static", true)))
	router.GET("/", func(c *gin.Context) {
		if s.DefaultPage != "" {
			c.Redirect(302, "/"+s.DefaultPage+"/read")
		} else {
			c.Redirect(302, "/"+randomAlliterateCombo())
		}
	})

	router.POST("/uploads", s.handleUpload)

	router.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")
		c.Redirect(302, "/"+page+"/view?"+c.Request.URL.RawQuery)
	})
	router.GET("/:page/*command", s.handlePageRequest)
	router.POST("/update", s.handlePageUpdate)
	router.POST("/relinquish", s.handlePageRelinquish) // relinquish returns the page no matter what (and destroys if nessecary)
	router.POST("/exists", s.handlePageExists)
	router.POST("/lock", s.handleLock)

	// Allow iframe/scripts in markup?
	allowInsecureHtml = s.AllowInsecure
	return router
}

func (s *Site) loadTemplates(list ...string) multitemplate.Render {
	r := multitemplate.New()

	for _, x := range list {
		templateString, err := Asset("templates/" + x)
		if err != nil {
			panic(err)
		}

		tmplMessage, err := template.New(x).Funcs(template.FuncMap{
			"sniffContentType": s.sniffContentType,
		}).Parse(string(templateString))
		if err != nil {
			panic(err)
		}

		r.Add(x, tmplMessage)
	}

	return r
}

func pageIsLocked(p *Page, c *gin.Context) bool {
	// it is easier to reason about when the page is actually unlocked
	var unlocked = !p.IsLocked ||
		(p.IsLocked && p.UnlockedFor == getSetSessionID(c))
	return !unlocked
}

func (s *Site) handlePageRelinquish(c *gin.Context) {
	type QueryJSON struct {
		Page string `json:"page"`
	}
	var json QueryJSON
	err := c.BindJSON(&json)
	if err != nil {
		s.Logger.Trace(err.Error())
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "Wrong JSON"})
		return
	}
	if len(json.Page) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "Must specify `page`"})
		return
	}
	message := "Relinquished"
	p := s.Open(json.Page)
	name := p.Meta
	if name == "" {
		name = json.Page
	}
	text := p.Text.GetCurrent()
	isLocked := pageIsLocked(p, c)
	if !isLocked {
		p.Erase()
		message = "Relinquished and erased"
	}
	c.JSON(http.StatusOK, gin.H{"success": true,
		"name":    name,
		"message": message,
		"text":    text,
		"locked":  isLocked,
	})
}

func getSetSessionID(c *gin.Context) (sid string) {
	var (
		session = sessions.Default(c)
		v       = session.Get("sid")
	)
	if v != nil {
		sid = v.(string)
	}
	if v == nil || sid == "" {
		sid = RandStringBytesMaskImprSrc(8)
		session.Set("sid", sid)
		session.Save()
	}
	return sid
}

func (s *Site) handlePageRequest(c *gin.Context) {
	page := c.Param("page")
	command := c.Param("command")

	if page == "favicon.ico" {
		data, _ := Asset("/static/img/cowyo/favicon.ico")
		c.Data(http.StatusOK, contentType("/static/img/cowyo/favicon.ico"), data)
		return
	} else if page == "static" {
		filename := page + command
		var data []byte
		if filename == "static/css/custom.css" {
			data = s.Css
		} else {
			var errAssset error
			data, errAssset = Asset(filename)
			if errAssset != nil {
				c.String(http.StatusInternalServerError, "Could not find data")
			}
		}
		c.Data(http.StatusOK, contentType(filename), data)
		return
	} else if page == "uploads" {
		if len(command) == 0 || command == "/" || command == "/edit" {
			if !s.Fileuploads {
				c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("uploads are disabled on this server"))
				return
			}
		} else {
			command = command[1:]
			if !strings.HasSuffix(command, ".upload") {
				command = command + ".upload"
			}
			pathname := path.Join(s.PathToData, command)

			if allowInsecureHtml {
				c.Header(
					"Content-Disposition",
					`inline; filename="`+c.DefaultQuery("filename", "upload")+`"`,
				)
			} else {
				// Prevent malicious html uploads by forcing type to plaintext and 'download-instead-of-view'
				c.Header("Content-Type", "text/plain")
				c.Header(
					"Content-Disposition",
					`attachment; filename="`+c.DefaultQuery("filename", "upload")+`"`,
				)
			}
			c.File(pathname)
			return
		}
	}

	if len(command) < 2 {
		c.Redirect(302, "/"+page+"/view")
		return
	}

	p := s.OpenOrInit(page, c.Request)

	// use the default lock
	if s.defaultLock() != "" && p.IsNew() {
		p.IsLocked = true
		p.PassphraseToUnlock = s.defaultLock()
	}

	//version := c.DefaultQuery("version", "ajksldfjl")
	isLocked := pageIsLocked(p, c)

	// Disallow anything but viewing locked pages
	if (isLocked) &&
		(command[0:2] != "/v" && command[0:2] != "/r") {
		c.Redirect(302, "/"+page+"/view")
		return
	}

	if command == "/erase" {
		if !isLocked {
			p.Erase()
			c.Redirect(302, "/"+page+"/edit")
		} else {
			c.Redirect(302, "/"+page+"/view")
		}
		return
	}
	rawText := p.Text.GetCurrent()
	rawHTML := p.RenderedPage

	// Check to see if an old version is requested
	// versionInt, versionErr := strconv.Atoi(version)
	// if versionErr == nil && versionInt > 0 {
	// 	versionText, err := p.Text.GetPreviousByTimestamp(int64(versionInt))
	// 	if err == nil {
	// 		rawText = versionText
	// 		rawHTML = GithubMarkdownToHTML(rawText)
	// 	}
	// }

	// Get history
	var versionsInt64 []int64
	var versionsChangeSums []int
	var versionsText []string
	if command[0:2] == "/h" {
		versionsInt64, versionsChangeSums = p.Text.GetMajorSnapshotsAndChangeSums(60) // get snapshots 60 seconds apart
		versionsText = make([]string, len(versionsInt64))
		for i, v := range versionsInt64 {
			versionsText[i] = time.Unix(v/1000000000, 0).Format("Mon Jan 2 15:04:05 MST 2006")
		}
		versionsText = reverseSliceString(versionsText)
		versionsInt64 = reverseSliceInt64(versionsInt64)
		versionsChangeSums = reverseSliceInt(versionsChangeSums)
	}

	if strings.HasPrefix(command, "/raw") {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Data(http.StatusOK, contentType(p.Identifier), []byte(rawText))
		return
	}

	if strings.HasPrefix(command, "/frontmatter") {
		c.Data(http.StatusOK, gin.MIMEJSON, p.FrontmatterJson)
		return
	}

	var DirectoryEntries []os.FileInfo
	if page == "ls" {
		command = "/view"
		DirectoryEntries = s.DirectoryList()
	}
	if page == "uploads" {
		command = "/view"
		var err error
		DirectoryEntries, err = s.UploadList()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"EditPage":    command[0:2] == "/e", // /edit
		"ViewPage":    command[0:2] == "/v", // /view
		"HistoryPage": command[0:2] == "/h", // /history
		"ReadPage":    command[0:2] == "/r", // /history
		"DontKnowPage": command[0:2] != "/e" &&
			command[0:2] != "/v" &&
			command[0:2] != "/l" &&
			command[0:2] != "/r" &&
			command[0:2] != "/h",
		"DirectoryPage":      page == "ls" || page == "uploads",
		"UploadPage":         page == "uploads",
		"DirectoryEntries":   DirectoryEntries,
		"Page":               page,
		"RenderedPage":       template.HTML(rawHTML),
		"RawPage":            rawText,
		"Versions":           versionsInt64,
		"VersionsText":       versionsText,
		"VersionsChangeSums": versionsChangeSums,
		"IsLocked":           isLocked,
		"Route":              "/" + page + command,
		"HasDotInName":       strings.Contains(page, "."),
		"RecentlyEdited":     getRecentlyEdited(page, c),
		"CustomCSS":          len(s.Css) > 0,
		"Debounce":           s.Debounce,
		"DiaryMode":          s.Diary,
		"Date":               time.Now().Format("2006-01-02"),
		"UnixTime":           time.Now().Unix(),
		"AllowFileUploads":   s.Fileuploads,
		"MaxUploadMB":        s.MaxUploadSize,
	})
}

func getRecentlyEdited(title string, c *gin.Context) []string {
	session := sessions.Default(c)
	var recentlyEdited string
	v := session.Get("recentlyEdited")
	editedThings := []string{}
	if v == nil {
		recentlyEdited = title
	} else {
		editedThings = strings.Split(v.(string), "|||")
		if !stringInSlice(title, editedThings) {
			recentlyEdited = v.(string) + "|||" + title
		} else {
			recentlyEdited = v.(string)
		}
	}
	session.Set("recentlyEdited", recentlyEdited)
	session.Save()
	editedThingsWithoutCurrent := make([]string, len(editedThings))
	i := 0
	for _, thing := range editedThings {
		if thing == title {
			continue
		}
		if strings.Contains(thing, "icon-") {
			continue
		}
		editedThingsWithoutCurrent[i] = thing
		i++
	}
	return editedThingsWithoutCurrent[:i]
}

func (s *Site) handlePageExists(c *gin.Context) {
	type QueryJSON struct {
		Page string `json:"page"`
	}
	var json QueryJSON
	err := c.BindJSON(&json)
	if err != nil {
		s.Logger.Trace(err.Error())
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "Wrong JSON", "exists": false})
		return
	}
	p := s.Open(json.Page)
	if len(p.Text.GetCurrent()) > 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": json.Page + " found", "exists": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": json.Page + " not found", "exists": false})
	}

}

func (s *Site) handlePageUpdate(c *gin.Context) {
	type QueryJSON struct {
		Page      string `json:"page"`
		NewText   string `json:"new_text"`
		FetchedAt int64  `json:"fetched_at"`
		Meta      string `json:"meta"`
	}
	var json QueryJSON
	err := c.BindJSON(&json)
	if err != nil {
		s.Logger.Trace(err.Error())
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "Wrong JSON"})
		return
	}
	if uint(len(json.NewText)) > s.MaxDocumentSize {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "Too much"})
		return
	}
	if len(json.Page) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "Must specify `page`"})
		return
	}
	s.Logger.Trace("Update: %v", json)
	p := s.Open(json.Page)
	var (
		message       string
		sinceLastEdit = time.Since(p.LastEditTime())
	)
	success := false
	if pageIsLocked(p, c) {
		if sinceLastEdit < minutesToUnlock {
			message = "This page is being edited by someone else"
		} else {
			// here what might have happened is that two people unlock without
			// editing thus they both suceeds but only one is able to edit
			message = "Locked, must unlock first"
		}
	} else if json.FetchedAt > 0 && p.LastEditUnixTime() > json.FetchedAt {
		message = "Refusing to overwrite others work"
	} else {
		p.Meta = json.Meta
		p.Update(json.NewText)
		p.Save()
		message = "Saved"
		success = true
	}
	c.JSON(http.StatusOK, gin.H{"success": success, "message": message, "unix_time": time.Now().Unix()})
}

func (s *Site) handleLock(c *gin.Context) {
	type QueryJSON struct {
		Page       string `json:"page"`
		Passphrase string `json:"passphrase"`
	}

	var json QueryJSON
	if c.BindJSON(&json) != nil {
		c.String(http.StatusBadRequest, "Problem binding keys")
		return
	}
	p := s.Open(json.Page)
	if s.defaultLock() != "" && p.IsNew() {
		p.IsLocked = true // IsLocked was replaced by variable wrt Context
		p.PassphraseToUnlock = s.defaultLock()
	}

	var (
		message       string
		sessionID     = getSetSessionID(c)
		sinceLastEdit = time.Since(p.LastEditTime())
	)

	// both lock/unlock ends here on locked&timeout combination
	if p.IsLocked &&
		p.UnlockedFor != sessionID &&
		p.UnlockedFor != "" &&
		sinceLastEdit.Minutes() < minutesToUnlock {

		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": fmt.Sprintf("This page is being edited by someone else! Will unlock automatically %2.0f minutes after the last change.", minutesToUnlock-sinceLastEdit.Minutes()),
		})
		return
	}
	if !pageIsLocked(p, c) {
		p.IsLocked = true
		p.PassphraseToUnlock = HashPassword(json.Passphrase)
		p.UnlockedFor = ""
		message = "Locked"
	} else {
		err2 := CheckPasswordHash(json.Passphrase, p.PassphraseToUnlock)
		if err2 != nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "Can't unlock"})
			return
		}
		p.UnlockedFor = sessionID
		message = "Unlocked only for you"
	}
	p.Save()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": message})
}

func (s *Site) handleUpload(c *gin.Context) {
	if !s.Fileuploads {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("uploads are disabled on this server"))
		return
	}

	file, info, err := c.Request.FormFile("file")
	defer file.Close()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	newName := "sha256-" + encodeBytesToBase32(h.Sum(nil))

	// Replaces any existing version, but sha256 collisions are rare as anything.
	outfile, err := os.Create(path.Join(s.PathToData, newName+".upload"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	file.Seek(0, io.SeekStart)
	_, err = io.Copy(outfile, file)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Location", "/uploads/"+newName+"?filename="+url.QueryEscape(info.Filename))
	return
}

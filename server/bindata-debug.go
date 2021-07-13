// Code generated for package server by go-bindata DO NOT EDIT. (@generated)
// sources:
// static/css/base-min.css
// static/css/default.css
// static/css/dropzone.css
// static/css/github-markdown.css
// static/css/highlight.css
// static/css/menus-min.css
// static/img/cowyo/android-icon-144x144.png
// static/img/cowyo/android-icon-192x192.png
// static/img/cowyo/android-icon-36x36.png
// static/img/cowyo/android-icon-48x48.png
// static/img/cowyo/android-icon-72x72.png
// static/img/cowyo/android-icon-96x96.png
// static/img/cowyo/apple-icon-114x114.png
// static/img/cowyo/apple-icon-120x120.png
// static/img/cowyo/apple-icon-144x144.png
// static/img/cowyo/apple-icon-152x152.png
// static/img/cowyo/apple-icon-180x180.png
// static/img/cowyo/apple-icon-57x57.png
// static/img/cowyo/apple-icon-60x60.png
// static/img/cowyo/apple-icon-72x72.png
// static/img/cowyo/apple-icon-76x76.png
// static/img/cowyo/apple-icon-precomposed.png
// static/img/cowyo/apple-icon.png
// static/img/cowyo/browserconfig.xml
// static/img/cowyo/favicon-16x16.png
// static/img/cowyo/favicon-32x32.png
// static/img/cowyo/favicon-96x96.png
// static/img/cowyo/favicon.ico
// static/img/cowyo/manifest.json
// static/img/cowyo/ms-icon-144x144.png
// static/img/cowyo/ms-icon-150x150.png
// static/img/cowyo/ms-icon-310x310.png
// static/img/cowyo/ms-icon-70x70.png
// static/img/logo.png
// static/js/cowyo.js
// static/js/dropzone.js
// static/js/highlight.min.js
// static/js/highlight.pack.js
// static/js/jquery-1.8.3.js
// static/text/adjectives
// static/text/adjectives.old
// static/text/animals
// static/text/animals.all
// static/text/howmany.py
// static/text/robots.txt
// static/text/sitemap.xml
// templates/index.tmpl
// +build debug

package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// staticCssBaseMinCss reads file data from disk. It returns an error on failure.
func staticCssBaseMinCss() (*asset, error) {
	path := "/workspaces/cowyo/static/css/base-min.css"
	name := "static/css/base-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticCssDefaultCss reads file data from disk. It returns an error on failure.
func staticCssDefaultCss() (*asset, error) {
	path := "/workspaces/cowyo/static/css/default.css"
	name := "static/css/default.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticCssDropzoneCss reads file data from disk. It returns an error on failure.
func staticCssDropzoneCss() (*asset, error) {
	path := "/workspaces/cowyo/static/css/dropzone.css"
	name := "static/css/dropzone.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticCssGithubMarkdownCss reads file data from disk. It returns an error on failure.
func staticCssGithubMarkdownCss() (*asset, error) {
	path := "/workspaces/cowyo/static/css/github-markdown.css"
	name := "static/css/github-markdown.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticCssHighlightCss reads file data from disk. It returns an error on failure.
func staticCssHighlightCss() (*asset, error) {
	path := "/workspaces/cowyo/static/css/highlight.css"
	name := "static/css/highlight.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticCssMenusMinCss reads file data from disk. It returns an error on failure.
func staticCssMenusMinCss() (*asset, error) {
	path := "/workspaces/cowyo/static/css/menus-min.css"
	name := "static/css/menus-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAndroidIcon144x144Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAndroidIcon144x144Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/android-icon-144x144.png"
	name := "static/img/cowyo/android-icon-144x144.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAndroidIcon192x192Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAndroidIcon192x192Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/android-icon-192x192.png"
	name := "static/img/cowyo/android-icon-192x192.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAndroidIcon36x36Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAndroidIcon36x36Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/android-icon-36x36.png"
	name := "static/img/cowyo/android-icon-36x36.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAndroidIcon48x48Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAndroidIcon48x48Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/android-icon-48x48.png"
	name := "static/img/cowyo/android-icon-48x48.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAndroidIcon72x72Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAndroidIcon72x72Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/android-icon-72x72.png"
	name := "static/img/cowyo/android-icon-72x72.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAndroidIcon96x96Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAndroidIcon96x96Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/android-icon-96x96.png"
	name := "static/img/cowyo/android-icon-96x96.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon114x114Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon114x114Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-114x114.png"
	name := "static/img/cowyo/apple-icon-114x114.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon120x120Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon120x120Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-120x120.png"
	name := "static/img/cowyo/apple-icon-120x120.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon144x144Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon144x144Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-144x144.png"
	name := "static/img/cowyo/apple-icon-144x144.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon152x152Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon152x152Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-152x152.png"
	name := "static/img/cowyo/apple-icon-152x152.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon180x180Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon180x180Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-180x180.png"
	name := "static/img/cowyo/apple-icon-180x180.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon57x57Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon57x57Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-57x57.png"
	name := "static/img/cowyo/apple-icon-57x57.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon60x60Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon60x60Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-60x60.png"
	name := "static/img/cowyo/apple-icon-60x60.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon72x72Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon72x72Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-72x72.png"
	name := "static/img/cowyo/apple-icon-72x72.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIcon76x76Png reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIcon76x76Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-76x76.png"
	name := "static/img/cowyo/apple-icon-76x76.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIconPrecomposedPng reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIconPrecomposedPng() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon-precomposed.png"
	name := "static/img/cowyo/apple-icon-precomposed.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoAppleIconPng reads file data from disk. It returns an error on failure.
func staticImgCowyoAppleIconPng() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/apple-icon.png"
	name := "static/img/cowyo/apple-icon.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoBrowserconfigXml reads file data from disk. It returns an error on failure.
func staticImgCowyoBrowserconfigXml() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/browserconfig.xml"
	name := "static/img/cowyo/browserconfig.xml"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoFavicon16x16Png reads file data from disk. It returns an error on failure.
func staticImgCowyoFavicon16x16Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/favicon-16x16.png"
	name := "static/img/cowyo/favicon-16x16.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoFavicon32x32Png reads file data from disk. It returns an error on failure.
func staticImgCowyoFavicon32x32Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/favicon-32x32.png"
	name := "static/img/cowyo/favicon-32x32.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoFavicon96x96Png reads file data from disk. It returns an error on failure.
func staticImgCowyoFavicon96x96Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/favicon-96x96.png"
	name := "static/img/cowyo/favicon-96x96.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoFaviconIco reads file data from disk. It returns an error on failure.
func staticImgCowyoFaviconIco() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/favicon.ico"
	name := "static/img/cowyo/favicon.ico"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoManifestJson reads file data from disk. It returns an error on failure.
func staticImgCowyoManifestJson() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/manifest.json"
	name := "static/img/cowyo/manifest.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoMsIcon144x144Png reads file data from disk. It returns an error on failure.
func staticImgCowyoMsIcon144x144Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/ms-icon-144x144.png"
	name := "static/img/cowyo/ms-icon-144x144.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoMsIcon150x150Png reads file data from disk. It returns an error on failure.
func staticImgCowyoMsIcon150x150Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/ms-icon-150x150.png"
	name := "static/img/cowyo/ms-icon-150x150.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoMsIcon310x310Png reads file data from disk. It returns an error on failure.
func staticImgCowyoMsIcon310x310Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/ms-icon-310x310.png"
	name := "static/img/cowyo/ms-icon-310x310.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgCowyoMsIcon70x70Png reads file data from disk. It returns an error on failure.
func staticImgCowyoMsIcon70x70Png() (*asset, error) {
	path := "/workspaces/cowyo/static/img/cowyo/ms-icon-70x70.png"
	name := "static/img/cowyo/ms-icon-70x70.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticImgLogoPng reads file data from disk. It returns an error on failure.
func staticImgLogoPng() (*asset, error) {
	path := "/workspaces/cowyo/static/img/logo.png"
	name := "static/img/logo.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsCowyoJs reads file data from disk. It returns an error on failure.
func staticJsCowyoJs() (*asset, error) {
	path := "/workspaces/cowyo/static/js/cowyo.js"
	name := "static/js/cowyo.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsDropzoneJs reads file data from disk. It returns an error on failure.
func staticJsDropzoneJs() (*asset, error) {
	path := "/workspaces/cowyo/static/js/dropzone.js"
	name := "static/js/dropzone.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsHighlightMinJs reads file data from disk. It returns an error on failure.
func staticJsHighlightMinJs() (*asset, error) {
	path := "/workspaces/cowyo/static/js/highlight.min.js"
	name := "static/js/highlight.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsHighlightPackJs reads file data from disk. It returns an error on failure.
func staticJsHighlightPackJs() (*asset, error) {
	path := "/workspaces/cowyo/static/js/highlight.pack.js"
	name := "static/js/highlight.pack.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticJsJquery183Js reads file data from disk. It returns an error on failure.
func staticJsJquery183Js() (*asset, error) {
	path := "/workspaces/cowyo/static/js/jquery-1.8.3.js"
	name := "static/js/jquery-1.8.3.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTextAdjectives reads file data from disk. It returns an error on failure.
func staticTextAdjectives() (*asset, error) {
	path := "/workspaces/cowyo/static/text/adjectives"
	name := "static/text/adjectives"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTextAdjectivesOld reads file data from disk. It returns an error on failure.
func staticTextAdjectivesOld() (*asset, error) {
	path := "/workspaces/cowyo/static/text/adjectives.old"
	name := "static/text/adjectives.old"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTextAnimals reads file data from disk. It returns an error on failure.
func staticTextAnimals() (*asset, error) {
	path := "/workspaces/cowyo/static/text/animals"
	name := "static/text/animals"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTextAnimalsAll reads file data from disk. It returns an error on failure.
func staticTextAnimalsAll() (*asset, error) {
	path := "/workspaces/cowyo/static/text/animals.all"
	name := "static/text/animals.all"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTextHowmanyPy reads file data from disk. It returns an error on failure.
func staticTextHowmanyPy() (*asset, error) {
	path := "/workspaces/cowyo/static/text/howmany.py"
	name := "static/text/howmany.py"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTextRobotsTxt reads file data from disk. It returns an error on failure.
func staticTextRobotsTxt() (*asset, error) {
	path := "/workspaces/cowyo/static/text/robots.txt"
	name := "static/text/robots.txt"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticTextSitemapXml reads file data from disk. It returns an error on failure.
func staticTextSitemapXml() (*asset, error) {
	path := "/workspaces/cowyo/static/text/sitemap.xml"
	name := "static/text/sitemap.xml"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesIndexTmpl reads file data from disk. It returns an error on failure.
func templatesIndexTmpl() (*asset, error) {
	path := "/workspaces/cowyo/templates/index.tmpl"
	name := "templates/index.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/css/base-min.css":                     staticCssBaseMinCss,
	"static/css/default.css":                      staticCssDefaultCss,
	"static/css/dropzone.css":                     staticCssDropzoneCss,
	"static/css/github-markdown.css":              staticCssGithubMarkdownCss,
	"static/css/highlight.css":                    staticCssHighlightCss,
	"static/css/menus-min.css":                    staticCssMenusMinCss,
	"static/img/cowyo/android-icon-144x144.png":   staticImgCowyoAndroidIcon144x144Png,
	"static/img/cowyo/android-icon-192x192.png":   staticImgCowyoAndroidIcon192x192Png,
	"static/img/cowyo/android-icon-36x36.png":     staticImgCowyoAndroidIcon36x36Png,
	"static/img/cowyo/android-icon-48x48.png":     staticImgCowyoAndroidIcon48x48Png,
	"static/img/cowyo/android-icon-72x72.png":     staticImgCowyoAndroidIcon72x72Png,
	"static/img/cowyo/android-icon-96x96.png":     staticImgCowyoAndroidIcon96x96Png,
	"static/img/cowyo/apple-icon-114x114.png":     staticImgCowyoAppleIcon114x114Png,
	"static/img/cowyo/apple-icon-120x120.png":     staticImgCowyoAppleIcon120x120Png,
	"static/img/cowyo/apple-icon-144x144.png":     staticImgCowyoAppleIcon144x144Png,
	"static/img/cowyo/apple-icon-152x152.png":     staticImgCowyoAppleIcon152x152Png,
	"static/img/cowyo/apple-icon-180x180.png":     staticImgCowyoAppleIcon180x180Png,
	"static/img/cowyo/apple-icon-57x57.png":       staticImgCowyoAppleIcon57x57Png,
	"static/img/cowyo/apple-icon-60x60.png":       staticImgCowyoAppleIcon60x60Png,
	"static/img/cowyo/apple-icon-72x72.png":       staticImgCowyoAppleIcon72x72Png,
	"static/img/cowyo/apple-icon-76x76.png":       staticImgCowyoAppleIcon76x76Png,
	"static/img/cowyo/apple-icon-precomposed.png": staticImgCowyoAppleIconPrecomposedPng,
	"static/img/cowyo/apple-icon.png":             staticImgCowyoAppleIconPng,
	"static/img/cowyo/browserconfig.xml":          staticImgCowyoBrowserconfigXml,
	"static/img/cowyo/favicon-16x16.png":          staticImgCowyoFavicon16x16Png,
	"static/img/cowyo/favicon-32x32.png":          staticImgCowyoFavicon32x32Png,
	"static/img/cowyo/favicon-96x96.png":          staticImgCowyoFavicon96x96Png,
	"static/img/cowyo/favicon.ico":                staticImgCowyoFaviconIco,
	"static/img/cowyo/manifest.json":              staticImgCowyoManifestJson,
	"static/img/cowyo/ms-icon-144x144.png":        staticImgCowyoMsIcon144x144Png,
	"static/img/cowyo/ms-icon-150x150.png":        staticImgCowyoMsIcon150x150Png,
	"static/img/cowyo/ms-icon-310x310.png":        staticImgCowyoMsIcon310x310Png,
	"static/img/cowyo/ms-icon-70x70.png":          staticImgCowyoMsIcon70x70Png,
	"static/img/logo.png":                         staticImgLogoPng,
	"static/js/cowyo.js":                          staticJsCowyoJs,
	"static/js/dropzone.js":                       staticJsDropzoneJs,
	"static/js/highlight.min.js":                  staticJsHighlightMinJs,
	"static/js/highlight.pack.js":                 staticJsHighlightPackJs,
	"static/js/jquery-1.8.3.js":                   staticJsJquery183Js,
	"static/text/adjectives":                      staticTextAdjectives,
	"static/text/adjectives.old":                  staticTextAdjectivesOld,
	"static/text/animals":                         staticTextAnimals,
	"static/text/animals.all":                     staticTextAnimalsAll,
	"static/text/howmany.py":                      staticTextHowmanyPy,
	"static/text/robots.txt":                      staticTextRobotsTxt,
	"static/text/sitemap.xml":                     staticTextSitemapXml,
	"templates/index.tmpl":                        templatesIndexTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"static": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"base-min.css":        &bintree{staticCssBaseMinCss, map[string]*bintree{}},
			"default.css":         &bintree{staticCssDefaultCss, map[string]*bintree{}},
			"dropzone.css":        &bintree{staticCssDropzoneCss, map[string]*bintree{}},
			"github-markdown.css": &bintree{staticCssGithubMarkdownCss, map[string]*bintree{}},
			"highlight.css":       &bintree{staticCssHighlightCss, map[string]*bintree{}},
			"menus-min.css":       &bintree{staticCssMenusMinCss, map[string]*bintree{}},
		}},
		"img": &bintree{nil, map[string]*bintree{
			"cowyo": &bintree{nil, map[string]*bintree{
				"android-icon-144x144.png":   &bintree{staticImgCowyoAndroidIcon144x144Png, map[string]*bintree{}},
				"android-icon-192x192.png":   &bintree{staticImgCowyoAndroidIcon192x192Png, map[string]*bintree{}},
				"android-icon-36x36.png":     &bintree{staticImgCowyoAndroidIcon36x36Png, map[string]*bintree{}},
				"android-icon-48x48.png":     &bintree{staticImgCowyoAndroidIcon48x48Png, map[string]*bintree{}},
				"android-icon-72x72.png":     &bintree{staticImgCowyoAndroidIcon72x72Png, map[string]*bintree{}},
				"android-icon-96x96.png":     &bintree{staticImgCowyoAndroidIcon96x96Png, map[string]*bintree{}},
				"apple-icon-114x114.png":     &bintree{staticImgCowyoAppleIcon114x114Png, map[string]*bintree{}},
				"apple-icon-120x120.png":     &bintree{staticImgCowyoAppleIcon120x120Png, map[string]*bintree{}},
				"apple-icon-144x144.png":     &bintree{staticImgCowyoAppleIcon144x144Png, map[string]*bintree{}},
				"apple-icon-152x152.png":     &bintree{staticImgCowyoAppleIcon152x152Png, map[string]*bintree{}},
				"apple-icon-180x180.png":     &bintree{staticImgCowyoAppleIcon180x180Png, map[string]*bintree{}},
				"apple-icon-57x57.png":       &bintree{staticImgCowyoAppleIcon57x57Png, map[string]*bintree{}},
				"apple-icon-60x60.png":       &bintree{staticImgCowyoAppleIcon60x60Png, map[string]*bintree{}},
				"apple-icon-72x72.png":       &bintree{staticImgCowyoAppleIcon72x72Png, map[string]*bintree{}},
				"apple-icon-76x76.png":       &bintree{staticImgCowyoAppleIcon76x76Png, map[string]*bintree{}},
				"apple-icon-precomposed.png": &bintree{staticImgCowyoAppleIconPrecomposedPng, map[string]*bintree{}},
				"apple-icon.png":             &bintree{staticImgCowyoAppleIconPng, map[string]*bintree{}},
				"browserconfig.xml":          &bintree{staticImgCowyoBrowserconfigXml, map[string]*bintree{}},
				"favicon-16x16.png":          &bintree{staticImgCowyoFavicon16x16Png, map[string]*bintree{}},
				"favicon-32x32.png":          &bintree{staticImgCowyoFavicon32x32Png, map[string]*bintree{}},
				"favicon-96x96.png":          &bintree{staticImgCowyoFavicon96x96Png, map[string]*bintree{}},
				"favicon.ico":                &bintree{staticImgCowyoFaviconIco, map[string]*bintree{}},
				"manifest.json":              &bintree{staticImgCowyoManifestJson, map[string]*bintree{}},
				"ms-icon-144x144.png":        &bintree{staticImgCowyoMsIcon144x144Png, map[string]*bintree{}},
				"ms-icon-150x150.png":        &bintree{staticImgCowyoMsIcon150x150Png, map[string]*bintree{}},
				"ms-icon-310x310.png":        &bintree{staticImgCowyoMsIcon310x310Png, map[string]*bintree{}},
				"ms-icon-70x70.png":          &bintree{staticImgCowyoMsIcon70x70Png, map[string]*bintree{}},
			}},
			"logo.png": &bintree{staticImgLogoPng, map[string]*bintree{}},
		}},
		"js": &bintree{nil, map[string]*bintree{
			"cowyo.js":          &bintree{staticJsCowyoJs, map[string]*bintree{}},
			"dropzone.js":       &bintree{staticJsDropzoneJs, map[string]*bintree{}},
			"highlight.min.js":  &bintree{staticJsHighlightMinJs, map[string]*bintree{}},
			"highlight.pack.js": &bintree{staticJsHighlightPackJs, map[string]*bintree{}},
			"jquery-1.8.3.js":   &bintree{staticJsJquery183Js, map[string]*bintree{}},
		}},
		"text": &bintree{nil, map[string]*bintree{
			"adjectives":     &bintree{staticTextAdjectives, map[string]*bintree{}},
			"adjectives.old": &bintree{staticTextAdjectivesOld, map[string]*bintree{}},
			"animals":        &bintree{staticTextAnimals, map[string]*bintree{}},
			"animals.all":    &bintree{staticTextAnimalsAll, map[string]*bintree{}},
			"howmany.py":     &bintree{staticTextHowmanyPy, map[string]*bintree{}},
			"robots.txt":     &bintree{staticTextRobotsTxt, map[string]*bintree{}},
			"sitemap.xml":    &bintree{staticTextSitemapXml, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"index.tmpl": &bintree{templatesIndexTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

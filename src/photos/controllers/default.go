package controllers

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

// Photo struct represents a pohto
type Photo struct {
	Origin string
	Small  string
	Title  string
}

// MainController main controller
type MainController struct {
	beego.Controller
}

// Get return main page
func (c *MainController) Get() {
	var photos []Photo
	c.Data["Photos"] = photos
	c.Data["Title"] = "Our Photos"
	c.TplName = "index.tpl"
}

// GetImage return image
func (c *MainController) GetImage() {
	path := c.Ctx.Input.Param(":path")
	rawPath, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		return
	}
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, string(rawPath))
}

// GetPage return a specified page
func (c *MainController) GetPage() {
	path := c.Ctx.Input.Param(":path")
	rawPath, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		return
	}
	var photos []Photo
	var dirs []string
	walkFunc := func(itemPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirs = append(dirs, itemPath)
			return nil
		}
		photos = append(photos, Photo{
			Origin: "/i/" + base64.StdEncoding.EncodeToString([]byte(path)),
			Small:  "/i/" + base64.StdEncoding.EncodeToString([]byte(path)),
			Title:  info.Name(),
		})
		return nil
	}
	filepath.Walk(string(rawPath), walkFunc)
	c.Data["Dirs"] = dirs
	c.Data["Photos"] = photos
	c.Data["Title"] = path
	c.TplName = "index.tpl"
}
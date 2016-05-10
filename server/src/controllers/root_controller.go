package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

type assets struct {
	Stylesheet string
	Javascript string
}

var assetsData = assetsPath()

func assetsPath() *assets {
	a := &assets{}

	if appengine.IsDevAppServer() {
		a.Javascript = "http://localhost:3500/assets/main.js"
		a.Stylesheet = "http://localhost:3500/assets/main.css"
	} else {
		buf, err := ioutil.ReadFile("public/assets/webpack-manifest.json")
		if err != nil {
			log.Fatalf("failed to load webpack manifest")
		}
		m, err := simplejson.NewJson(buf)
		if err != nil {
			log.Fatalf("failed to load webpack manifest json")
		}
		a.Javascript = "/assets/" + m.Get("main.js").MustString()
		a.Stylesheet = "/assets/" + m.Get("main.css").MustString()
	}

	return a
}

func GetMainView(c echo.Context) error {
	data := map[string]interface{}{}
	data["AssetData"] = assetsData
	return c.Render(http.StatusOK, "layout", data)
}

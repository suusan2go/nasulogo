package nasulogo

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"controllers"
	"middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initEcho() *echo.Echo {
	e := echo.New()

	s := standard.New("")
	s.SetHandler(e)
	e.Use(middleware.LoadDotEnv())
	http.Handle("/", s)

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.SetRenderer(t)

	return e
}

func init() {
	var router = initEcho()
	router.GET("/", controllers.GetMainView)
	router.GET("/auth/login/:provider", controllers.GetAuth)
	router.GET("/auth/callback/:provider", controllers.CallbackAuth)
}

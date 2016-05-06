package nasulogo

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"google.golang.org/appengine"

	"controllers"
)

func initEcho() *echo.Echo {
	e := echo.New()

	s := standard.New("")
	s.SetHandler(e)
	e.Use(loadDotEnv())
	http.Handle("/", s)

	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func loadDotEnv() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if appengine.IsDevAppServer() {
				err := godotenv.Load()
				if err != nil {
					log.Fatal("Error loading .env file")
				}
			}
			return next(c)
		}
	}
}

func init() {

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	var router = initEcho()
	router.SetRenderer(t)
	router.GET("/", func(c echo.Context) error { return c.Render(http.StatusOK, "layout", nil) })
	router.GET("/auth/login/:provider", controllers.GetAuth)
	router.GET("/auth/callback/:provider", controllers.CallbackAuth)
}

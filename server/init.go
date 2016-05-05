package nasulogo

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func initEcho() *echo.Echo {
	e := echo.New()

	s := standard.New("")
	s.SetHandler(e)
	http.Handle("/", s)

	return e
}

func init() {
	var router = initEcho()
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!\n")
	})
}

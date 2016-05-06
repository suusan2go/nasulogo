package middleware

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"log"
)

func LoadDotEnv() echo.MiddlewareFunc {
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

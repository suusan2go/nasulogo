package models

import (
	"net/http"

	"github.com/mjibson/goon"
)

func makeGoon(r *http.Request) *goon.Goon {
	g := goon.NewGoon(r)
	return g
}

package models

import (
	"net/http"

	"github.com/mjibson/goon"
)

type User struct {
	Id        string `datastore:"-" goon:"id"`
	Name      string `datastore:name"`
	Email     string `datastore:email"`
	AvatarUrl string `datasotre:avatarUrl,noindex`
}

func GetUser(r *http.Request, u *User) error {
	g := goon.NewGoon(r)
	if err := g.Get(u); err != nil {
		return err
	}
	return nil
}

func CreateUser(r *http.Request, u *User) error {
	g := goon.NewGoon(r)
	if _, err := g.Put(u); err != nil {
		return err
	}
	return nil
}

package models

import (
	"net/http"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

type User struct {
	Id        int64  `datastore:"-" goon:"id"`
	Name      string `datastore:name"`
	Email     string `datastore:email"`
	Uid       string `datasotre:uId`
	AvatarUrl string `datasotre:avatarUrl,noindex`
}

func GetOrCreateUser(r *http.Request, u *User) (*User, error) {
	user, err := GetUser(r, &User{Uid: u.Uid})
	return user, err
}

func GetUser(r *http.Request, u *User) (*User, error) {
	g := goon.NewGoon(r)
	if err := g.Get(u); err != nil {
		if err != datastore.ErrNoSuchEntity {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return u, nil
}

func CreateUser(r *http.Request, u *User) (*User, error) {
	g := goon.NewGoon(r)
	if _, err := g.Put(u); err != nil {
		return u, err
	}
	return u, nil
}

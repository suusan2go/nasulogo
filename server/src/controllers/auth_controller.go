package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"github.com/stretchr/signature"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	"models"
)

// 本当はinitで処理したいが、GAEではinitだと環境変数を読み出せない
// またapp engineではhttp.Transportなどが使用できないため別のメソッドに差し替える必要がある
// SEE: https://github.com/stretchr/gomniauth/pull/23
func initGomniauth(ctx *context.Context) {
	hostname := appengine.DefaultVersionHostname(*ctx)
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://"+hostname+"/auth/callback/google",
		),
	)
	t := new(urlfetch.Transport)
	t.Context = *ctx
	common.SetRoundTripper(t)
}

func GetAuth(c echo.Context) error {
	ctx := appengine.NewContext(c.Request().(*standard.Request).Request)
	initGomniauth(&ctx)

	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		log.Errorf(ctx, "認証プロバイダーの取得に失敗しました", provider, "-", err)
		return err
	}
	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Errorf(ctx, "GetBiginAuthURLの呼び出し中にエラーが発生しました", provider, "-", err)
		return err
	}
	return c.Redirect(http.StatusTemporaryRedirect, loginUrl)
}

func CallbackAuth(c echo.Context) error {
	r := c.Request().(*standard.Request).Request
	ctx := appengine.NewContext(r)

	initGomniauth(&ctx)

	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		log.Errorf(ctx, "認証プロバイダーの取得に失敗しました", provider, "-", err)
		return err
	}
	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
	if err != nil {
		log.Errorf(ctx, "認証を完了できませんんでした", provider, "-", err)
		return err
	}
	pu, err := provider.GetUser(creds)
	if err != nil {
		log.Errorf(ctx, "ユーザーの取得に失敗しました", provider, "-", err)
		return err
	}

	user := &models.User{Id: pu.IDForProvider(c.Param("provider"))}
	if err := models.GetUser(r, user); err == datastore.ErrNoSuchEntity {
		user = &models.User{
			Id:        pu.IDForProvider(c.Param("provider")),
			Name:      pu.Name(),
			Email:     pu.Email(),
			AvatarUrl: pu.AvatarURL(),
		}
		if err := models.CreateUser(r, user); err != nil {
			log.Errorf(ctx, "ユーザーの作成に失敗しました", provider, "-", err)
			return err
		}
	} else if err != nil {
		log.Errorf(ctx, "ユーザーの取得に失敗しました", provider, "-", err)
		return err
	}
	log.Debugf(ctx, user.Id, user.Name)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

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
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

// 本当はinitで処理したいが、GAEではinitだと環境変数を読み出せない
// またapp engineではhttp.Transportなどが使用できないため別のメソッドに差し替える必要がある
// SEE: https://github.com/stretchr/gomniauth/pull/23
func initGomniauth(ctx *context.Context) {
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://localhost:8080/auth/callback/google",
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
	}
	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Errorf(ctx, "GetBiginAuthURLの呼び出し中にエラーが発生しました", provider, "-", err)
	}
	return c.Redirect(http.StatusTemporaryRedirect, loginUrl)
}

func CallbackAuth(c echo.Context) error {
	request := c.Request().(*standard.Request).Request
	ctx := appengine.NewContext(request)

	initGomniauth(&ctx)

	provider, err := gomniauth.Provider(c.Param("provider"))
	if err != nil {
		log.Errorf(ctx, "認証プロバイダーの取得に失敗しました", provider, "-", err)
	}
	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(request.URL.RawQuery))
	if err != nil {
		log.Errorf(ctx, "認証を完了できませんんでした", provider, "-", err)
	}
	user, err := provider.GetUser(creds)
	if err != nil {
		log.Errorf(ctx, "ユーザーの取得に失敗しました", provider, "-", err)
	}
	log.Debugf(ctx, user.Name())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

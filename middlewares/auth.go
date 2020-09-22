package middlewares

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func verifyFirebaseIDToken(ctx echo.Context, auth *auth.Client) (*auth.Token, error) {
	headerAuth := ctx.Request().Header.Get("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)
	logrus.Debug(err)

	return jwtToken, err
}

// FirebaseGurad ログインしている場合にのみ使えるAPIに対して認証を行う
func FirebaseGurad() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, err := verifyFirebaseIDToken(c, authClient)

			if err != nil {
				return c.JSON(fasthttp.StatusUnauthorized, "Not Authenticated")
			}

			c.Set("auth", jwtToken)

			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}

// FirebaseAuth ログインしていなくとも利用できるAPIに対して認証を行う
func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, _ := verifyFirebaseIDToken(c, authClient)

			c.Set("auth", jwtToken)

			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}

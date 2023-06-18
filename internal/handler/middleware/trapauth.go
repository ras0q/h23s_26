package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_26/internal/pkg/config"
	"golang.org/x/oauth2"
)

func TrapAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get(config.SessionName, c)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
			}

			tok, ok := sess.Values[config.TokenKey].(*oauth2.Token)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			c.Set(string(config.TokenKey), tok)

			userID, ok := sess.Values[config.TraqIDKey].(string)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			c.Set(string(config.TraqIDKey), userID)

			return next(c)
		}
	}
}

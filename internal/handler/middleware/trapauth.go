package middleware

import "github.com/labstack/echo/v4"

func TRAPAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO: https://wiki.trap.jp/SysAd/docs/NeoShowcase#head16 の部員認証を実装する
			c.Set("userID", "user1")

			return next(c)
		}
	}
}

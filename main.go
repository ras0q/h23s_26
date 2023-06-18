package main

import (
	"encoding/gob"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/traP-jp/h23s_26/internal/handler"
	"github.com/traP-jp/h23s_26/internal/migration"
	"github.com/traP-jp/h23s_26/internal/pkg/config"
	"github.com/traP-jp/h23s_26/internal/repository"
	"golang.org/x/oauth2"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// register types for session
	gob.Register(config.SessionKey(""))
	gob.Register(&oauth2.Token{})

	e := echo.New()

	// middlewares
	corsConfig := middleware.DefaultCORSConfig
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, http.MethodOptions)
	e.Use(middleware.CORSWithConfig(corsConfig))

	proxyConfig := middleware.DefaultProxyConfig
	proxyConfig.Balancer = middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: config.ClientURL(),
		},
	})

	proxyConfig.Skipper = func(c echo.Context) bool {
		if strings.HasPrefix(c.Path(), "/api/v1") {
			return true
		}
		c.Request().Host = config.ClientURL().Host
		return false
	}
	proxyConfig.Rewrite = map[string]string{
		"/dashboard": "/dashboard",
		"/ranking":   "/ranking",
		"/missions":  "/missions",
	}

	e.Use(middleware.ProxyWithConfig(proxyConfig))

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// connect to database
	db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// migrate tables
	if err := migration.MigrateTables(db.DB); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(repo, config.TraqOAuth2())
	v1API := e.Group("/api/v1")
	h.SetupRoutes(v1API)

	e.Logger.Fatal(e.Start(config.AppAddr()))
}

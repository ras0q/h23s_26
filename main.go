package main

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
	"github.com/traP-jp/h23s_26/internal/handler"
	"github.com/traP-jp/h23s_26/internal/pkg/config"
	"github.com/traP-jp/h23s_26/internal/repository"
	"github.com/traP-jp/h23s_26/internal/repository/migration"
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

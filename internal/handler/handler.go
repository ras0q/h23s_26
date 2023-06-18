package handler

import (
	traqoauth2 "github.com/ras0q/traq-oauth2"
	"github.com/traP-jp/h23s_26/internal/handler/middleware"
	"github.com/traP-jp/h23s_26/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo             *repository.Repository
	traqOauth2Config *traqoauth2.Config
}

func New(repo *repository.Repository, traqOauth2Config *traqoauth2.Config) *Handler {
	return &Handler{
		repo:             repo,
		traqOauth2Config: traqOauth2Config,
	}
}

func (h *Handler) SetupRoutes(api *echo.Group) {
	// Ping API
	pingAPI := api.Group("/ping")
	{
		pingAPI.GET("", h.Ping)
	}

	// User API
	userAPI := api.Group("/users")
	{
		userAPI.GET("", h.GetUsers)
		userAPI.POST("", h.PostUser)
		userAPI.GET("/:userID", h.GetUser)
		userAPI.GET("/me", h.GetMe, middleware.TrapAuth())
		userAPI.PATCH("/:userID/missions/:missionID", h.PatchMission, middleware.TrapAuth())
	}

	// Mission api
	missionAPI := api.Group("/missions")
	{
		missionAPI.GET("", h.GetMissions)
		missionAPI.POST("", h.PostMission, middleware.TrapAuth())
		missionAPI.GET("/:missionID", h.GetMission)
	}

	// Ranking API
	rankingAPI := api.Group("/ranking")
	{
		rankingAPI.GET("", h.GetRanking)
	}

	// Oauth2 API
	oauth2API := api.Group("/oauth2")
	{
		oauth2API.GET("/authorize", h.Authorize)
		oauth2API.GET("/callback", h.Callback)
	}
}

package handler

import (
	"github.com/traP-jp/h23s_26/internal/handler/middleware"
	"github.com/traP-jp/h23s_26/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
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
		// userAPI.GET("/me", h.GetMe, middleware.TRAPAuth())
		// userAPI.PATCH("/me/missions/:missionID", h.PatchMission, middleware.TRAPAuth())
	}

	// Mission api
	missionAPI := api.Group("/missions")
	{
		missionAPI.GET("", h.GetMissions)
		missionAPI.POST("", h.PostMission, middleware.TRAPAuth())
		missionAPI.GET("/:missionID", h.GetMission)
	}

	// // Ranking API
	// rankingAPI := api.Group("/ranking")
	// {
	// 	rankingAPI.GET("", h.GetRanking)
	// }
}

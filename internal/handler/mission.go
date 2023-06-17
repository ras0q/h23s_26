package handler

import (
	"fmt"
	"net/http"

	vd "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_26/internal/repository"
)

// スキーマ定義
type (
	GetMissionsResponse []GetMissionResponse

	GetMissionResponse struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Achievers   []string  `json:"achievers"`
	}

	CreateMissionRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	CreateMissionResponse struct {
		ID uuid.UUID
	}
)

// GET /api/v1/missions/
func (h *Handler) GetMissions(c echo.Context) error {
	missions, err := h.repo.GetMissions(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := make(GetMissionsResponse, len(missions))
	for i, mission := range missions {
		res[i] = GetMissionResponse{
			ID:          mission.ID,
			Name:        mission.Name,
			Description: mission.Description,
			Achievers:   mission.Achievers,
		}
	}

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/missions/:missionID
func (h *Handler) GetMission(c echo.Context) error {
	missionID, err := uuid.Parse(c.Param("missionID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid missionID").SetInternal(err)
	}

	mission, err := h.repo.GetMission(c.Request().Context(), missionID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := GetMissionResponse{
		ID:          mission.ID,
		Name:        mission.Name,
		Description: mission.Description,
		Achievers:   mission.Achievers,
	}

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/missions :traP authorization needed:
func (h *Handler) PostMission(c echo.Context) error {
	req := new(CreateMissionRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}
	err := vd.ValidateStruct(
		req,
		vd.Field(&req.Name, vd.Required),
		vd.Field(&req.Description, vd.Required),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err).Error()).SetInternal(err)
	}

	params := repository.CreateMissionParams{
		Name:        req.Name,
		Description: req.Description,
		CreatorID:   "user1",
	}

	missionID, err := h.repo.PostMission(c.Request().Context(), params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := CreateMissionResponse{
		ID: missionID,
	}

	return c.JSON(http.StatusOK, res)

}

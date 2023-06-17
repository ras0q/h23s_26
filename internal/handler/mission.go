package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//スキーマ定義
type (
	GetMissionsResponse []GetMissionResponse

	GetMissionResponse struct {
		ID uuid.UUID `db:"id"`
		Name string `db:"name"`
		Description string `db:"description"`
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
			ID:    mission.ID,
			Name:  mission.Name,
			Description: mission.Description,
		}
	}

	return c.JSON(http.StatusOK, res)
}


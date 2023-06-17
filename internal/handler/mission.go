package handler

import (
	//"fmt"
	//"go-backend-sample/internal/repository"
	"net/http"

	//vd "github.com/go-ozzo/ozzo-validation"
	//"github.com/go-ozzo/ozzo-validation/is"
	//"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//スキーマ定義
type (
	GetMissionsResponse []GetMissionResponse

	GetMissionResponse struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Achivers []string `json:"achivers"`
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
			Achivers: mission.Achivers,
		}
	}

	return c.JSON(http.StatusOK, res)
}


package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// スキーマ定義
type (
	GetRankingResponse struct {
		Ranking []string `json:"ranking"`
	}
)

// GET /api/v1/ranking
func (h *Handler) GetRanking(c echo.Context) error {

	ranking, err := h.repo.GetRanking(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := GetRankingResponse{
		Ranking: ranking.Ranking,
	}

	return c.JSON(http.StatusOK, res)

}

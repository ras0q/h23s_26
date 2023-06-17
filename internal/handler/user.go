package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	GetUsersResponse []GetUserResponse

	GetUserResponse struct {
		ID string `json:"id"`
		// Ranking  int         `json:"ranking"`
		Achieves []uuid.UUID `json:"achieves"`
	}
)

// GET /users
func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := make(GetUsersResponse, len(users))
	for i, user := range users {
		res[i] = GetUserResponse{
			ID: user.ID,
			// Ranking:  user.Ranking,
			Achieves: user.AchieveMissions,
		}
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetUser(c echo.Context) error {
	user, err := h.repo.GetUser(c.Request().Context(), c.Param("userID"))

	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound).SetInternal(err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := GetUserResponse{
		ID:       user.ID,
		Achieves: user.AchieveMissions,
	}

	return c.JSON(http.StatusOK, res)

}

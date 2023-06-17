package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	GetUsersResponse []GetUserResponse

	GetUserResponse struct {
		ID string `json:"id"`
		// Ranking  int         `json:"ranking"`
		// Achieves []uuid.UUID `json:"achieves"`
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
		}
	}

	return c.JSON(http.StatusOK, res)
}

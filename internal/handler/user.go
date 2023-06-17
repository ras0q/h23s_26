package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	vd "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_26/internal/repository"
)

type (
	GetUsersResponse []GetUserResponse

	GetUserResponse struct {
		ID string `json:"id"`
		// Ranking  int         `json:"ranking"`
		Achieves []uuid.UUID `json:"achieves"`
	}

	CreateUserRequest struct {
		ID string `json:"id"`
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

// POST /users
func (h *Handler) PostUser(c echo.Context) error {
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}

	err := vd.ValidateStruct(
		req,
		vd.Field(&req.ID, vd.Required),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err).Error()).SetInternal(err)
	}

	param := repository.CreateUserParams{
		ID: req.ID,
	}

	err = h.repo.PostUser(c.Request().Context(), param)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.JSON(http.StatusOK, param)

}

// GET /users/:userID
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

package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"

	vd "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_26/internal/pkg/config"
	"github.com/traP-jp/h23s_26/internal/repository"
)

type (
	GetUsersResponse []GetUserResponse

	GetUserResponse struct {
		ID       string      `json:"id"`
		Ranking  int         `json:"ranking"`
		Achieves []uuid.UUID `json:"achieves"`
	}

	CreateUserRequest struct {
		ID string `json:"id"`
	}

	PatchMissionRequest struct {
		Clear     bool      `json:"clear"`
		ClearedAt time.Time `json:"cleared_at"`
	}
)

// GET /users
func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	sort.Slice(users, func(i, j int) bool {
		return len(users[i].AchieveMissions) > len(users[j].AchieveMissions)
	})

	res := make(GetUsersResponse, len(users))
	for i, user := range users {
		res[i] = GetUserResponse{
			ID:       user.ID,
			Ranking:  i + 1,
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

	return c.NoContent(http.StatusCreated)

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

// GET /users/me
func (h *Handler) GetMe(c echo.Context) error {
	userID := c.Get(string(config.TraqIDKey)).(string)

	user, err := h.repo.GetUser(c.Request().Context(), userID)
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

// PATCH /users/:userID/missions/:missionID

func (h *Handler) PatchMission(c echo.Context) error {
	missionID, err := uuid.Parse(c.Param("missionID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid missionID").SetInternal(err)
	}

	req := new(PatchMissionRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}

	err = vd.ValidateStruct(
		req,
		vd.Field(&req.Clear),
		vd.Field(&req.ClearedAt, vd.Required),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err).Error()).SetInternal(err)
	}

	param := repository.PatchMissionParams{
		Clear:     req.Clear,
		UserID:    c.Param("userID"),
		MissionID: missionID,
	}

	err = h.repo.PatchMission(c.Request().Context(), param)

	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound).SetInternal(err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)

}

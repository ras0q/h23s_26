package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	// missons table
	Mission struct {
		ID          uuid.UUID `db:"id"`
		Name        string    `db:"name"`
		Description string    `db:"description"`
		CreatorID   string    `db:"creator_id"`
		Achievers   []string
	}

	CreateMissionParams struct {
		Name        string
		Description string
		CreatorID   string
	}
)

func (r *Repository) GetMissions(ctx context.Context) ([]*Mission, error) {
	missions := []*Mission{}
	if err := r.db.SelectContext(ctx, &missions, "SELECT * FROM missions"); err != nil {
		return nil, fmt.Errorf("select missions: %w", err)
	}

	if err := r.db.GetContext(ctx, &missions, "SELECT  FROM missions"); err != nil {
		return nil, fmt.Errorf("select missions: %w", err)
	}

	achieveUsers := make([]*UserMissionRelation, 0)

	for _, mission := range missions {
		for _, achieveUser := range achieveUsers {
			if mission.ID == achieveUser.MissionID{
				mission.Achievers = append(mission.Achievers, achieveUser.UserID)
			}
		}
	}

	return missions, nil
}

func (r *Repository) GetMission(ctx context.Context, missionID uuid.UUID) (*Mission, error) {
	mission := Mission{}
	if err := r.db.GetContext(ctx, &mission, "SELECT * FROM missions WHERE id = ?", missionID); err != nil {
		return nil, fmt.Errorf("select missions: %w", err)
	}

	return &mission, nil
}

func (r *Repository) PostMission(ctx context.Context, params CreateMissionParams) (uuid.UUID, error) {
	missionID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO missions (id, name, description,creator_id) VALUES (?, ?, ?, ?)", missionID, params.Name, params.Description, params.CreatorID); err != nil {
		return uuid.Nil, fmt.Errorf("insert mission: %w", err)
	}

	return missionID, nil
}

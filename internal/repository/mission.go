package repository

import (
	"context"
	"fmt"

	//"github.com/google/uuid"
)

type(
	// missons table
	Mission struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Achivers []string `json:"achivers"`
	}

)

func (r *Repository) GetMissions(ctx context.Context) ([]*Mission, error) {
	missions := []*Mission{}
	if err := r.db.SelectContext(ctx, &missions, "SELECT * FROM missions"); err != nil {
		return nil, fmt.Errorf("select missions: %w", err)
	}

	return missions, nil
}
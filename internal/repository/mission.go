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
	}
)

func (r *Repository) GetMissions(ctx context.Context) ([]*Mission, error) {
	missions := []*Mission{}
	if err := r.db.SelectContext(ctx, &missions, "SELECT * FROM missions"); err != nil {
		return nil, fmt.Errorf("select missions: %w", err)
	}

	return missions, nil
}

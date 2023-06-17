package repository

import (
	"context"
	"fmt"
)

type (
	Ranking struct {
		Ranking []string `json:"ranking"`
	}
)

func (r *Repository) GetRanking(ctx context.Context) (*Ranking, error) {
	ranking := Ranking{}
	if err := r.db.SelectContext(ctx, &ranking.Ranking, "SELECT user_id FROM user_mission_relations GROUP BY user_id ORDER BY COUNT(*) DESC LIMIT 5"); err != nil {
		return nil, fmt.Errorf("ranking: %w", err)
	}

	return &ranking,nil

}

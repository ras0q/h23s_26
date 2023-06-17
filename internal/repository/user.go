package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	// users table
	User struct {
		ID              string      `db:"id"` // primary key
		AchieveMissions []uuid.UUID `db:"-"`
	}

	// user_mission_relations table
	UserMissionRelation struct {
		ID        uuid.UUID `db:"id"`         // primary key
		UserID    string    `db:"user_id"`    // foreign key
		MissionID uuid.UUID `db:"mission_id"` // foreign key
	}
)

func (r *Repository) GetUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := r.db.SelectContext(ctx, &users, "SELECT * FROM users"); err != nil {
		return nil, fmt.Errorf("get users from db: %w", err)
	}

	achieveMissons := []*UserMissionRelation{}
	if err := r.db.SelectContext(ctx, &achieveMissons, "SELECT * FROM user_mission_relations"); err != nil {
		return nil, fmt.Errorf("get user_mission_relations from db: %w", err)
	}

	for _, user := range users {
		user.AchieveMissions = []uuid.UUID{}
		for _, achieveMission := range achieveMissons {
			if user.ID == achieveMission.UserID {
				user.AchieveMissions = append(user.AchieveMissions, achieveMission.ID)
			}
		}
	}

	return users, nil
}

func (r *Repository) GetUser(ctx context.Context, userID string) (*User, error) {
	user := User{}
	if err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id= ? ", userID); err != nil {
		return nil, fmt.Errorf("get users from db: %w", err)
	}

	if err := r.db.SelectContext(ctx, &user.AchieveMissions, "SELECT mission_id FROM user_mission_relations WHERE user_id= ? ", userID); err != nil {
		return nil, fmt.Errorf("get user_mission_relations from db: %w", err)
	}

	return &user, nil

}

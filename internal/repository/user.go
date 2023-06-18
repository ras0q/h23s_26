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
		Rank            int         `db:"-"`
		AchieveMissions []uuid.UUID `db:"-"`
	}

	// user_mission_relations table
	UserMissionRelation struct {
		ID        uuid.UUID `db:"id"`         // primary key
		UserID    string    `db:"user_id"`    // foreign key
		MissionID uuid.UUID `db:"mission_id"` // foreign key
	}

	CreateUserParams struct {
		ID string
	}

	PatchMissionParams struct {
		Clear     bool      `db:"-"`
		UserID    string    `db:"user_id"`
		MissionID uuid.UUID `db:"mission_id"`
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

	if err := r.db.GetContext(ctx, &user.Rank, "SELECT COUNT(*)+1 FROM (SELECT user_id FROM user_mission_relations GROUP BY user_id HAVING COUNT(*) > (SELECT COUNT(*) FROM user_mission_relations WHERE user_id = ?)) AS tmp", userID); err != nil {
		return nil, fmt.Errorf("get user_mission_relations from db: %w", err)
	}

	user.AchieveMissions = []uuid.UUID{}
	if err := r.db.SelectContext(ctx, &user.AchieveMissions, "SELECT mission_id FROM user_mission_relations WHERE user_id= ? ", userID); err != nil {
		return nil, fmt.Errorf("get user_mission_relations from db: %w", err)
	}

	return &user, nil
}

func (r *Repository) PostUser(ctx context.Context, params CreateUserParams) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO users (id) VALUES (?)", params.ID); err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}

func (r *Repository) PatchMission(ctx context.Context, params PatchMissionParams) error {
	if params.Clear {
		patchID := uuid.New()

		if _, err := r.db.ExecContext(ctx, "INSERT IGNORE INTO user_mission_relations(id,user_id,mission_id) VALUES(?,?,?)", patchID, params.UserID, params.MissionID); err != nil {
			return fmt.Errorf("patch mission: %w", err)
		}
		return nil

	}

	if _, err := r.db.ExecContext(ctx, "DELETE FROM user_mission_relations WHERE user_id= ? AND mission_id= ? ", params.UserID, params.MissionID); err != nil {
		return fmt.Errorf("patch mission: %w", err)
	}

	return nil
}

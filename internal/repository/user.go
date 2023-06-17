package repository

import (
	"context"
	"fmt"
)

type (
	// users table
	User struct {
		ID string `db:"id"` // primary key
	}
)

func (r *Repository) GetUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)

	if err := r.db.SelectContext(ctx, &users, "SELECT * FROM users"); err != nil {
		return nil, fmt.Errorf("get users from db: %w", err)
	}

	return users, nil
}

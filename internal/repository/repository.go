package repository

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

//go:embed schema.sql
var schema string

func (r *Repository) SetupTables() error {
	// TODO: 初期化処理何とかする
	var user User
	if err := r.db.Get(&user, "SELECT * FROM users LIMIT 1"); err == nil {
		return nil
	}

	for _, query := range strings.Split(schema, ";") {
		if len(query) > 0 {
			if _, err := r.db.Exec(query); err != nil {
				return fmt.Errorf("setup tables: %w", err)
			}
		}
	}

	return nil
}

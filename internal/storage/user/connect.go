package user

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type userRepo struct {
	db *sql.DB
}

func New(dbURl string) (*userRepo, error) {
	db, err := sql.Open("postgres", dbURl)
	if err != nil {
		return nil, fmt.Errorf("connect open a db driver: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to a db: %w", err)
	}
	return &userRepo{db}, nil
}

func (db *userRepo) tx(ctx context.Context) (*sql.Tx, error) {
	return db.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
}

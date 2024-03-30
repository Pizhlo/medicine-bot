package drugs

import (
	"context"
	"database/sql"
	"fmt"
)

type drugsRepo struct {
	db *sql.DB
}

func New(dbURl string) (*drugsRepo, error) {
	db, err := sql.Open("postgres", dbURl)
	if err != nil {
		return nil, fmt.Errorf("connect open a db driver: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("cannot connect to a db: %w", err)
	}
	return &drugsRepo{db}, nil
}

func (db *drugsRepo) tx(ctx context.Context) (*sql.Tx, error) {
	return db.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
}

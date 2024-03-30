package user

import (
	"context"
	"fmt"
)

func (db *userRepo) Save(ctx context.Context, tgID int64) error {
	tx, err := db.tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `insert into users.users (tg_id) values($1) on conflict (tg_id) do nothing`, tgID)
	if err != nil {
		return fmt.Errorf("error while saving user in DB: %w", err)
	}

	return tx.Commit()
}

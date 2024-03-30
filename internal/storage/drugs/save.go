package drugs

import (
	"context"
	"fmt"

	"github.com/Pizhlo/medicine-bot/internal/model"
)

func (db *drugsRepo) Save(ctx context.Context, tgID int64, drug model.Drug) error {
	tx, err := db.tx(ctx)
	if err != nil {
		return nil
	}

	_, err = tx.ExecContext(ctx, `insert into drugs.drugs (user_id, name, desription, created, take_today, today_count) values((select id from users.users where tg_id = $1), $2, $3, $4, $5, $6)`,
		tgID, drug.Name, drug.Description, drug.Created, drug.TakeToday, drug.TodayCount)
	if err != nil {
		return fmt.Errorf("error while saving user in DB: %w", err)
	}

	return tx.Commit()

}

package drugs

import (
	"context"
	"fmt"

	"github.com/Pizhlo/medicine-bot/internal/model"
)

func (db *drugsRepo) Save(ctx context.Context, drug *model.Drug) error {
	tx, err := db.tx(ctx)
	if err != nil {
		return nil
	}

	_, err = tx.ExecContext(ctx, `insert into drugs.drugs (user_id, name, description, take_today, today_count) values((select id from users.users where tg_id = $1), $2, $3, $4, $5)`,
		drug.UserID, drug.Name, drug.Description.String, drug.TakeToday, drug.TodayCount)
	if err != nil {
		return fmt.Errorf("error while saving user in DB: %w", err)
	}

	return tx.Commit()

}

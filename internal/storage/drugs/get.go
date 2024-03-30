package drugs

import (
	"context"
	"database/sql"
	"errors"

	api_errors "github.com/Pizhlo/medicine-bot/internal/errors"
	"github.com/Pizhlo/medicine-bot/internal/model"
)

func (db *drugsRepo) GetbyUser(ctx context.Context, tgId int64) ([]model.Drug, error) {
	rows, err := db.db.QueryContext(ctx, "select drug_number, drugs.drugs.user_id, name, desription, created, take_today, today_count from drugs.drugs join drugs.drugs_view on drugs.drugs_view.user_id = drugs.drugs.user_id where drugs.drugs.user_id = (select id from users.users where tg_id = $1);", tgId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.ErrDrugsNotFound
		}
		return nil, err
	}

	result := []model.Drug{}

	for rows.Next() {
		drug := model.Drug{}
		err := rows.Scan(&drug.ViewID, &drug.UserID, &drug.Name, &drug.Description, &drug.Created, &drug.TakeToday, &drug.TodayCount)
		if err != nil {
			return nil, err
		}

		result = append(result, drug)
	}

	if len(result) == 0 {
		return nil, api_errors.ErrDrugsNotFound
	}

	return result, nil
}

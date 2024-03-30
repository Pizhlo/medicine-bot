package user

import "context"

func (db *userRepo) GetAll(ctx context.Context) ([]int64, error) {
	rows, err := db.db.QueryContext(ctx, "select tg_id from users.users")
	if err != nil {
		return nil, err
	}

	result := []int64{}

	for rows.Next() {
		id := int64(0)
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		result = append(result, id)
	}

	return result, nil
}

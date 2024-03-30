package model

import (
	"database/sql"
)

type Drug struct {
	ID          int
	ViewID      int
	UserID      int64
	Name        string
	Description sql.NullString
	TakeToday   bool
	TodayCount  int
}

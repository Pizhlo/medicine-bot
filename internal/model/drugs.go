package model

import (
	"database/sql"
	"time"
)

type Drug struct {
	ID          int
	ViewID      int
	UserID      int64
	Name        string
	Description sql.NullString
	Created     time.Time
	TakeToday   bool
	TodayCount  int
}

package data

import "time"

//
// データを保持するための構造体
type Thread struct {
	ID        int
	Uuid      string
	Topic     string
	UserID    string
	CreatedAt time.Time
}

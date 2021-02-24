package model

import "time"

type User struct {
	ID          uint64 `gorm:"primaryKey"`
	Username    string
	Email       string
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

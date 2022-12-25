package entity

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar; unique;not null"`
	Amount    int       `json:"amount" gorm:"type:integer; null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}

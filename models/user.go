package models

import "time"

type Favorite struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	VideoID   string    `json:"video_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	User User
}

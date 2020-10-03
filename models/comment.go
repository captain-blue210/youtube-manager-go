package models

import "time"

type Comment struct {
	ID        uint       `gorm:"primary_key"`
	UserID    uint       `json:"user_id"`
	VideoID   string     `json:"video_id"`
	Comment   string     `json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`

	User User
}

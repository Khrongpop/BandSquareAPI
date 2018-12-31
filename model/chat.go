package model

import "time"

type Chat struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	UserID       uint      `json:"user_id"`
	ToID         uint      `json:"to_id"`
	Message      string    `json:"message"`
	Seen         bool      `json:"seen"`
	User         User      `gorm:"foreignkey:UserID" json:"user"`
	ToUser       User      `gorm:"foreignkey:ToID" json:"to_user"`
	NotSeenCount *int      `json:"not_seen_count"`
	CreatedAt    time.Time `json:"created_at"`
}

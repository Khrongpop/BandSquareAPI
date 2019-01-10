package model

import "time"

type Notification struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    int       `json:"user_id"`
	Status    int       `gorm:"default:1" json:"status"`
	BookingID *int      `json:"booking_id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `gorm:"foreignkey:user_id" json:"user"`
	Booking   *Booking  `gorm:"foreignkey:booking_id" json:"booking"`
}

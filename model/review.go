package model

import "time"

type Review struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	BandID    uint      `json:"band_id"`
	BookingID uint      `json:"booking_id"`
	Rate      float32   `json:"rate"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	User      *User     `json:"user"`
	Booking   *Booking  `json:"booking"`
}

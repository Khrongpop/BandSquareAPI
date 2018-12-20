package model

import (
	"time"
)

type Booking struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	UserID           uint      `json:"user_id"`
	BandID           uint      `json:"band_id"`
	CategoryID       uint      `json:"category_id"`
	TypeID           uint      `json:"type_id"`
	Status           int       `gorm:"default:0" json:"status"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	Location         string    `json:"location"`
	Date             time.Time `json:"date"`
	Time             time.Time `json:"time"`
	Duration         time.Time `json:"duration"`
	Price            float64   `json:"price"`
	Instrument       bool      `gorm:"default:false" json:"instrument"`
	InstrumentDetail *string   `json:"instrument_detail"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	User             User      `json:"user"`
	Band             *Band     `json:"band"`
	Category         Category  `json:"category"`
	Type             Type      `json:"type"`
}

package model

type Notification struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserID    int    `json:"user_id"`
	Status    int    `gorm:"default:1" json:"status"`
	BookingID *int   `json:"booking_id"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
}

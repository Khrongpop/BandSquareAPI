package model

import (
	"time"
)

type Role struct {
	ID   uint   `gorm:"primary_key" json:id`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type User struct {
	// gorm.Model
	ID        uint   `gorm:"primary_key" json:id`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
	Active    uint   `json:active`
	RoleID    uint   `json:role_id`
	Role      Role
	Band      Band
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}

type Band struct {
	// gorm.Model
	ID           uint      `gorm:"primary_key" json:id`
	About        string    `json:"about"`
	Member       string    `json:"member"`
	WorkLocation string    `json:"work_location"`
	MinPrice     int64     `json:"min_price"`
	MaxPrice     int64     `json:"max_price"`
	Cover        string    `json:"cover"`
	UserID       uint      `json:user_id`
	CreatedAt    time.Time `json:created_at`
	User         *User
}

type BandType struct {
	BandID uint   `gorm:"primary_key" json:band_id`
	TypeID uint   `gorm:"primary_key" json:type_id`
	Detail string `json:"detail"`
	Band   Band
	Type   Type
}
type Category struct {
	ID   uint   `gorm:"primary_key" json:id`
	Name string `json:"name"`
}
type Genre struct {
	ID   uint   `gorm:"primary_key" json:id`
	Name string `json:"name"`
}
type Type struct {
	ID   uint   `gorm:"primary_key" json:id`
	Name string `json:"name"`
}

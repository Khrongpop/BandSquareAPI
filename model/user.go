package model

import (
	"time"
)

type Role struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type User struct {
	// gorm.Model
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Image     string    `json:"image"`
	Thumbnail string    `json:"thumbnail"`
	Active    bool      `gorm:"default:0" json:"active"`
	RoleID    int8      `gorm:"default:1" json:"role_id"`
	Role      Role      `json:"role"`
	Band      *Band     `json:"band"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SocailAccount struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	UserID     uint   `json:"user_id"`
	ProviderID string `json:"provider_id"`
	Provider   string `json:"provider"`
	// 	// 2343960062310644
}
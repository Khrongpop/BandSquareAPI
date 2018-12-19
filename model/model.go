package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
	RoleID    uint   `json:role_id`
	Role      Role
}
type Role struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
}

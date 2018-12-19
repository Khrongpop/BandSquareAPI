package seeder

import (
	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

var db *gorm.DB

func RoleSeed(db *gorm.DB) {
	db.Create(&model.Role{
		Name: "Member",
		Slug: "Member",
	})
	db.Create(&model.Role{
		Name: "Musician",
		Slug: "Music",
	})
	db.Create(&model.Role{
		Name: "Administrator",
		Slug: "Admin",
	})
}

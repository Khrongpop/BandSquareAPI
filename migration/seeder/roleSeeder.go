package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

var db *gorm.DB

func RoleSeed(db *gorm.DB) {
	fmt.Println("seed role ...")
	db.Create(&model.Role{
		ID:   1,
		Name: "Member",
		Slug: "Member",
	})
	db.Create(&model.Role{
		ID:   2,
		Name: "Musician",
		Slug: "Music",
	})
	db.Create(&model.Role{
		ID:   3,
		Name: "Administrator",
		Slug: "Admin",
	})
}

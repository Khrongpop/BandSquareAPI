package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func CategorySeed(db *gorm.DB) {
	fmt.Println("seed category ...")
	db.Create(&model.Category{
		ID:   1,
		Name: "Wedding",
	})
	db.Create(&model.Category{
		ID:   2,
		Name: "Event",
	})
	db.Create(&model.Category{
		ID:   3,
		Name: "Party",
	})
	db.Create(&model.Category{
		ID:   4,
		Name: "Other",
	})
}

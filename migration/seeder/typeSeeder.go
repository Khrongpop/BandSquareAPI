package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func TypeSeed(db *gorm.DB) {
	fmt.Println("seed type ...")
	db.Create(&model.Type{
		ID:   1,
		Name: "Acoustic",
	})
	db.Create(&model.Type{
		ID:   2,
		Name: "FullBand",
	})
	db.Create(&model.Type{
		ID:   3,
		Name: "DJ",
	})
	db.Create(&model.Type{
		ID:   4,
		Name: "String",
	})
	db.Create(&model.Type{
		ID:   5,
		Name: "Jazz",
	})
}

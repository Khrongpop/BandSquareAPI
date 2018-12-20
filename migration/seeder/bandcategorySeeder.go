package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BandCategorySeed(db *gorm.DB) {
	fmt.Println("seed BandGenre ...")
	db.Create(&model.BandCategory{
		BandID:     1,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     1,
		CategoryID: 2,
	})
	db.Create(&model.BandCategory{
		BandID:     1,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     1,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     2,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     2,
		CategoryID: 2,
	})
	db.Create(&model.BandCategory{
		BandID:     2,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     2,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     3,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     3,
		CategoryID: 2,
	})
	db.Create(&model.BandCategory{
		BandID:     3,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     3,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     4,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     4,
		CategoryID: 2,
	})
	db.Create(&model.BandCategory{
		BandID:     4,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     4,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     5,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     5,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     5,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     6,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     6,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     7,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     8,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     8,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     9,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     9,
		CategoryID: 2,
	})
	db.Create(&model.BandCategory{
		BandID:     9,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     9,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     10,
		CategoryID: 1,
	})
	db.Create(&model.BandCategory{
		BandID:     10,
		CategoryID: 2,
	})
	db.Create(&model.BandCategory{
		BandID:     10,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     10,
		CategoryID: 4,
	})
	db.Create(&model.BandCategory{
		BandID:     11,
		CategoryID: 2,
	})
	db.Create(&model.BandCategory{
		BandID:     11,
		CategoryID: 3,
	})
	db.Create(&model.BandCategory{
		BandID:     11,
		CategoryID: 4,
	})
}

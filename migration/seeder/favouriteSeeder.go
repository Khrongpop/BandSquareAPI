package seeder

import (
	"fmt"

	"github.com/khrongpop/BandSquareAPI/model"

	"github.com/jinzhu/gorm"
)

func FavouriteSeed(db *gorm.DB) {
	fmt.Println("seed Favourite ...")
	db.Create(&model.Favourite{
		UserID: 2,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 2,
		BandID: 3,
	})
	db.Create(&model.Favourite{
		UserID: 2,
		BandID: 4,
	})
	db.Create(&model.Favourite{
		UserID: 2,
		BandID: 7,
	})
	db.Create(&model.Favourite{
		UserID: 13,
		BandID: 4,
	})
	db.Create(&model.Favourite{
		UserID: 13,
		BandID: 7,
	})
	db.Create(&model.Favourite{
		UserID: 13,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 13,
		BandID: 2,
	})
	db.Create(&model.Favourite{
		UserID: 13,
		BandID: 3,
	})
	db.Create(&model.Favourite{
		UserID: 14,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 14,
		BandID: 2,
	})
	db.Create(&model.Favourite{
		UserID: 14,
		BandID: 3,
	})
	db.Create(&model.Favourite{
		UserID: 14,
		BandID: 4,
	})
	db.Create(&model.Favourite{
		UserID: 14,
		BandID: 5,
	})
	db.Create(&model.Favourite{
		UserID: 15,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 15,
		BandID: 2,
	})
	db.Create(&model.Favourite{
		UserID: 15,
		BandID: 5,
	})
	db.Create(&model.Favourite{
		UserID: 15,
		BandID: 8,
	})
	db.Create(&model.Favourite{
		UserID: 15,
		BandID: 9,
	})
	db.Create(&model.Favourite{
		UserID: 16,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 16,
		BandID: 2,
	})
	db.Create(&model.Favourite{
		UserID: 16,
		BandID: 4,
	})
	db.Create(&model.Favourite{
		UserID: 16,
		BandID: 7,
	})
	db.Create(&model.Favourite{
		UserID: 16,
		BandID: 6,
	})
	db.Create(&model.Favourite{
		UserID: 17,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 17,
		BandID: 2,
	})
	db.Create(&model.Favourite{
		UserID: 17,
		BandID: 3,
	})
	db.Create(&model.Favourite{
		UserID: 17,
		BandID: 4,
	})
	db.Create(&model.Favourite{
		UserID: 18,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 18,
		BandID: 4,
	})
	db.Create(&model.Favourite{
		UserID: 18,
		BandID: 5,
	})
	db.Create(&model.Favourite{
		UserID: 18,
		BandID: 7,
	})
	db.Create(&model.Favourite{
		UserID: 19,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 19,
		BandID: 2,
	})
	db.Create(&model.Favourite{
		UserID: 19,
		BandID: 3,
	})
	db.Create(&model.Favourite{
		UserID: 19,
		BandID: 5,
	})
	db.Create(&model.Favourite{
		UserID: 19,
		BandID: 6,
	})
	db.Create(&model.Favourite{
		UserID: 20,
		BandID: 1,
	})
	db.Create(&model.Favourite{
		UserID: 20,
		BandID: 2,
	})
	db.Create(&model.Favourite{
		UserID: 20,
		BandID: 3,
	})
	db.Create(&model.Favourite{
		UserID: 20,
		BandID: 4,
	})
	db.Create(&model.Favourite{
		UserID: 20,
		BandID: 5,
	})
}

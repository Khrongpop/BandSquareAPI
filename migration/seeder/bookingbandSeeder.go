package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BookingBandSeed(db *gorm.DB) {
	fmt.Println("seed BookingBandSeed ...")
	db.Create(&model.BookingBand{
		BookingID: 1,
		BandID:    1,
	})
	db.Create(&model.BookingBand{
		BookingID: 1,
		BandID:    2,
	})
	db.Create(&model.BookingBand{
		BookingID: 1,
		BandID:    3,
	})
}

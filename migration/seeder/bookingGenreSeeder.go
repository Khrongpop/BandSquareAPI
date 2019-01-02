package seeder

import (
	"fmt"

	"github.com/khrongpop/BandSquareAPI/model"

	"github.com/jinzhu/gorm"
)

func BookingGenreSeed(db *gorm.DB) {
	fmt.Println("seed bookingGenreSeed ...")
	db.Create(&model.BookingGenre{
		BookingID: 1,
		GenreID:   1,
	})
	db.Create(&model.BookingGenre{
		BookingID: 1,
		GenreID:   2,
	})

}

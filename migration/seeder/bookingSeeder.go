package seeder

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BookingSeed(db *gorm.DB) {
	fmt.Println("seed booking ...")
	db.Create(&model.Booking{
		ID:         1,
		BandID:     nil,
		UserID:     2,
		CategoryID: 1,
		TypeID:     2,
		Status:     0,
		Latitude:   13.915030,
		Longitude:  100.551289,
		Location:   `Silpakorn University`,
		Date:       time.Now().AddDate(1, 2, 0),
		Time:       time.Now().AddDate(1, 2, 0).Add(20 * time.Hour).Add(30 * time.Minute),
		Duration:   `1:30:00`,
		Price:      15000,
		Instrument: false,
	})

	db.Create(&model.Booking{
		ID:         2,
		BandID:     getID(1),
		UserID:     2,
		CategoryID: 2,
		TypeID:     1,
		Status:     2,
		Latitude:   13.915030,
		Longitude:  100.551289,
		Location:   `Silpakorn University`,
		Date:       time.Now().AddDate(1, 4, 0),
		Time:       time.Now().AddDate(1, 4, 0).Add(10 * time.Hour).Add(25 * time.Minute),
		Duration:   `1:30:00`,
		Price:      15000,
		Instrument: false,
	})
}

func getID(x uint) *uint {
	return &x
}

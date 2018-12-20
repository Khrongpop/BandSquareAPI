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
		DateTime:   time.Now().AddDate(1, 2, 0).Add(20 * time.Hour).Add(30 * time.Minute),
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
		DateTime:   time.Now().AddDate(1, 4, 0).Add(10 * time.Hour).Add(25 * time.Minute),
		Duration:   `1:30:00`,
		Price:      15000,
		Instrument: false,
	})

	db.Create(&model.Booking{
		ID:         3,
		BandID:     getID(2),
		UserID:     2,
		CategoryID: 3,
		TypeID:     2,
		Status:     3,
		Latitude:   13.915030,
		Longitude:  100.551289,
		Location:   `Silpakorn University`,
		DateTime:   time.Now().AddDate(0, -5, 0).Add(40 * time.Hour).Add(45 * time.Minute),
		Duration:   `1:30:00`,
		Price:      12000,
		Instrument: false,
	})

	db.Create(&model.Booking{
		ID:         4,
		BandID:     getID(3),
		UserID:     16,
		CategoryID: 1,
		TypeID:     2,
		Status:     3,
		Latitude:   13.915030,
		Longitude:  100.551289,
		Location:   `Silpakorn University`,
		DateTime:   time.Now().AddDate(0, -1, 0).Add(27 * time.Hour).Add(18 * time.Minute),
		Duration:   `1:30:00`,
		Price:      17000,
		Instrument: false,
	})

	db.Create(&model.Booking{
		ID:         5,
		BandID:     getID(4),
		UserID:     15,
		CategoryID: 1,
		TypeID:     2,
		Status:     3,
		Latitude:   13.915030,
		Longitude:  100.551289,
		Location:   `Silpakorn University`,
		DateTime:   time.Now().AddDate(0, -7, 0).Add(47 * time.Hour).Add(58 * time.Minute),
		Duration:   `1:30:00`,
		Price:      17000,
		Instrument: false,
	})

	db.Create(&model.Booking{
		ID:         6,
		BandID:     getID(5),
		UserID:     13,
		CategoryID: 1,
		TypeID:     2,
		Status:     3,
		Latitude:   13.915030,
		Longitude:  100.551289,
		Location:   `Silpakorn University`,
		DateTime:   time.Now().AddDate(0, -2, 0).Add(10 * time.Hour).Add(30 * time.Minute),
		Duration:   `1:30:00`,
		Price:      10000,
		Instrument: false,
	})

	db.Create(&model.Booking{
		ID:         7,
		BandID:     getID(3),
		UserID:     17,
		CategoryID: 1,
		TypeID:     2,
		Status:     3,
		Latitude:   13.915030,
		Longitude:  100.551289,
		Location:   `Silpakorn University`,
		DateTime:   time.Now().AddDate(0, -2, 0).Add(40 * time.Hour).Add(50 * time.Minute),
		Duration:   `1:30:00`,
		Price:      10000,
		Instrument: false,
	})
}

func getID(x uint) *uint {
	return &x
}

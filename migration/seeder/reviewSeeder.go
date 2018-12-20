package seeder

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func ReviewSeed(db *gorm.DB) {
	fmt.Println("seed review ...")
	db.Create(&model.Review{
		ID:        1,
		BandID:    2,
		UserID:    2,
		BookingID: 3,
		Rate:      5,
		Detail:    `Good Job`,
		CreatedAt: time.Now().AddDate(0, -5, 0).Add(50 * time.Hour).Add(55 * time.Minute),
	})
	db.Create(&model.Review{
		ID:        2,
		BandID:    3,
		UserID:    16,
		BookingID: 4,
		Rate:      4,
		Detail:    `Well!!!`,
		CreatedAt: time.Now().AddDate(0, -1, 0).Add(57 * time.Hour).Add(58 * time.Minute),
	})
	db.Create(&model.Review{
		ID:        3,
		BandID:    15,
		UserID:    4,
		BookingID: 5,
		Rate:      5,
		Detail:    `The Best!!!`,
		CreatedAt: time.Now().AddDate(0, -7, 0).Add(57 * time.Hour).Add(58 * time.Minute),
	})
	db.Create(&model.Review{
		ID:        4,
		BandID:    5,
		UserID:    13,
		BookingID: 6,
		Rate:      4,
		Detail:    `Good Band`,
		CreatedAt: time.Now().AddDate(0, -2, 0).Add(50 * time.Hour).Add(50 * time.Minute),
	})
}

package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BandGenreSeed(db *gorm.DB) {
	fmt.Println("seed BandGenre ...")
	db.Create(&model.BandGenre{
		BandID:  1,
		GenreID: 1,
	})
	db.Create(&model.BandGenre{
		BandID:  1,
		GenreID: 3,
	})
	db.Create(&model.BandGenre{
		BandID:  1,
		GenreID: 4,
	})
	db.Create(&model.BandGenre{
		BandID:  1,
		GenreID: 5,
	})
	// db.Create(&model.BandGenre{
	// 	BandID:  2,
	// 	GenreID: 1,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  2,
	// 	GenreID: 2,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  3,
	// 	GenreID: 1,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  3,
	// 	GenreID: 4,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  4,
	// 	GenreID: 1,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  4,
	// 	GenreID: 2,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  4,
	// 	GenreID: 3,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  4,
	// 	GenreID: 7,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  4,
	// 	GenreID: 8,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  5,
	// 	GenreID: 1,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  5,
	// 	GenreID: 2,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  5,
	// 	GenreID: 8,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  6,
	// 	GenreID: 1,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  6,
	// 	GenreID: 2,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  6,
	// 	GenreID: 10,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  7,
	// 	GenreID: 1,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  7,
	// 	GenreID: 2,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  7,
	// 	GenreID: 3,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  7,
	// 	GenreID: 5,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  7,
	// 	GenreID: 8,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  7,
	// 	GenreID: 9,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  8,
	// 	GenreID: 2,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  8,
	// 	GenreID: 8,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  8,
	// 	GenreID: 9,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  9,
	// 	GenreID: 3,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  9,
	// 	GenreID: 11,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  10,
	// 	GenreID: 12,
	// })
	// db.Create(&model.BandGenre{
	// 	BandID:  11,
	// 	GenreID: 13,
	// })
}

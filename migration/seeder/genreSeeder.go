package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func GenreSeed(db *gorm.DB) {
	fmt.Println("seed genre ...")
	db.Create(&model.Genre{
		ID:   1,
		Name: "Pop",
	})
	db.Create(&model.Genre{
		ID:   2,
		Name: "Rock",
	})
	db.Create(&model.Genre{
		ID:   3,
		Name: "R&B",
	})
	db.Create(&model.Genre{
		ID:   4,
		Name: "Disco",
	})
	db.Create(&model.Genre{
		ID:   5,
		Name: "Soul",
	})
	db.Create(&model.Genre{
		ID:   6,
		Name: "Fung",
	})
	db.Create(&model.Genre{
		ID:   7,
		Name: "Jazz",
	})
	db.Create(&model.Genre{
		ID:   8,
		Name: "80-90s",
	})
	db.Create(&model.Genre{
		ID:   9,
		Name: "เพื่อชีวิต",
	})
	db.Create(&model.Genre{
		ID:   10,
		Name: "ลูกทุ่ง/ลูกกรุง",
	})
	db.Create(&model.Genre{
		ID:   11,
		Name: "Hiphop",
	})
	db.Create(&model.Genre{
		ID:   12,
		Name: "EDM",
	})
	db.Create(&model.Genre{
		ID:   13,
		Name: "Reggae/Ska",
	})
}

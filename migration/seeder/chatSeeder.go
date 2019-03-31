package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func ChatSeed(db *gorm.DB) {
	fmt.Println("seed chat ...")
	// db.Create(&model.Chat{
	// 	UserID:    2,
	// 	ToID:      1,
	// 	Message:   `สวัสดีครับ`,
	// 	Seen:      true,
	// 	CreatedAt: time.Now().AddDate(0, -2, 0).Add(10 * time.Hour).Add(30 * time.Minute),
	// })
	// db.Create(&model.Chat{
	// 	UserID:    2,
	// 	ToID:      1,
	// 	Message:   `สนใจเล่นงานแต่งไหมครับ`,
	// 	Seen:      true,
	// 	CreatedAt: time.Now().AddDate(0, -2, 0).Add(10 * time.Hour).Add(31 * time.Minute),
	// })
	// db.Create(&model.Chat{
	// 	UserID:    2,
	// 	ToID:      1,
	// 	Message:   `ราคา 15000 โรงแรม Ibis`,
	// 	Seen:      true,
	// 	CreatedAt: time.Now().AddDate(0, -2, 0).Add(10 * time.Hour).Add(32 * time.Minute),
	// })
	// db.Create(&model.Chat{
	// 	UserID:    1,
	// 	ToID:      2,
	// 	Message:   `OK ครับ`,
	// 	Seen:      true,
	// 	CreatedAt: time.Now().AddDate(0, -2, 0).Add(10 * time.Hour).Add(35 * time.Minute),
	// })
	// db.Create(&model.Chat{
	// 	UserID:    1,
	// 	ToID:      2,
	// 	Message:   `วันที่เท่าไหร่ครับ`,
	// 	Seen:      true,
	// 	CreatedAt: time.Now().AddDate(0, -2, 0).Add(10 * time.Hour).Add(36 * time.Minute),
	// })
	// db.Create(&model.Chat{
	// 	UserID:    2,
	// 	ToID:      1,
	// 	Message:   `16 ธันวา ครับ`,
	// 	Seen:      false,
	// 	CreatedAt: time.Now().AddDate(0, -2, 0).Add(10 * time.Hour).Add(40 * time.Minute),
	// })
	// db.Create(&model.Chat{
	// 	UserID:    2,
	// 	ToID:      3,
	// 	Message:   `สวัสดีครับ`,
	// 	Seen:      false,
	// 	CreatedAt: time.Now().AddDate(0, -3, 0).Add(10 * time.Hour).Add(40 * time.Minute),
	// })
}

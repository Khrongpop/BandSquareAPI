package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BandVideoSeed(db *gorm.DB) {
	fmt.Println("seed BandVideo ...")
	db.Create(&model.BandVideo{
		BandtypeID: 1,
		Video:      `https://www.youtube.com/watch?v=boU-Q9ml35I`,
		Code:       `boU-Q9ml35I`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 1,
		Video:      `https://www.youtube.com/watch?v=OSdKm2ll0vk`,
		Code:       `OSdKm2ll0vk`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 2,
		Video:      `https://www.youtube.com/watch?v=F20dqI7J9CI`,
		Code:       `F20dqI7J9CI`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 2,
		Video:      `https://www.youtube.com/watch?v=VgHsFpDjhwk`,
		Code:       `VgHsFpDjhwk`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 3,
		Video:      `https://www.youtube.com/watch?v=on59wzIc08Y`,
		Code:       `on59wzIc08Y`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 3,
		Video:      `https://www.youtube.com/watch?v=nHqZM1BmeHk`,
		Code:       `nHqZM1BmeHk`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 4,
		Video:      `https://www.youtube.com/watch?v=quNdO3TEfN0`,
		Code:       `quNdO3TEfN0`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 4,
		Video:      `https://www.youtube.com/watch?v=jam1aW4NAH4`,
		Code:       `jam1aW4NAH4`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 5,
		Video:      `https://www.youtube.com/watch?v=ANjZXe-95LA`,
		Code:       `ANjZXe-95LA`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 5,
		Video:      `https://www.youtube.com/watch?v=sM31H-olvD4`,
		Code:       `sM31H-olvD4`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 6,
		Video:      `https://www.youtube.com/watch?v=dhvrVXuENLM`,
		Code:       `dhvrVXuENLM`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 6,
		Video:      `https://www.youtube.com/watch?v=D0Nt3Op7n4M`,
		Code:       `D0Nt3Op7n4M`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 6,
		Video:      `https://www.youtube.com/watch?v=C4B6_qI2Ak4`,
		Code:       `C4B6_qI2Ak4`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 7,
		Video:      `https://www.youtube.com/watch?v=FY2bpvwbsQg`,
		Code:       `FY2bpvwbsQg`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 7,
		Video:      `https://www.youtube.com/watch?v=WFVxSaXjKbI`,
		Code:       `WFVxSaXjKbI`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 7,
		Video:      `https://www.youtube.com/watch?v=tjPhMDBcsBA`,
		Code:       `tjPhMDBcsBA`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 8,
		Video:      `https://www.youtube.com/watch?v=qVVjxu6d0HU`,
		Code:       `qVVjxu6d0HU`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 8,
		Video:      `https://www.youtube.com/watch?v=WkpCUQNQrog`,
		Code:       `WkpCUQNQrog`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 8,
		Video:      `https://www.youtube.com/watch?v=UJ_FB-8N_f0`,
		Code:       `UJ_FB-8N_f0`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 9,
		Video:      `https://www.youtube.com/watch?v=txkIqVQRVZI`,
		Code:       `txkIqVQRVZI`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 9,
		Video:      `https://www.youtube.com/watch?v=_TgUudFnZWk`,
		Code:       `_TgUudFnZWk`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 9,
		Video:      `https://www.youtube.com/watch?v=2Qe2k7fk504`,
		Code:       `2Qe2k7fk504`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 10,
		Video:      `https://www.youtube.com/watch?v=vOOKKtz6ilc`,
		Code:       `vOOKKtz6ilc`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 10,
		Video:      `https://www.youtube.com/watch?v=mhdODgztdAg`,
		Code:       `mhdODgztdAg`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 10,
		Video:      `https://www.youtube.com/watch?v=jA2l3uPQclc`,
		Code:       `jA2l3uPQclc`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 11,
		Video:      `https://www.youtube.com/watch?v=37BsmWQtQq4`,
		Code:       `37BsmWQtQq4`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 11,
		Video:      `https://www.youtube.com/watch?v=ANjKKf1Gp3A`,
		Code:       `ANjKKf1Gp3A`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 11,
		Video:      `https://www.youtube.com/watch?v=o804Aw9WxmQ`,
		Code:       `o804Aw9WxmQ`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 12,
		Video:      `https://www.youtube.com/watch?v=JFG02g_phaA`,
		Code:       `JFG02g_phaA`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 12,
		Video:      `https://www.youtube.com/watch?v=hcjtNGqw2ec`,
		Code:       `hcjtNGqw2ec`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 13,
		Video:      `https://www.youtube.com/watch?v=7A1bu5DpszQ`,
		Code:       `7A1bu5DpszQ`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 13,
		Video:      `https://www.youtube.com/watch?v=6GIsyFhchDo`,
		Code:       `6GIsyFhchDo`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 13,
		Video:      `https://www.youtube.com/watch?v=nLEIdLYHxYI`,
		Code:       `nLEIdLYHxYI`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 14,
		Video:      `https://www.youtube.com/watch?v=LlNk8g-DQLA`,
		Code:       `LlNk8g-DQLA`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 14,
		Video:      `https://www.youtube.com/watch?v=1gCPV1xTL70`,
		Code:       `1gCPV1xTL70`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 14,
		Video:      `https://www.youtube.com/watch?v=bfURENUQc7U`,
		Code:       `bfURENUQc7U`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 15,
		Video:      `https://www.youtube.com/watch?v=qI2LERva23c`,
		Code:       `qI2LERva23c`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 16,
		Video:      `https://www.youtube.com/watch?v=u4VOOcCh5w4`,
		Code:       `u4VOOcCh5w4`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 16,
		Video:      `https://www.youtube.com/watch?v=syufhzvnqsc`,
		Code:       `syufhzvnqsc`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 16,
		Video:      `https://www.youtube.com/watch?v=YUNHHPYcDtI`,
		Code:       `YUNHHPYcDtI`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 17,
		Video:      `https://www.youtube.com/watch?v=i3S7a7IZWKQ`,
		Code:       `i3S7a7IZWKQ`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 17,
		Video:      `https://www.youtube.com/watch?v=ijswb1vIZCk`,
		Code:       `ijswb1vIZCk`,
	})
	db.Create(&model.BandVideo{
		BandtypeID: 17,
		Video:      `https://www.youtube.com/watch?v=KEBV-XxUIHk`,
		Code:       `KEBV-XxUIHk`,
	})
}

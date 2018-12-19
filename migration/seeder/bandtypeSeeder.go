package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BandTypeSeed(db *gorm.DB) {
	fmt.Println("seed category ...")
	db.Create(&model.BandType{
		BandID: 1,
		TypeID: 1,
		Detail: "วงดนตรีเซ็ตเล็กฟังสบาย ประกอบด้วยนักร้องหญิง กีตาร์ ดับเบิ้ลเบส แซ็กโซโฟน (สามารถเพิ่มกลองไฟฟ้า หรือคีย์บอร์ดได้)",
	})
	db.Create(&model.BandType{
		BandID: 1,
		TypeID: 1,
		Detail: "",
	})
	db.Create(&model.BandType{
		BandID: 1,
		TypeID: 1,
		Detail: "",
	})
	db.Create(&model.BandType{
		BandID: 1,
		TypeID: 1,
		Detail: "",
	})
}

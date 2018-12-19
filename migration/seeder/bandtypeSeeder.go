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
		BandID: 2,
		TypeID: 1,
		Detail: `เป็นวงดนตรีที่มีตั้งแต่ 2 ชิ้นขึ้นไป (กีตาร์+ร้อง) และในกรณี 3 ชิ้นขึ้นไป ท่านสามารถเลือกได้ว่า "ท่านต้องการได้อะไร"​ (กลอง,แซ๊กโซโฟน หรืออื่นๆ) อีกทั้งท่านยังสามารถคัดเลือกนักร้องชายหญิงได้ตามที่ท่านต้องการ`,
	})
	db.Create(&model.BandType{
		BandID: 2,
		TypeID: 2,
		Detail: `-สามารถเพิ่มจำนวนชิ้นที่มีตั้งแต่ 4 คนขึ้นไป / ท่านสามารถคัดเลือกนักร้องได้ด้วยตัวของท่านเอง โดยจะมีทั้งหมด 3คนด้วยกัน (ไฮท์,เติ้ลและแบ๊ด) / สามารถเพิ่มคียบอร์ดหรือแดนซ์เซอร์ได้ / รีเควสแนวดนตรีที่ท่านชอบได้`,
	})
	// db.Create(&model.BandType{
	// 	BandID: 1,
	// 	TypeID: 1,
	// 	Detail: ``,
	// })
}

package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BandSeed(db *gorm.DB) {
	fmt.Println("seed band ...")
	// #####################   1  ########################
	db.Create(&model.Band{
		ID:           1,
		UserID:       1,
		About:        "Darin วงดนตรีที่เหมาะสำหรับงานเลี้ยงส่วนตัว งานแต่งงานสุดโรแมนติก เพื่อเพิ่มบรรยากาศที่อบอุ่น หรูหรา ให้แก่งานของคุณผ่านเสียงดนตรี วงสามารถเล่นเพลงได้หลากหลายแนวทั้ง Pop, R&B/Soul, Jazz, Funk รวมถึงเพลงลูกทุ่งก็เช่นกัน  สร้างความแตกต่างให้งานของคุณด้วยเสียงร้องสุดหวานของพลอย นักร้องนำ พร้อมด้วยกีตาร์ ดับเบิ้ลเบส และแซกโซโฟน (รวมถึงสามารถเพิ่มคียบอร์ด และกลองไฟฟ้าตามความต้องการของลูกค้าได้)",
		Member:       "พลอย - นักร้อง จาก I Can See Your Voice Thailand, ร้องแลกไลค์ไทยแลนด์ '\n' นักดนตรี - Back-Up ฟลุ๊ค The Star, อิมเมจ The Voice, ฯลฯ'",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     10000,
		MaxPrice:     30000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/540b0ce1-b68b-42d1-a500-1ca7233abf3f.jpeg",
	})

	// #####################   2  ########################
	db.Create(&model.Band{
		ID:           2,
		UserID:       3,
		About:        "พวกเราวง BIPOLOR '\n' - สามารถรับงานได้ตั้งแต่ Folk Song 2-3 ชิ้น จนไปถึง งาน Full Band '\n' - สามารถเล่นได้ทุกแนว ตั้งแต่เพลงตลาดไทยเก่า/ใหม่ แนววาไรตี้ รวมถึงเพลงสากล '\n' - หนึ่งในสมาชิกในวงเป็นนักแสดง (IG:tape_worrachai) '\n' - สามารถเพิ่มสมาชิกให้มีนักร้องหญิงได้",
		Member:       "วรชัย ศิริคงสุวรรณ (เทป) - กลอง '\n' พลกฤษณ์ อยู่บัว (ไฮท์) - ร้องนำ '\n' มนัส สุขเกษม (โรเบิด) - กีต้าร์ '\n' สามารถเพิ่มสมาชิกเครื่องอื่นๆเช่น เบส, คีย์บอร์ด หรือเพิ่มนักร้องหญิงได้ตามที่ท่านต้องการ",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     10000,
		MaxPrice:     20000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/c547bf66-0a45-4d7d-b9f3-fada75713e53.jpeg",
	})

	// #####################   21  ########################
	// db.Create(&model.Band{
	// 	About:        "",
	// 	Member:       "",
	// 	WorkLocation: "",
	// 	MinPrice:     1000,
	// 	MaxPrice:     200,
	// 	Cover:        "",
	// })

	// // #####################   21  ########################
	// db.Create(&model.Band{
	// 	About:        "",
	// 	Member:       "",
	// 	WorkLocation: "",
	// 	MinPrice:     1000,
	// 	MaxPrice:     200,
	// 	Cover:        "",
	// })

}

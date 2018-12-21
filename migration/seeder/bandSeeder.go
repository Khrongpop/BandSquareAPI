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
		Member:       "พลอย - นักร้อง จาก I Can See Your Voice Thailand, ร้องแลกไลค์ไทยแลนด์ '\n'นักดนตรี - Back-Up ฟลุ๊ค The Star, อิมเมจ The Voice, ฯลฯ'",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     10000,
		MaxPrice:     30000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/540b0ce1-b68b-42d1-a500-1ca7233abf3f.jpeg",
	})

	// #####################   2  ########################
	db.Create(&model.Band{
		ID:           2,
		UserID:       3,
		About:        "พวกเราวง BIPOLOR \n - สามารถรับงานได้ตั้งแต่ Folk Song 2-3 ชิ้น จนไปถึง งาน Full Band \n - สามารถเล่นได้ทุกแนว ตั้งแต่เพลงตลาดไทยเก่า/ใหม่ แนววาไรตี้ รวมถึงเพลงสากล \n - หนึ่งในสมาชิกในวงเป็นนักแสดง (IG:tape_worrachai) \n - สามารถเพิ่มสมาชิกให้มีนักร้องหญิงได้",
		Member:       "วรชัย ศิริคงสุวรรณ (เทป) - กลอง \nพลกฤษณ์ อยู่บัว (ไฮท์) - ร้องนำ \nมนัส สุขเกษม (โรเบิด) - กีต้าร์ \nสามารถเพิ่มสมาชิกเครื่องอื่นๆเช่น เบส, คีย์บอร์ด หรือเพิ่มนักร้องหญิงได้ตามที่ท่านต้องการ",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     10000,
		MaxPrice:     20000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/c547bf66-0a45-4d7d-b9f3-fada75713e53.jpeg",
	})

	// #####################   3  ########################
	db.Create(&model.Band{
		ID:           3,
		UserID:       4,
		About:        "วง Oneday เป็นกลุ่มนักศึกษาจากวิทยาลัยดุริยางคศิลป์ ม.มหิดล ที่มากด้วยฝีมือ และประสบการณ์ สามารถเล่นเพลงได้หลากหลายแนวมากๆ ตั้งแต่ Pop, Jazz, Rock,B     ไปจนถึงลูกทุ่ง สามารถรับได้ตั้งแต่งานแต่งงาน งานเลี้ยง งานอีเวนท์ งานปาร์ตี้",
		Member:       "นายธรรมศาสตร์ หนูชัยแก้ว - เปียโน คีย์บอร์ด \nนายศิวัช ไม้ค้าง - กีต้าร์ไฟฟ้า กีต้าร์โปร่ง \nนายปณิธาน คชสวัสดิ์ - เบสไฟฟ้า \nนายศรายุทธ ลิขิตปัทมสิงห์ - กลองชุด คาฮอง \nนายศรราม คามบุตร - แซกโซโฟน \nนางสาวณัฐชยา ราชวงษ์ - นักร้องหญิง \nนายธนพล ศรีรัตน์ - นักร้องชาย",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     9000,
		MaxPrice:     20000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/4eddcbd9-58fd-4a75-95f8-be6e1dc2a9ca.jpeg",
	})

	// #####################   4  ########################
	db.Create(&model.Band{
		ID:           4,
		UserID:       5,
		About:        "Big Tan เป็นวงดนตรีแนว Pop/R&B/Jazz ที่รวมตัวกันจากชุมนุม TU Folksong จากมหาวิทยาลัยธรรมศาสตร์ \n พวกเราได้นำเพลงฮิตติดหูทุกยุคสมัย ทั้งไทยและสากลมา Arrange ใหม่ในแบบที่ฟังสบายหู มี Groove ที่ชวนโยกตาม และรวมตัวนักดนตรีที่มากประสบการณ์ พร้อมที่จะให้บริการดนตรีสำหรับงานแต่งงานและอีเว้นต์ต่างๆ หากทางลูกค้าต้องการเครื่องดนตรีเฉพาะเพิ่ม อย่างเช่นเครื่องเป่า ทางเราก็สามารถจัดหานักดนตรีอาชีพมาเสริมได้",
		Member:       "ตาล - ร้องนำ \nโจ - กีตาร์ \nเต้ย - เบส \nดิว - คีย์บอร์ด, คอรัส \nเจ - กลอง \nอัพ - กีตาร์ \nเจมี่ - Saxophone",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     8000,
		MaxPrice:     30000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/bc36b679-6c6a-44ba-af51-3f23343fbb97.jpeg",
	})

	// #####################   5  ########################
	db.Create(&model.Band{
		ID:           5,
		UserID:       6,
		About:        "สำราญใจ วงดนตรีหลากหลายรูปแบบ ทั้งอคูสติกและฟูลแบนด์ ทั้งซึ้ง ไพเราะ และมันส์สนุกสุดเหวี่ยง รวมถึงเพลงยุคเก่า 80-90 ที่คุณเคยประทับใจและนึกไม่ถึง เราจะเอามาเล่นให้คุณฟัง ด้วยประสบการณ์เล่นงานแต่งและอีเวนท์มาตลอดระยะเวลากว่า 10 ปี \n  **ทางวงมีนักร้องให้เลือกหลากหลายคน แต่ละคนล้วนมีประสบการณ์ผ่านการประกวดทั้ง The Voice Thailand หรือ ผ่านรายการ I Can See Your Voice มาทั้งสิ้น",
		Member:       "ปราชญ์ - กีตาร์ & เบส \nกลม อรวี (The Voice Thailand Season 1) - ร้องนำ \nลี่ - กลอง \nต้น - ร้องนำ \nสามารถหานักร้องหญิงได้ถ้าลูกค้าต้องการ",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     10000,
		MaxPrice:     30000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/e6833d22-6a54-4b18-9e42-7d7ea2b7e289.png",
	})

	// #####################   6  ########################
	db.Create(&model.Band{
		ID:           6,
		UserID:       7,
		About:        "พวกเราวง Zerox จากเมืองพัทยา จังหวัดชลบุรี พร้อมที่จะสร้างความสุขแก่ท่านผู้ฟังและท่านผู้ชมด้วยเสียงเพลงเสียงดนตรี เพลงที่พวกเราเน้นเล่นสดจะเป็นแนว วาไรตี้ ลูกทุ่ง สตริง Pop Rock หรือ ตามความต้องการของลูกค้า",
		Member:       "แพน - ร้องนำ \nทอมโบ้ - กีตาร์ \nทาม - คีย์บอร์ด \nปอ - เบส \nเซน - กลอง",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     8000,
		MaxPrice:     10000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/0bb10aed-f95e-4973-96e4-4369665b43f6.jpeg",
	})

	// #####################   7  ########################
	db.Create(&model.Band{
		ID:           7,
		UserID:       8,
		About:        "รวบรวมศิลปินที่ผ่านงานมามากมาย ทั้งแบ็คอัพศิลปินดัง เบิร์ด ธงชัย,เป๊ก ผลิตโชคและร่วมร้องกับศิลปินดังมากมายหลายซิงเกิล ผ่านเวทีประกวดมามากมายทั้งThe voice,KPN,First stage",
		Member:       "ป่าน - นักร้องหญิง \nปอ - กีตาร์/ร้องนำ \nพุทธ - คาฮอง/ร้องนำ \nฮุย - saxsophone \nตั้ว - คีย์บอร์ด/ร้องนำ \nปอ - กีตาร์ \nวิว - กลอง \nเบียร์ - กีตาร์/ร้องนำ\n นก - ร้องนำ",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     10000,
		MaxPrice:     40000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/a2baaa55-1e5f-41a2-a4f8-f0d775f717b4.jpeg",
	})

	// #####################   8  ########################
	db.Create(&model.Band{
		ID:           8,
		UserID:       9,
		About:        "Middleway เป็นวงดนตรีแนว Rock/Numetal มีสมาชิกทั้งหมด 4 คน โดยวง Middleway สามารถเล่นดนตรีได้หลากหลายแนว เช่น ป็อป, เพื่อชีวิต, วาไรตี้สนุกๆ, ร็อคไทยเก่ายันใหม่ รวมถึงเพลงสากล ที่จะเน้นไปทาง Rock ช่วงยุค 80’s ขึ้นไป \n ที่มาของชื่อวง Middleway ก็คือ Middle = กลาง , Way = ทาง, ถนน รวมกันเป็น ทางสายกลาง ซึ่งจะบ่งบอกถึงความเป็นกลางของวงว่า Middleway สามารถเล่นดนตรีได้หลากหลายแนวนั่นเอง",
		Member:       "เบิ้ม - ร้องนำ \nวสันต์ - กีต้าร์ \nกอล์ฟ - เบส \nบอล - กลอง",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     6000,
		MaxPrice:     10000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/dc676617-7744-4ece-a642-66c296a2fb48.jpeg",
	})

	// #####################   9  ########################
	db.Create(&model.Band{
		ID:           9,
		UserID:       10,
		About:        "ด้วยความที่หลงไหลในเสียงเพลงเเละการเเสดงออกโดยใช้ดนตรีทำให้ตัดสินใจเป็นดีเจในปลายปี 2014 โดยเริ่มจากการฝึกฝนโดยพี่ชายที่รู้จักกันในขอนแก่นจนสามารถที่จะสร้างความมันส์ให้เกิดกับทุกๆปาร์ตี้ในงานต่างๆได้ด้วยเพลงเเละบีทที่มีจังหวะสนุกๆเเละชวนให้ทุกคนได้ขยับตัวไปกับบทเพลง \n เเนวเพลงที่ถนัด \n \n EDM ///// HOUSE MUSIC (esp. Future House) ///// HIPHOP-R&B ///// TRAP ///// \n \n -ผมอยากจะสร้างสรรค์บทเพลงให้ทุกงานมีความสนุกและร่วมสร้างความทรงจำดีๆให้กับทุกๆคน-",
		Member:       "DJ Bright",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     5000,
		MaxPrice:     10000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/35391212-4f78-40a8-96cd-9dd9725f912b.jpeg",
	})

	// #####################   10  ########################
	db.Create(&model.Band{
		ID:           10,
		UserID:       11,
		About:        "DJ Aristo เริ่มเข้าสู่วงการดีเจ ปี2010 ผ่านการเรียนโดยใช้หลักkey mixing แนวเพลงที่ถนัด ตั้งแต่ยุค 90' s จนถึง EDM ที่กำลังเป็นที่นิยมในปัจจุบัน \n \n Type : Hiphop / Tech House / Basshouse / Trap / Twerk / Bigroom / Groove 90 \n \n ติดต่อขอฟัง Demo / mixtape เพิ่มเติมได้ \n \n Mixcloud : https://www.mixcloud.com/arthur-papan-paper/full-set/ \n \n **สามารถติดต่อเครื่องเสียงเพิ่มได้ถ้าลูกค้าต้องการ",
		Member:       "Arth (DJ Aristo)",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     5000,
		MaxPrice:     10000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7aeae8e1-4760-4c73-875a-84e33715dd2f.jpeg",
	})

	// #####################   11  ########################
	db.Create(&model.Band{
		ID:           11,
		UserID:       12,
		About:        "Hexthree the reggae dub band from nonthaburi,Thailand contain with 5 instruments (Drumset,Bass,Keyboards,Trombone,Saxophone) with Live Dub mixing on board.  We start with Love & deep interest in reggae dub music and the feeling of creator to spread the positive vibes through our imagination to everybody. Our independent band do it ourself in all method : Produce,Lyrics,Arrangement,Record,Mixing and Mastering ; All our creation is for spread our feeling & imagination through our loved music. \n Hexthree (เฮ็กซ์-ทรี) วงดนตรีอิสระ แนวเพลง reggae dub จากจังหวัดนนทบุรี,ประเทศไทย ประกอบด้วยเครื่องดนตรี 5 ชิ้น (กลองชุด, เบส, คีย์บอร์ด, ทรอมโบน, แซกโซโฟน) ผสมผสานกับการมิกซ์เสียงแบบ Dub สดๆ \n พวกเราเริ่มต้นจาก ความรักและคลั่งไคล้ในแนวดนตรีเร้กเก้ดั๊บ และความต้องการในการทำเพลงที่ส่งความรู้สึกด้านบวกผ่านจินตนาการของพวกเราแก่ทุกคน  วงดนตรีอิสระของพวกเราทำกันเองในทุกกระบวนการทั้ง โปรดิ๊วซ์,เขียนเพลง,เรียบเรียง,การอัด,การมิกซ์ และ การมาสเตอร์ริ่ง รวมถึงการทำ Artwork CD ตัดต่อ Video ทำM V  ทุกๆการสร้างสรรค์ทั้งหมด เพื่อ ถ่ายทอดความรู้สึกของพวกเราผ่านเนื้อหา และ จินตนาการทางเสียงดนตรีที่พวกเราชื่นชอบ",
		Member:       "1. จิรภัทร กลิ่นอุทัย - ร้องนำ,เบส \n2. จันท์จุฑา วิมลศิลปิน - ร้องคอรัส,กลอง \n3. ชลน ทองมี - กีตาร์ \n4. วิชญะ วงศ์สกุล - ทรอมโบน \n5. ราชนาวี สนธิสวน - แซกโซโฟน \n6. ภูวเนศ ชุมจันทร์ - Dubman",
		WorkLocation: "กรุงเทพมหานคร",
		MinPrice:     7000,
		MaxPrice:     20000,
		Cover:        "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/6c08c037-a537-4aa6-aa01-a121684eae5b.jpeg",
	})

}

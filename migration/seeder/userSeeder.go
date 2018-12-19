package seeder

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
	"golang.org/x/crypto/bcrypt"
)

func UserSeed(db *gorm.DB) {
	t := time.Now()
	t.Format("2006-01-02 15:04:05")
	// #####################   1 ########################
	db.Create(&model.User{
		Name:      "Darin",
		Email:     "darin@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/85f3122e-83fd-4dc5-ac5f-3131f9e7b259.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/85f3122e-83fd-4dc5-ac5f-3131f9e7b259.jpeg",
		Active:    0,
		RoleID:    11,
	})

	// #####################   11 ########################
	db.Create(&model.User{
		Name:      "Muang",
		Email:     "muang@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://graph.facebook.com/2343960062310644/picture?type=large&return_ssl_resources=1",
		Thumbnail: "https://graph.facebook.com/2343960062310644/picture?type=large&return_ssl_resources=1",
		Active:    0,
		RoleID:    1,
	})

	// #####################   21 ########################
	db.Create(&model.User{
		Name:      "BIPOLOR",
		Email:     "bipolor@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/9881bf96-3d36-4491-97c6-2a0ee78ac4cb.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/9881bf96-3d36-4491-97c6-2a0ee78ac4cb.jpeg",
		Active:    0,
		RoleID:    11,
	})

	// #####################   31 ########################
	db.Create(&model.User{
		Name:      "Oneday",
		Email:     "oneday@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/5e55560e-ef11-483f-82fc-2dd7483a4d6b.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/5e55560e-ef11-483f-82fc-2dd7483a4d6b.jpeg",
		Active:    1,
		RoleID:    11,
	})

	// #####################   41 ########################
	db.Create(&model.User{
		Name:      "Big Tan",
		Email:     "bigtan@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/913d4942-0547-430c-a758-4f5d53ee9b9c.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/913d4942-0547-430c-a758-4f5d53ee9b9c.jpeg",
		Active:    0,
		RoleID:    11,
	})

	// #####################   51 ########################
	db.Create(&model.User{
		Name:      "สำราญใจ",
		Email:     "samran@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/25197a16-c77a-4abe-94e3-7cc8cca81798.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/25197a16-c77a-4abe-94e3-7cc8cca81798.jpeg",
		Active:    1,
		RoleID:    11,
	})

	// #####################   61 ########################
	db.Create(&model.User{
		Name:      "Zerox",
		Email:     "zerox@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/9d2e3f65-79f3-421a-bb8b-83ded63d8710.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/9d2e3f65-79f3-421a-bb8b-83ded63d8710.jpeg",
		Active:    1,
		RoleID:    11,
	})

	// #####################   71 ########################
	db.Create(&model.User{
		Name:      "Mouth Smell Milk Music Group",
		Email:     "msmmg@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/f2881fb5-686f-45bc-9eee-96ec1486cba6.png",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/f2881fb5-686f-45bc-9eee-96ec1486cba6.png",
		Active:    0,
		RoleID:    11,
	})

	// #####################   81 ########################
	db.Create(&model.User{
		Name:      "DJ Bright",
		Email:     "dj_bright@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/bca29532-fdc7-482f-815e-32f73bf07863.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/bca29532-fdc7-482f-815e-32f73bf07863.jpeg",
		Active:    1,
		RoleID:    11,
	})

	// #####################   91 ########################
	db.Create(&model.User{
		Name:      "DJ ARiSTO",
		Email:     "dj_aristo@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7aeae8e1-4760-4c73-875a-84e33715dd2f.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7aeae8e1-4760-4c73-875a-84e33715dd2f.jpeg",
		Active:    0,
		RoleID:    11,
	})

	// #####################   101 ########################
	db.Create(&model.User{
		Name:      "HEXTHREE",
		Email:     "HEXTHREE@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/301436fc-2520-479e-b464-f35ddc81d49a.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/301436fc-2520-479e-b464-f35ddc81d49a.jpeg",
		Active:    0,
		RoleID:    11,
	})

	// #####################   111 ########################
	db.Create(&model.User{
		Name:      "Arm",
		Email:     "arm@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2Farm.jpg?alt=media&token=4cbca5a3-3245-4860-9f5d-dec27e277718",
		Thumbnail: "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2Farm.jpg?alt=media&token=4cbca5a3-3245-4860-9f5d-dec27e277718",
		Active:    0,
		RoleID:    11,
	})

	// #####################   121 ########################
	db.Create(&model.User{
		Name:      "Bank",
		Email:     "bank@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2Fbank.jpg?alt=media&token=35569c4e-7358-4cf4-9e2b-089641452d25",
		Thumbnail: "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2Fbank.jpg?alt=media&token=35569c4e-7358-4cf4-9e2b-089641452d25",
		Active:    0,
		RoleID:    11,
	})

	// #####################   131 ########################
	db.Create(&model.User{
		Name:      "PuenXXX",
		Email:     "puen@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FPeun.jpg?alt=media&token=63f73b67-f7c7-42a8-b868-cf3b9ff0f62d",
		Thumbnail: "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FPeun.jpg?alt=media&token=63f73b67-f7c7-42a8-b868-cf3b9ff0f62d",
		Active:    0,
		RoleID:    11,
	})

	// #####################   141 ########################
	db.Create(&model.User{
		Name:      "TubWanXXX",
		Email:     "tubtub@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FTubwan.jpg?alt=media&token=c1b96b63-8f1e-4d63-ae70-a979fa977564",
		Thumbnail: "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FTubwan.jpg?alt=media&token=c1b96b63-8f1e-4d63-ae70-a979fa977564",
		Active:    0,
		RoleID:    11,
	})

	// #####################   151 ########################
	db.Create(&model.User{
		Name:      "Tle",
		Email:     "tle@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FTle.jpg?alt=media&token=9aab747c-8fba-4f54-b8bc-71e794789cc8",
		Thumbnail: "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FTle.jpg?alt=media&token=9aab747c-8fba-4f54-b8bc-71e794789cc8",
		Active:    0,
		RoleID:    11,
	})

	// #####################   161 ########################
	db.Create(&model.User{
		Name:      "Boss",
		Email:     "boss@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FBoss.jpg?alt=media&token=02d3c05f-ebf0-43de-ba4f-717588c320f6",
		Thumbnail: "https://firebasestorage.googleapis.com/v0/b/thesis-4ef45.appspot.com/o/image%2Fusers%2FBoss.jpg?alt=media&token=02d3c05f-ebf0-43de-ba4f-717588c320f6",
		Active:    0,
		RoleID:    11,
	})

}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

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
	db.Create(&model.User{
		Name:      "Darin",
		Email:     "darin@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/85f3122e-83fd-4dc5-ac5f-3131f9e7b259.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/85f3122e-83fd-4dc5-ac5f-3131f9e7b259.jpeg",
		RoleID:    11,
	})

	db.Create(&model.User{
		Name:      "Muang",
		Email:     "muang@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://graph.facebook.com/2343960062310644/picture?type=large&return_ssl_resources=1",
		Thumbnail: "https://graph.facebook.com/2343960062310644/picture?type=large&return_ssl_resources=1",
		RoleID:    1,
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

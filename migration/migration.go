package migration

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/migration/seeder"
	"github.com/khrongpop/BandSquareAPI/model"
)

var db *gorm.DB

func DropTable(db *gorm.DB) {
	db.DropTable(&model.BandType{}, &model.Band{}, &model.User{})
}

func DBSetup(db *gorm.DB) {
	// db.DropTable(&model.Band{}, &model.User{}, &model.Role{})
	db.AutoMigrate(&model.User{}, &model.SocailAccount{}, &model.Band{}, &model.BandType{}, &model.Booking{}, &model.Review{})
	// dataSeed(db)
}

func dataSeed(db *gorm.DB) {
	// t := time.Now()
	// t.Format("2006-01-02 15:04:05")

	// seeder.RoleSeed(db)
	// seeder.CategorySeed(db)
	// seeder.GenreSeed(db)
	// seeder.TypeSeed(db)
	// seeder.UserSeed(db)
	// seeder.BandSeed(db)
	// seeder.BandTypeSeed(db)
	// seeder.BandGenreSeed(db)
	// seeder.BookingSeed(db)
	seeder.ReviewSeed(db)
	db.Model(&model.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "RESTRICT")
	db.Model(&model.SocailAccount{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Band{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.BandType{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.BandType{}).AddForeignKey("type_id", "types(id)", "cascade", "RESTRICT")
	db.Model(&model.BandGenre{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.BandGenre{}).AddForeignKey("genre_id", "genres(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("category_id", "categories(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("type_id", "types(id)", "cascade", "RESTRICT")
	db.Model(&model.Review{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Review{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.Review{}).AddForeignKey("booking_id", "bookings(id)", "cascade", "RESTRICT")
}

func RefreshDB(db *gorm.DB) {
	fmt.Println("drop table ...")
	// db.DropTable(&model.BandType{}, &model.Band{}, &model.User{}, &model.Category{}, &model.Genre{}, &model.Type{}, &model.Role{})
	// db.DropTable(&model.BandType{}, &model.Band{}, &model.User{})
	// db.DropTable(&model.Booking{})
	fmt.Println("migrate table ...")
	db.AutoMigrate(&model.User{}, &model.SocailAccount{}, &model.Band{}, &model.BandType{}, &model.Booking{}, &model.Review{})
	dataSeed(db)
	fmt.Println("migrate complete !!!")
}

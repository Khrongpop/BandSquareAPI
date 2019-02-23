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
	seeder.UserSeed(db)
	seeder.BandSeed(db)
	seeder.BandCategorySeed(db)
	seeder.BandTypeSeed(db)
	seeder.BandImageSeed(db)
	seeder.BandVideoSeed(db)
	seeder.BandGenreSeed(db)
	seeder.BookingSeed(db)
	seeder.ReviewSeed(db)
	seeder.FavouriteSeed(db)
	seeder.ChatSeed(db)
	seeder.BookingBandSeed(db)
	seeder.BookingGenreSeed(db)
	addForeignKey(db)
}

func addForeignKey(db *gorm.DB) {
	db.Model(&model.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "RESTRICT")
	db.Model(&model.SocailAccount{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Band{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.BandCategory{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.BandCategory{}).AddForeignKey("category_id", "categories(id)", "cascade", "RESTRICT")
	db.Model(&model.BandType{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.BandType{}).AddForeignKey("type_id", "types(id)", "cascade", "RESTRICT")
	db.Model(&model.BandImage{}).AddForeignKey("bandtype_id", "band_types(id)", "cascade", "RESTRICT")
	db.Model(&model.BandVideo{}).AddForeignKey("bandtype_id", "band_types(id)", "cascade", "RESTRICT")
	db.Model(&model.BandGenre{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.BandGenre{}).AddForeignKey("genre_id", "genres(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("category_id", "categories(id)", "cascade", "RESTRICT")
	db.Model(&model.Booking{}).AddForeignKey("type_id", "types(id)", "cascade", "RESTRICT")
	db.Model(&model.BookingBand{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.BookingBand{}).AddForeignKey("booking_id", "bookings(id)", "cascade", "RESTRICT")
	db.Model(&model.BookingGenre{}).AddForeignKey("genre_id", "genres(id)", "cascade", "RESTRICT")
	db.Model(&model.BookingGenre{}).AddForeignKey("booking_id", "bookings(id)", "cascade", "RESTRICT")
	db.Model(&model.Review{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Review{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.Review{}).AddForeignKey("booking_id", "bookings(id)", "cascade", "RESTRICT")
	db.Model(&model.Favourite{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Favourite{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.Chat{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Chat{}).AddForeignKey("to_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.Notification{}).AddForeignKey("booking_id", "bookings(id)", "cascade", "RESTRICT")
	db.Model(&model.Notification{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")

}

func RefreshDB(db *gorm.DB) {
	fmt.Println("drop table ...")

	// DropTable to user
	// db.DropTable(&model.Notification{}, &model.Chat{}, &model.Favourite{}, &model.Review{}, &model.BookingBand{}, &model.BookingGenre{}, &model.Booking{}, &model.BandImage{}, &model.BandVideo{}, &model.BandGenre{}, &model.BandCategory{}, &model.BandType{}, &model.Band{}, &model.SocailAccount{}, &model.User{})

	// db.DropTable(&model.Notification{})
	// db.DropTable(&model.Chat{})
	// db.DropTable(&model.Favourite{})
	// db.DropTable(&model.Review{})
	// db.DropTable(&model.BookingGenre{})
	// db.DropTable(&model.BookingBand{})
	// db.DropTable(&model.Booking{})
	// db.DropTable(&model.BandImage{})
	// db.DropTable(&model.BandVideo{})
	// db.DropTable(&model.BandGenre{})
	// db.DropTable(&model.BandCategory{})
	// db.DropTable(&model.BandType{})
	// db.DropTable(&model.Band{})
	// db.DropTable(&model.SocailAccount{})
	// db.DropTable(&model.User{})

	fmt.Println("migrate table ...")

	// AutoMigrate to user
	// db.AutoMigrate(&model.User{}, &model.SocailAccount{}, &model.Band{}, &model.BandType{}, &model.BandCategory{}, &model.BandGenre{}, &model.BandVideo{}, &model.BandImage{}, &model.Booking{}, &model.BookingBand{}, &model.BookingGenre{}, &model.Review{}, &model.Favourite{}, &model.Chat{}, &model.Notification{})

	// db.AutoMigrate(&model.User{})
	// db.AutoMigrate(&model.SocailAccount{})
	// db.AutoMigrate(&model.Band{})
	// db.AutoMigrate(&model.BandType{})
	// db.AutoMigrate(&model.BandCategory{})
	// db.AutoMigrate(&model.BandGenre{})
	// db.AutoMigrate(&model.BandVideo{})
	// db.AutoMigrate(&model.BandImage{})
	// db.AutoMigrate(&model.Booking{})
	// db.AutoMigrate(&model.BookingBand{})
	// db.AutoMigrate(&model.BookingGenre{})
	// db.AutoMigrate(&model.Review{})
	// db.AutoMigrate(&model.Favourite{})
	// db.AutoMigrate(&model.Chat{})
	db.AutoMigrate(&model.Notification{})

	fmt.Println("seed table ...")
	// dataSeed(db)
	fmt.Println("migrate complete !!!")
}

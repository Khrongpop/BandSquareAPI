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
	db.AutoMigrate(&model.Role{}, &model.User{}, &model.Band{}, &model.BandType{})
	dataSeed(db)
}

func dataSeed(db *gorm.DB) {
	// t := time.Now()
	// t.Format("2006-01-02 15:04:05")

	// seeder.RoleSeed(db)git
	// seeder.CategorySeed(db)
	// seeder.GenreSeed(db)
	// seeder.TypeSeed(db)
	// seeder.UserSeed(db)
	// seeder.BandSeed(db)
	seeder.BandTypeSeed(db)
	// db.Model(&model.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "RESTRICT")
	// db.Model(&model.Band{}).AddForeignKey("user_id", "users(id)", "cascade", "RESTRICT")
	db.Model(&model.BandType{}).AddForeignKey("band_id", "bands(id)", "cascade", "RESTRICT")
	db.Model(&model.BandType{}).AddForeignKey("type_id", "types(id)", "cascade", "RESTRICT")
}

func RefreshDB(db *gorm.DB) {
	fmt.Println("drop table ...")
	// db.DropTable(&model.BandType{}, &model.Band{}, &model.User{}, &model.Category{}, &model.Genre{}, &model.Type{}, &model.Role{})
	db.DropTable(&model.BandType{})
	fmt.Println("migrate table ...")
	db.AutoMigrate(&model.Role{}, &model.User{}, &model.Band{}, &model.BandType{})
	dataSeed(db)
}

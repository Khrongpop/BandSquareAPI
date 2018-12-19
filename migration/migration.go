package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/migration/seeder"
	"github.com/khrongpop/BandSquareAPI/model"
)

var db *gorm.DB

func DropTable(db *gorm.DB) {
	db.DropTable(&model.User{}, &model.Role{})
}

func DBSetup(db *gorm.DB) {
	db.AutoMigrate(&model.Role{}, &model.User{})
	dataSeed(db)
}

func dataSeed(db *gorm.DB) {
	// t := time.Now()
	// t.Format("2006-01-02 15:04:05")
	seeder.RoleSeed(db)
	seeder.UserSeed(db)
	db.Model(&model.User{}).AddForeignKey("role_id", "roles(id)", "cascade", "RESTRICT")
}

func RefreshDB(db *gorm.DB) {
	DropTable(db)
	db.AutoMigrate(&model.Role{}, &model.User{})
	dataSeed(db)
}

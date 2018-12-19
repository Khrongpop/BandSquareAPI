package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
	"github.com/labstack/echo"
)

var db *gorm.DB
var err error

func List(db *gorm.DB, c echo.Context) error {
	users := []model.User{}
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

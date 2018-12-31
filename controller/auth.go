package controller

import (
	"net/http"

	"github.com/khrongpop/BandSquareAPI/model"
	echo "github.com/labstack/echo"
)

func login(c echo.Context) error {
	var user model.User
	// if err := c.Bind(&user); err != nil {
	// 	return err
	// }
	// db.Create(&user)
	// name := c.FormValue("name")
	return c.JSON(http.StatusOK, user)
}

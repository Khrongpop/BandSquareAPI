package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var db *gorm.DB
var err error

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:created_at`
	Updated_at string `json:updated_at`
}

func main() {
	db, err = gorm.Open("mysql", "muang:muang@tcp(35.247.180.6)/bandsquare")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/create", create)
	e.GET("/users", list)
	e.GET("/users/:id", view)
	e.PUT("/users/:id", update)
	e.DELETE("/users/:id", delete)
	e.Logger.Fatal(e.Start(":1323"))
}

func create(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	db.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func list(c echo.Context) error {
	users := []User{}
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func view(c echo.Context) error {
	var user User
	db.First(&user, c.Param("id"))
	return c.JSON(http.StatusOK, user)
}

func update(c echo.Context) error {
	var user User
	db.First(&user, c.Param("id"))
	name := c.FormValue("name")
	db.Model(&user).Update("Name", name)
	return c.JSON(http.StatusOK, user)
}

func delete(c echo.Context) error {
	var user User
	db.First(&user, c.Param("id"))
	db.Delete(&user)
	return c.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/khrongpop/BandSquareAPI/migration"
	"github.com/khrongpop/BandSquareAPI/model"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB
var err error

func main() {
	viper.AutomaticEnv()
	// port := ":" + viper.GetString("port")
	port := ":1323"
	datasource := viper.GetString("CLEARDB_DATABASE_URL")

	mysqlUser := "b85b02f8218929"
	mysqlPass := "1642c1e7"
	mysqlHost := "us-cdbr-iron-east-01.cleardb.net"
	mysqlName := "heroku_a5a40c45511bb84"
	mysqlConfig := "?parseTime=true"

	databaseURL := fmt.Sprintf("%v:%v@tcp(%v)/%v%v", mysqlUser, mysqlPass, mysqlHost, mysqlName, mysqlConfig)
	db, err = gorm.Open("mysql", databaseURL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// migration.DBSetup(db)
	// fmt.Println("pass")
	// pass := "123456"
	// passLog := hashAndSalt([]byte(pass))
	// check := comparePasswords(passLog, []byte(pass))
	// check2 := comparePasswords(passLog, []byte("1234"))

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! test"+datasource)
	})

	e.POST("/create", create)
	e.GET("/users", list)
	e.GET("/users/:id", view)
	e.PUT("/users/:id", update)
	e.DELETE("/users/:id", delete)

	e.GET("/bands", bandslist)
	e.GET("/db/refesh", refreshDB)

	e.Logger.Fatal(e.Start(port))

}

func create(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	db.Create(&user)
	// name := c.FormValue("name")
	return c.JSON(http.StatusOK, user)
}

func list(c echo.Context) error {
	users := []model.User{}
	db.Find(&users)
	for i, _ := range users {
		db.Model(&users[i]).Related(&users[i].Role)
		db.Model(&users[i]).Related(&users[i].Band)
		// users[i].Band = band
	}
	return c.JSON(http.StatusOK, users)
}

func bandslist(c echo.Context) error {
	bands := []model.Band{}
	db.Find(&bands)
	for i, _ := range bands {
		// db.Model(&users[i]).Related(&users[i].Role)
		// db.Model(&users[i]).Related(&users[i].Band)
		// users[i].Band = band
		db.Model(&bands[i]).Related(&bands[i].Types)
		// for i, _ := range bands[i].Types {
		// db.Model(&bands[i].Types[i]).Related(&bands[i].Types[i].Type)
		// }
	}
	return c.JSON(http.StatusOK, bands)
}

func view(c echo.Context) error {
	var user model.User
	db.First(&user, c.Param("id"))
	db.Model(&user).Related(&user.Role)
	db.Model(&user).Related(&user.Band)

	return c.JSON(http.StatusOK, user)
}

func update(c echo.Context) error {
	var user model.User
	db.First(&user, c.Param("id"))
	name := c.FormValue("name")
	db.Model(&user).Update("Name", name)
	return c.JSON(http.StatusOK, user)
}

func delete(c echo.Context) error {
	var user model.User
	db.First(&user, c.Param("id"))
	db.Delete(&user)
	return c.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
}

func refreshDB(c echo.Context) error {
	migration.RefreshDB(db)
	return c.JSON(http.StatusOK, echo.Map{
		"result": "refres-success",
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

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

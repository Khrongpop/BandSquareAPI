package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
	RoleID    uint   `json:role_id`
	Role      Role
}
type Role struct {
	gorm.Model
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func main() {
	viper.AutomaticEnv()
	port := ":" + viper.GetString("port")
	// port := ":1323"
	datasource := viper.GetString("CLEARDB_DATABASE_URL")

	mysqlUser := "b85b02f8218929"
	mysqlPass := "1642c1e7"
	mysqlHost := "us-cdbr-iron-east-01.cleardb.net"
	mysqlName := "heroku_a5a40c45511bb84"

	DATABASE_URL := fmt.Sprintf("%v:%v@tcp(%v)/%v", mysqlUser, mysqlPass, mysqlHost, mysqlName)
	// DATABASE_URL := "b85b02f8218929:1642c1e7@tcp(us-cdbr-iron-east-01.cleardb.net)/heroku_a5a40c45511bb84"
	db, err = gorm.Open("mysql", DATABASE_URL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// fmt.Println("pass")
	// pass := "123456"
	// passLog := hashAndSalt([]byte(pass))
	// check := comparePasswords(passLog, []byte(pass))
	// check2 := comparePasswords(passLog, []byte("1234"))

	dbSetup()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! test"+datasource)
	})

	e.POST("/create", create)
	e.GET("/users", list)
	e.GET("/users/:id", view)
	e.PUT("/users/:id", update)
	e.DELETE("/users/:id", delete)
	e.Logger.Fatal(e.Start(port))

}

func create(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	// db.Create(&user)
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

func dbSetup() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	// dataSeed()
}
func dataSeed() {
	// t := time.Now()
	// t.Format("2006-01-02 15:04:05")
	userSeed()
	roleSeed()
	db.Model(&User{}).AddForeignKey("role_id", "roles(id)", "cascade", "RESTRICT")
}

func userSeed() {

	db.Create(&User{
		Name:      "Darin",
		Email:     "darin@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/85f3122e-83fd-4dc5-ac5f-3131f9e7b259.jpeg",
		Thumbnail: "https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/85f3122e-83fd-4dc5-ac5f-3131f9e7b259.jpeg",
		RoleID:    11,
	})

	db.Create(&User{
		Name:      "Muang",
		Email:     "muang@gmail.com",
		Password:  hashAndSalt([]byte("123456")),
		Image:     "https://graph.facebook.com/2343960062310644/picture?type=large&return_ssl_resources=1",
		Thumbnail: "https://graph.facebook.com/2343960062310644/picture?type=large&return_ssl_resources=1",
		RoleID:    1,
	})
}

func roleSeed() {
	db.Create(&Role{
		Name: "Member",
		Slug: "Member",
	})
	db.Create(&Role{
		Name: "Musician",
		Slug: "Music",
	})
	db.Create(&Role{
		Name: "Administrator",
		Slug: "Admin",
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

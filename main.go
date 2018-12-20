package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

	// input := "2017-08-31"
	// layout := "2006-01-02"
	// t, _ := time.Parse(layout, input)
	// fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
	// fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017

	viper.AutomaticEnv()
	port := ":" + viper.GetString("port")
	// port := ":1323"
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
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! test"+datasource)
	})

	auth := e.Group("/auth")
	auth.POST("/login", login)
	auth.POST("/register", register)
	auth.POST("/fblogin", fblogin)

	e.POST("/create", create)
	e.GET("/users", list)
	e.GET("/users/:id", view)
	e.PUT("/users/:id", update)
	e.DELETE("/users/:id", delete)

	bands := e.Group("/band")
	bands.GET("/bands", bandslist)
	bands.POST("/recommend", bandslist)
	bands.GET("/onlines", bandOnline)
	bands.POST("/news", bandOnline)
	bands.GET("/cheaps", bandCheap)
	bands.GET("/detail/:id", bandDetail)

	e.GET("/db/refesh", refreshDB)

	e.Logger.Fatal(e.Start(port))

}

func login(c echo.Context) error {
	var user model.User
	name := c.FormValue("name")
	if err := db.Where("name = ?", name).Or("email = ?", name).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, "invalid user")
	}
	password := c.FormValue("password")
	check := comparePasswords(user.Password, []byte(password))
	if check {
		db.Model(&user).Related(&user.Role)
		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusOK, "invalid password")
	}
}

func register(c echo.Context) error {
	var user model.User
	name := c.FormValue("name")
	if err := db.Where("name = ?", name).Or("email = ?", name).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		user.Name = name
		user.Email = c.FormValue("email")
		user.Active = true
		user.Password = hashAndSalt([]byte(c.FormValue("password")))
		user.Image = c.FormValue("image")
		user.Thumbnail = c.FormValue("image")
		user.RoleID = 1
		db.Create(&user)
		db.Model(&user).Related(&user.Role)
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusOK, "already have an user")
}

func fblogin(c echo.Context) error {
	var fb model.SocailAccount
	var user model.User
	providerID := c.FormValue("id")

	if err := db.Where("provider_id = ?", providerID).First(&fb).Error; gorm.IsRecordNotFoundError(err) {
		name := c.FormValue("name")
		email := c.FormValue("email")
		if err := db.Where("email = ?", email).First(&user).Error; gorm.IsRecordNotFoundError(err) {
			user.Name = name
			user.Email = email
			user.Active = true
			user.Password = hashAndSalt([]byte(strconv.Itoa(rand.Intn(100))))
			user.Image = "https://graph.facebook.com/" + providerID + "/picture?type=large&return_ssl_resources=1"
			user.Thumbnail = "https://graph.facebook.com/" + providerID + "/picture"
			user.RoleID = 1
			db.Create(&user)
			db.Model(&user).Related(&user.Role)
			fb.UserID = user.ID
			fb.ProviderID = providerID
			fb.Provider = c.FormValue("provider")
			db.Create(&fb)
			return c.JSON(http.StatusOK, user)
		}
		db.Model(&user).Related(&user.Role)
		fb.UserID = user.ID
		fb.ProviderID = providerID
		fb.Provider = c.FormValue("provider")
		db.Create(&fb)
		return c.JSON(http.StatusOK, user)
	}
	if err := db.First(&user, fb.UserID).Error; gorm.IsRecordNotFoundError(err) {
		log.Fatal(err)
	}
	db.Model(&user).Related(&user.Role)
	return c.JSON(http.StatusOK, user)
}

func bandOnline(c echo.Context) error {
	var bands []model.Band
	db.Joins("JOIN users ON bands.user_id = users.id AND users.active = ?", 1).Find(&bands)
	for i, band := range bands {
		var user model.User
		db.Model(&user).Related(&user.Role)
		db.First(&user, band.UserID)
		bands[i].User = &user
	}
	return c.JSON(http.StatusOK, bands)
}

func bandCheap(c echo.Context) error {
	var bands []model.Band
	db.Joins("JOIN users ON bands.user_id = users.id AND bands.max_price < ?", 25000).Find(&bands)
	for i, band := range bands {
		var user model.User
		db.First(&user, band.UserID)
		db.Model(&user).Related(&user.Role)
		bands[i].User = &user
	}
	return c.JSON(http.StatusOK, bands)
}

func bandDetail(c echo.Context) error {
	var band model.Band
	var user model.User
	var bandType []model.BandType

	db.First(&band, c.Param("id"))
	db.First(&user, band.UserID)
	band.User = &user
	if err := db.Where("band_id = ?", band.ID).Find(&bandType).Error; gorm.IsRecordNotFoundError(err) {

	}

	categoriesList := ""
	for i := range bandType {
		db.Model(&bandType[i]).Related(&bandType[i].Type)
		categoriesList += bandType[i].Type.Name
		if i != len(bandType)-1 {
			categoriesList += " , "
		}
	}
	band.Types = bandType
	band.CategoryList = &categoriesList

	db.Model(&band).Related(&band.Genres, "Genres")
	genresList := model.GetGenreList(band)
	band.GenreList = &genresList

	db.Model(&band).Related(&band.Bookings)
	for i := range band.Bookings {
		db.Model(&band.Bookings[i]).Related(&band.Bookings[i].User)
		db.Model(&band.Bookings[i]).Related(&band.Bookings[i].Category)
		db.Model(&band.Bookings[i]).Related(&band.Bookings[i].Type)
	}

	db.Model(&band).Related(&band.Reviews)
	rateAvg := model.GetRateAVG(band.Reviews)
	band.RateAvg = &rateAvg
	return c.JSON(http.StatusOK, band)
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
	for i := range users {
		db.Model(&users[i]).Related(&users[i].Role)
		// db.Model(&users[i]).Related(&users[i].Band)
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
	// db.Model(&user).Related(&user.Band)

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

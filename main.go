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

	// mysqlUser := viper.GetString("MYSQLUSER")
	// mysqlPass := viper.GetString("MYSQLPASS")
	// mysqlHost := viper.GetString("MYSQLHOST")
	// mysqlName := viper.GetString("MYSQLName")
	mysqlConfig := "charset=utf8&parseTime=true"

	databaseURL := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", mysqlUser, mysqlPass, mysqlHost, mysqlName, mysqlConfig)
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
	auth.POST("/enable", enableUser)
	auth.POST("/disable", disableUser)

	bands := e.Group("/band")
	bands.GET("/bands", bandslist)
	bands.GET("/recommend", bandRecommend)
	bands.GET("/detail/:id", testbandDetail)
	bands.POST("/recommend", bandRecommend)
	bands.POST("/onlines", bandOnline)
	bands.POST("/news", bandNew)
	bands.POST("/cheaps", bandCheap)
	bands.POST("/detail", bandDetail)
	bands.POST("/info", bandInfo)
	bands.POST("/types", bandTypes)
	bands.POST("/bookings", bandBookings)
	bands.POST("/reviews", bandReviews)
	bands.POST("/favourite", favourite)
	bands.POST("/favourite_check", checkFavourite)

	chats := e.Group("/chat")
	chats.GET("/testgetchats", testgetChats)
	chats.GET("/testgetchatuser/:uID/:tID", testgetChatUser)

	bands.GET("/info/:id", testbandInfo)
	bands.GET("/types/:id", testbandTypes)
	bands.GET("/bookings/:id", testbandBookings)
	bands.GET("/reviews/:id", testbandReviews)

	e.GET("/db/refesh", refreshDB)

	e.Logger.Fatal(e.Start(port))

}

func login(c echo.Context) error {
	user := model.User{}
	name := c.FormValue("name")
	if err := db.Where("name = ?", name).Or("email = ?", name).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, "invalid user")
	}
	password := c.FormValue("password")
	check := comparePasswords(user.Password, []byte(password))
	if check {
		db.Model(&user).Related(&user.Role)
		var favourites []model.Band
		db.Joins("JOIN favourites ON favourites.band_id = bands.id AND favourites.user_id = ?", user.ID).Find(&favourites)
		for i, band := range favourites {
			var userBand model.User
			db.First(&userBand, band.UserID)
			favourites[i].User = &userBand
		}
		user.Favourites = favourites

		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusOK, "invalid password")
	}
}

func register(c echo.Context) error {
	user := model.User{}
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
	fb := model.SocailAccount{}
	user := model.User{}
	providerID := c.FormValue("id")

	if err := db.First(&fb, "provider_id = ?", providerID).Error; gorm.IsRecordNotFoundError(err) {
		name := c.FormValue("name")
		email := c.FormValue("email")
		if err := db.First(&user, "email = ?", email).Error; gorm.IsRecordNotFoundError(err) {
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

func enableUser(c echo.Context) error {
	user := model.User{}
	db.First(&user, c.FormValue("id"))
	db.Model(&user).Update("active", true)
	return c.JSON(http.StatusOK, `enable_user_sucsees`)
}

func disableUser(c echo.Context) error {
	user := model.User{}
	db.First(&user, c.FormValue("id"))
	db.Model(&user).Update("active", false)
	return c.JSON(http.StatusOK, `disable_user_sucsees`)
}

func favourite(c echo.Context) error {
	favourite := model.Favourite{}
	response := Response{}
	userID := c.FormValue("user_id")
	bandID := c.FormValue("band_id")
	UserID, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	BandID, err := strconv.ParseUint(bandID, 10, 32)
	if err := db.First(&favourite, `user_id = ? AND band_id = ?`, userID, bandID).Error; gorm.IsRecordNotFoundError(err) {
		db.Create(&model.Favourite{
			UserID: uint(UserID),
			BandID: uint(BandID),
		})
		response.Message = `favourite sucsses`
		return c.JSON(http.StatusOK, response)
	}
	db.Delete(&favourite)
	response.Message = `unfavourite sucsses`
	return c.JSON(http.StatusOK, response)
}

func checkFavourite(c echo.Context) error {
	favourite := model.Favourite{}
	response := model.FavouriteCheck{}
	if err = db.First(&favourite, `user_id = ? AND band_id = ?`, c.FormValue("user_id"), c.FormValue("band_id")).Error; gorm.IsRecordNotFoundError(err) {
		response.Status = false
		return c.JSON(http.StatusOK, response)
	}
	response.Status = true
	return c.JSON(http.StatusOK, response)
}

func bandRecommend(c echo.Context) error {
	bands := []model.Band{}
	db.Table("bands").Select("* ,AVG(reviews.rate) as avg").Joins("JOIN reviews ON reviews.band_id = bands.id ").Group("bands.user_id").Order("avg desc").Scan(&bands)
	for i := range bands {
		bands[i] = getBandTitle(bands[i])
		// bands[i] = getBandDetail(bands[i])
	}

	return c.JSON(http.StatusOK, bands)
}

func bandOnline(c echo.Context) error {
	bands := []model.Band{}
	db.Joins("JOIN users ON bands.user_id = users.id AND users.active = ?", 1).Find(&bands)
	for i := range bands {
		bands[i] = getBandTitle(bands[i])
		// bands[i] = getBandDetail(bands[i])
	}
	return c.JSON(http.StatusOK, bands)
}

func bandNew(c echo.Context) error {
	bands := []model.Band{}
	db.Table("bands").Select("* ").Joins("left join bookings ON bookings.band_id = bands.id ").Where("bookings.id IS NULL").Group("bands.user_id").Scan(&bands)
	for i := range bands {
		bands[i] = getBandTitle(bands[i])
		// bands[i] = getBandDetail(bands[i])
	}
	return c.JSON(http.StatusOK, bands)
}

func bandCheap(c echo.Context) error {
	bands := []model.Band{}
	db.Find(&bands, "max_price < ?", 15000)
	for i := range bands {
		bands[i] = getBandTitle(bands[i])
		// bands[i] = getBandDetail(bands[i])
	}
	return c.JSON(http.StatusOK, bands)
}

func bandDetail(c echo.Context) error {
	band := model.Band{}
	// db.First(&band, c.Param("id"))
	db.First(&band, c.FormValue("band_id"))

	band = getBandTitle(band)
	band = getBandDetail(band)

	return c.JSON(http.StatusOK, band)
}

func bandInfo(c echo.Context) error {
	band := model.Band{}
	db.First(&band, c.FormValue("band_id"))

	band = getBandTitle(band)
	// band = getBandDetail(band)

	return c.JSON(http.StatusOK, band)
}

func bandTypes(c echo.Context) error {
	bandType := []model.BandType{}
	if err := db.Find(&bandType, "band_id = ?", c.FormValue("band_id")).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range bandType {
		db.Model(&bandType[i]).Related(&bandType[i].Type)
		var images []model.BandImage
		var videos []model.BandVideo
		db.Find(&images, `bandtype_id = ?`, bandType[i].ID)
		db.Find(&videos, `bandtype_id = ?`, bandType[i].ID)
		bandType[i].Images = images
		bandType[i].Videos = videos
	}
	return c.JSON(http.StatusOK, bandType)
}

func bandBookings(c echo.Context) error {
	bookings := []model.Booking{}
	if err := db.Find(&bookings, "band_id = ?", c.FormValue("band_id")).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range bookings {
		db.Model(&bookings[i]).Related(&bookings[i].User)
		db.Model(&bookings[i]).Related(&bookings[i].Category)
		db.Model(&bookings[i]).Related(&bookings[i].Type)
	}
	return c.JSON(http.StatusOK, bookings)
}
func bandReviews(c echo.Context) error {
	reviews := []model.Review{}
	if err := db.Find(&reviews, "band_id = ?", c.FormValue("band_id")).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range reviews {
		var user model.User
		var booking model.Booking
		db.Find(&user, reviews[i].UserID)
		db.Find(&booking, reviews[i].BookingID)
		db.Model(&booking).Related(&booking.Category)
		db.Model(&booking).Related(&booking.Type)
		reviews[i].User = &user
		reviews[i].Booking = &booking
	}
	return c.JSON(http.StatusOK, reviews)
}

func testbandDetail(c echo.Context) error {
	band := model.Band{}
	db.First(&band, c.Param("id"))

	band = getBandTitle(band)
	band = getBandDetail(band)

	return c.JSON(http.StatusOK, band)
}

func testbandInfo(c echo.Context) error {
	band := model.Band{}
	db.First(&band, c.Param("id"))

	band = getBandTitle(band)
	// band = getBandDetail(band)

	return c.JSON(http.StatusOK, band)
}

func testbandTypes(c echo.Context) error {
	bandType := []model.BandType{}
	if err := db.Find(&bandType, "band_id = ?", c.Param("id")).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range bandType {
		db.Model(&bandType[i]).Related(&bandType[i].Type)
		images := []model.BandImage{}
		videos := []model.BandVideo{}
		db.Find(&images, `bandtype_id = ?`, bandType[i].ID)
		db.Find(&videos, `bandtype_id = ?`, bandType[i].ID)
		bandType[i].Images = images
		bandType[i].Videos = videos
	}
	return c.JSON(http.StatusOK, bandType)
}

func testbandBookings(c echo.Context) error {
	bookings := []model.Booking{}
	if err := db.Find(&bookings, "band_id = ?", c.Param("id")).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range bookings {
		db.Model(&bookings[i]).Related(&bookings[i].User)
		db.Model(&bookings[i]).Related(&bookings[i].Category)
		db.Model(&bookings[i]).Related(&bookings[i].Type)
	}
	return c.JSON(http.StatusOK, bookings)
}
func testbandReviews(c echo.Context) error {
	reviews := []model.Review{}
	if err := db.Find(&reviews, "band_id = ?", c.Param("id")).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range reviews {
		user := model.User{}
		booking := model.Booking{}
		db.Find(&user, reviews[i].UserID)
		db.Find(&booking, reviews[i].BookingID)
		db.Model(&booking).Related(&booking.Category)
		db.Model(&booking).Related(&booking.Type)
		reviews[i].User = &user
		reviews[i].Booking = &booking
	}
	return c.JSON(http.StatusOK, reviews)
}

func testgetChats(c echo.Context) error {
	chats := []model.Chat{}
	db.Find(&chats)
	for i := range chats {
		db.Model(&chats[i]).Related(&chats[i].User)
		db.Model(&chats[i]).Related(&chats[i].ToUser, "ToID")
	}
	return c.JSON(http.StatusOK, chats)
}

func testgetChatUser(c echo.Context) error {
	chats := []model.Chat{}
	if err := db.Where(`user_id = ? AND to_id = ?`, c.Param(`uID`), c.Param(`tID`)).Or(`to_id = ? AND user_id = ?`, c.Param(`uID`), c.Param(`tID`)).Find(&chats).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range chats {
		db.Model(&chats[i]).Related(&chats[i].User)
		db.Model(&chats[i]).Related(&chats[i].ToUser, "ToID")
	}
	return c.JSON(http.StatusOK, chats)
}

func getBandTitle(band model.Band) model.Band {
	var user model.User
	db.First(&user, band.UserID)
	db.Model(&user).Related(&user.Role)
	band.User = &user

	db.Model(&band).Related(&band.Reviews)
	rateAvg := model.GetRateAVG(band.Reviews)
	band.RateAvg = &rateAvg

	db.Model(&band).Related(&band.Categories, "categories")
	categoriesList := model.GetCategoryList(band)
	band.CategoryList = &categoriesList

	db.Model(&band).Related(&band.Genres, "genres")
	genresList := model.GetGenreList(band)
	band.GenreList = &genresList

	return band
}

func getBandDetail(band model.Band) model.Band {
	var bandType []model.BandType
	if err := db.Find(&bandType, "band_id = ?", band.ID).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range bandType {
		db.Model(&bandType[i]).Related(&bandType[i].Type)
		var images []model.BandImage
		var videos []model.BandVideo
		db.Find(&images, `bandtype_id = ?`, bandType[i].ID)
		db.Find(&videos, `bandtype_id = ?`, bandType[i].ID)
		bandType[i].Images = images
		bandType[i].Videos = videos
	}
	band.Types = bandType

	db.Model(&band).Related(&band.Bookings)
	for i := range band.Bookings {
		db.Model(&band.Bookings[i]).Related(&band.Bookings[i].User)
		db.Model(&band.Bookings[i]).Related(&band.Bookings[i].Category)
		db.Model(&band.Bookings[i]).Related(&band.Bookings[i].Type)
	}

	for i, review := range band.Reviews {
		var user model.User
		var booking model.Booking
		db.Find(&user, review.UserID)
		db.Find(&booking, review.BookingID)
		db.Model(&booking).Related(&booking.Category)
		db.Model(&booking).Related(&booking.Type)
		band.Reviews[i].User = &user
		band.Reviews[i].Booking = &booking

	}
	return band
}

func create(c echo.Context) error {
	user := model.User{}
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

type Response struct {
	Message string `json:"message"`
}

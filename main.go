package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/khrongpop/BandSquareAPI/migration"
	"github.com/khrongpop/BandSquareAPI/model"
	"github.com/khrongpop/BandSquareAPI/notification"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB
var err error

func main() {

	// input := "2019-04-21 05:34:25"
	// layout := "2019-04-21"
	// t, _ := time.Parse(layout, input)

	// input := c.FormValue(`date`)
	// layout := "2006-01-02"

	// input := "2017-08-31"
	// layout := "2006-01-02 15:04:05"
	// t, _ := time.Parse(layout, input)
	// fmt.Println(t)                               // 2017-08-31 00:00:00 +0000 UTC
	// fmt.Println(t.Format("2006-01-02 15:04:05")) // 31-Aug-2017

	viper.AutomaticEnv()
	port := ":" + viper.GetString("PORT")

	mysqlUser := viper.GetString("MYSQLUSER")
	mysqlPass := viper.GetString("MYSQLPASS")
	mysqlHost := viper.GetString("MYSQLHOST")
	mysqlName := viper.GetString("MYSQLNAME")
	mysqlConfig := "charset=utf8&parseTime=true"

	databaseURL := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", mysqlUser, mysqlPass, mysqlHost, mysqlName, mysqlConfig)
	db, err = gorm.Open("mysql", databaseURL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// migration.DBSetup(db)
	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))

	// appID := viper.GetString("OSAPPID")
	// APIKEY := viper.GetString("OSAPIKEY")
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello Word!!!!!",
			// "appID":   appID,
			// "APIKEY":  APIKEY,
		})
	})

	auth := e.Group("/auth")
	auth.POST("/login", login)
	auth.POST("/register", register)
	auth.POST("/fblogin", fblogin)
	auth.POST("/enable", enableUser)
	auth.POST("/disable", disableUser)
	auth.POST("/favourite", getFavourite)
	auth.POST("/band_register", bandRegister)
	auth.POST("/current_band", getCurrentBand)
	auth.POST("/get_currentwork", getCurrentWork)
	auth.POST("/get_work_thisday", getWorkThisDay)
	auth.POST("/get_overview", getOverView)

	bands := e.Group("/band")
	bands.GET("/bands", bandslist)
	bands.GET("/recommend", bandOnline)
	bands.GET("/onlines", bandOnline)
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
	bands.POST("/bands_categories", bandCategories)
	bands.GET("/bands_categories/:id", bandCategories)

	chats := e.Group("/chat")
	chats.GET("/testgetchats/:uID", testgetChats)
	chats.GET("/testgetchatuser/:uID/:tID", testgetChatUser)
	chats.POST("/getchatuser", getChatUser)
	chats.POST("/getchats", getChat)
	chats.POST("/store", storeChat)

	bookings := e.Group("/booking")
	bookings.POST("/current_booking", getCurrentBooking)
	bookings.POST("/current_booking_band", getCurrentBookingBand)
	bookings.POST("/quick_booking", quickBook)
	bookings.POST("/booking", booking)
	bookings.POST("/select_band", selectBandBooking)
	bookings.POST("/band_accept", bandAcceptBooking)
	bookings.POST("/band_discard", bandDiscardtBooking)
	bookings.POST("/payment", paymentBandBooking)
	bookings.POST("/confirm", confirmBooking)
	bookings.POST("/review", reviewBooking)
	bookings.GET("/testcurbooking/:id", getCurrentBookingBand)

	notifications := e.Group("/notification")
	notifications.POST("/get_user_noti", getNotification)
	notifications.POST("/seen_noti", readNotification)
	notifications.POST("/add_player_id", addPlayerID)
	notifications.POST("/remove_player_id", removePlayerID)

	admin := e.Group("/admin")
	admin.POST("/login", adminLogin)
	admin.POST("/get_user", getUsers)
	admin.POST("/get_works", getWorks)

	// TEST API
	bands.GET("/info/:id", testbandInfo)
	bands.GET("/types/:id", testbandTypes)
	bands.GET("/bookings/:id", testbandBookings)
	bands.GET("/reviews/:id", testbandReviews)

	// e.GET("/favourite", getFavourite)

	e.GET("/db/refesh", refreshDB)
	e.GET("/testnoti", testNoti)

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
		favourites := []model.Band{}
		db.Joins("JOIN favourites ON favourites.band_id = bands.id AND favourites.user_id = ?", user.ID).Find(&favourites)
		for i, band := range favourites {
			userBand := model.User{}
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
	favourites := []model.Band{}
	db.Joins("JOIN favourites ON favourites.band_id = bands.id AND favourites.user_id = ?", user.ID).Find(&favourites)
	for i, band := range favourites {
		userBand := model.User{}
		db.First(&userBand, band.UserID)
		favourites[i].User = &userBand
	}
	user.Favourites = favourites
	return c.JSON(http.StatusOK, user)
}

func enableUser(c echo.Context) error {
	user := model.User{}
	db.Model(&user).Where("id = ?", c.FormValue("id")).Update("active", true)
	return c.JSON(http.StatusOK, `enable_user_sucsees`)
}

func disableUser(c echo.Context) error {
	user := model.User{}
	db.Model(&user).Where("id = ?", c.FormValue("id")).Update("active", false)
	return c.JSON(http.StatusOK, `disable_user_sucsees`)
}

func getCurrentBand(c echo.Context) error {
	band := model.Band{}
	db.First(&band, `user_id = ?`, c.FormValue("user_id"))
	db.Model(&band).Related(&band.Reviews)
	rateAvg := model.GetRateAVG(band.Reviews)
	band.RateAvg = &rateAvg
	return c.JSON(http.StatusOK, band)
}

func getFavourite(c echo.Context) error {
	userID := c.FormValue("user_id")
	favourites := []model.Band{}
	db.Joins("JOIN favourites ON favourites.band_id = bands.id AND favourites.user_id = ?", userID).Find(&favourites)
	for i, band := range favourites {
		userBand := model.User{}
		db.First(&userBand, band.UserID)
		favourites[i].User = &userBand
	}
	return c.JSON(http.StatusOK, favourites)
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

func getCurrentWork(c echo.Context) error {
	works := []model.Booking{}
	band := model.Band{}
	db.First(&band, "user_id = ? ", c.FormValue(`user_id`))
	db.Find(&works, "band_id = ?", band.ID)
	for i := range works {
		db.Model(&works[i]).Related(&works[i].User)
		db.Model(&works[i]).Related(&works[i].Category)
		db.Model(&works[i]).Related(&works[i].Type)
	}
	return c.JSON(http.StatusOK, works)
}

func getOverView(c echo.Context) error {
	bookings := []model.Booking{}
	band := model.Band{}
	reviews := []model.Review{}
	favourites := []model.Favourite{}
	complete := 0
	working := 0

	db.First(&band, `user_id = ?`, c.FormValue(`user_id`))
	db.Find(&bookings, `band_id = ?`, band.ID)
	for _, booking := range bookings {
		if booking.Status == 4 {
			complete++
		} else if booking.Status == 3 {
			working++
		}

	}
	if err := db.Find(&reviews, "band_id = ?", band.ID).Error; gorm.IsRecordNotFoundError(err) {

	}
	if err := db.Find(&favourites, "band_id = ?", band.ID).Error; gorm.IsRecordNotFoundError(err) {

	}

	return c.JSON(http.StatusOK, echo.Map{
		`complete`:   complete,
		`working`:    working,
		`review`:     len(reviews),
		`favourites`: len(favourites),
	})
}

func getWorkThisDay(c echo.Context) error {
	work := model.Booking{}
	band := model.Band{}
	dateFormat := time.Now().Format("2006-01-02")

	db.First(&band, "user_id = ? ", c.FormValue(`user_id`))
	if err := db.First(&work, "band_id = ? AND date = ?", band.ID, dateFormat).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "not have an work",
			"status":  404,
		})
	}

	db.Model(&work).Related(&work.User)
	db.Model(&work).Related(&work.Category)
	db.Model(&work).Related(&work.Type)

	return c.JSON(http.StatusOK, work)
}

func getNotification(c echo.Context) error {
	notifitions := []model.Notification{}

	db.Order("created_at desc").Find(&notifitions, `user_id = ?`, c.FormValue(`user_id`))
	for i := range notifitions {
		db.Model(&notifitions[i]).Related(&notifitions[i].User)
		booking := model.Booking{}
		db.First(&booking, `id = ?`, notifitions[i].BookingID)
		db.Model(&booking).Related(&booking.User)
		db.Model(&booking).Related(&booking.Category)
		db.Model(&booking).Related(&booking.Type)
		notifitions[i].Booking = &booking
	}
	return c.JSON(http.StatusOK, notifitions)
}

func addPlayerID(c echo.Context) error {

	player := model.PlayerID{}
	if err := db.First(&player, `user_id = ? AND player_id = ?`, c.FormValue(`user_id`), c.FormValue(`player_id`)).Error; gorm.IsRecordNotFoundError(err) {

		userID, err := strconv.ParseUint(c.FormValue(`user_id`), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		player.PlayerID = c.FormValue(`player_id`)
		player.UserID = uint(userID)
		db.Create(&player)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "add player suscess",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "already have player",
	})
}

func removePlayerID(c echo.Context) error {
	player := model.PlayerID{}
	if err := db.First(&player, `user_id = ? AND player_id = ?`, c.FormValue(`user_id`), c.FormValue(`player_id`)).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "not found player ",
		})
	}
	db.Delete(&player)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "remove player suscess",
	})
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
	db.Joins("JOIN users ON bands.user_id = users.id AND users.active = ?", 1).Limit(5).Find(&bands)
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

func bandCategories(c echo.Context) error {
	bands := []model.Band{}
	db.Joins("JOIN band_categories on bands.id = band_categories.band_id AND band_categories.category_id = ?", c.FormValue(`cat_id`)).Find(&bands)
	for i := range bands {
		bands[i] = getBandTitle(bands[i])
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
	if err := db.Order("created_at desc").Find(&bookings, "band_id = ?", c.FormValue("band_id")).Error; gorm.IsRecordNotFoundError(err) {

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
	if err := db.Order("created_at desc").Find(&reviews, "band_id = ?", c.FormValue("band_id")).Error; gorm.IsRecordNotFoundError(err) {

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

func getChat(c echo.Context) error {
	chats := []model.Chat{}
	userChats := []model.Chat{}
	if err := db.Where(`user_id = ? `, c.FormValue(`uID`)).Or(`to_id = ? `, c.FormValue(`uID`)).Order("created_at desc").Find(&chats).Error; gorm.IsRecordNotFoundError(err) {

	}
	for _, chat := range chats {

		if len(userChats) > 0 {
			valCheck := 0
			for _, user := range userChats {

				if user.UserID == chat.UserID && user.ToID == chat.ToID {
					valCheck++
				} else if user.UserID == chat.ToID && user.ToID == chat.UserID {
					valCheck++
				}
			}

			if valCheck == 0 {
				db.Model(&chat).Related(&chat.User)
				db.Model(&chat).Related(&chat.ToUser, "ToID")
				userChats = append(userChats, chat)
			}
		} else {
			db.Model(&chat).Related(&chat.User)
			db.Model(&chat).Related(&chat.ToUser, "ToID")
			userChats = append(userChats, chat)
		}

	}
	return c.JSON(http.StatusOK, userChats)
}
func getChatUser(c echo.Context) error {
	chats := []model.Chat{}
	userID, err := strconv.ParseInt(c.FormValue(`uID`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	if err := db.Where(`user_id = ? AND to_id = ?`, c.FormValue(`uID`), c.FormValue(`tID`)).Or(`to_id = ? AND user_id = ?`, c.FormValue(`uID`), c.FormValue(`tID`)).Find(&chats).Error; gorm.IsRecordNotFoundError(err) {

	}
	for i := range chats {
		if chats[i].ToID == int(userID) && chats[i].Seen == false {
			db.Model(&chats[i]).Update("Seen", true)
		}
		db.Model(&chats[i]).Related(&chats[i].User)
		db.Model(&chats[i]).Related(&chats[i].ToUser, "ToID")
	}
	return c.JSON(http.StatusOK, chats)
}

func storeChat(c echo.Context) error {
	chat := model.Chat{}

	userID, err := strconv.ParseInt(c.FormValue(`user_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	toID, err := strconv.ParseInt(c.FormValue(`to_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	chat.UserID = int(userID)
	chat.ToID = int(toID)
	chat.Message = c.FormValue(`message`)
	db.Create(&chat)

	// Create norification
	players := []model.PlayerID{}
	db.Find(&players, `user_id = ?`, chat.ToID)
	data := `{
		"page": "chat",
		"to_id": "` + c.FormValue(`user_id`) + `"
	}`
	user := model.User{}
	db.First(&user, chat.ToID)
	message := user.Name + `: ` + c.FormValue(`message`)
	notification.SendPushNotiByPlayerID(players, data, message)

	res := Response{}
	res.Message = `create chat sucsess`

	return c.JSON(http.StatusOK, res)
}

func bandRegister(c echo.Context) error {
	res := Response{}

	userID, err := strconv.ParseUint(c.FormValue(`user_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	minPrice, err := strconv.ParseInt(c.FormValue(`minPrice`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	maxPrice, err := strconv.ParseInt(c.FormValue(`maxPrice`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	band := model.Band{
		WorkLocation: c.FormValue(`location`),
		MinPrice:     minPrice,
		MaxPrice:     maxPrice,
		About:        c.FormValue(`about`),
		Member:       c.FormValue(`member`),
		UserID:       uint(userID),
		Cover:        c.FormValue(`coverURL`),
	}

	db.Create(&band)

	res.Message = `test`

	genresSTR := strings.Split(c.FormValue(`genres_id`), `,`)
	genres := []int64{}
	for _, gen := range genresSTR {
		genreID, err := strconv.ParseInt(gen, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		genres = append(genres, genreID)
	}

	categoriesSTR := strings.Split(c.FormValue(`categores_id`), `,`)
	categories := []int64{}
	for _, cate := range categoriesSTR {
		categoryID, err := strconv.ParseInt(cate, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		categories = append(categories, categoryID)
	}

	typesSTR := strings.Split(c.FormValue(`types_id`), `,`)
	types := []int64{}
	for _, ty := range typesSTR {
		typeID, err := strconv.ParseInt(ty, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		types = append(types, typeID)
	}

	for _, categoryID := range categories {
		db.Create(&model.BandCategory{
			BandID:     int(band.ID),
			CategoryID: int(categoryID),
		})
	}

	for _, genreID := range genres {
		db.Create(&model.BandGenre{
			BandID:  int(band.ID),
			GenreID: int(genreID),
		})
	}

	acousticImageStr := strings.Split(c.FormValue(`acousticImageStr`), `,`)
	acousticImages := []string{}
	for _, image := range acousticImageStr {
		acousticImages = append(acousticImages, image)
	}

	acousticVideoStr := strings.Split(c.FormValue(`acousticVideoStr`), `,`)
	acousticVideos := []string{}
	for _, video := range acousticVideoStr {
		acousticVideos = append(acousticVideos, video)
	}

	fullbandImageStr := strings.Split(c.FormValue(`fullbandImageStr`), `,`)
	fullbandImages := []string{}
	for _, image := range fullbandImageStr {
		fullbandImages = append(fullbandImages, image)
	}

	fullbandVideoStr := strings.Split(c.FormValue(`fullbandVideoStr`), `,`)
	fullbandVideos := []string{}
	for _, video := range fullbandVideoStr {
		fullbandVideos = append(fullbandVideos, video)
	}

	djImageStr := strings.Split(c.FormValue(`djImageStr`), `,`)
	djImages := []string{}
	for _, image := range djImageStr {
		djImages = append(djImages, image)
	}

	djVideoStr := strings.Split(c.FormValue(`djVideoStr`), `,`)
	djVideos := []string{}
	for _, video := range djVideoStr {
		djVideos = append(djVideos, video)
	}

	stringImageStr := strings.Split(c.FormValue(`stringImageStr`), `,`)
	stringImages := []string{}
	for _, image := range stringImageStr {
		stringImages = append(stringImages, image)
	}

	stringVideoStr := strings.Split(c.FormValue(`stringVideoStr`), `,`)
	stringVideos := []string{}
	for _, video := range stringVideoStr {
		stringVideos = append(stringVideos, video)
	}

	jazzImageStr := strings.Split(c.FormValue(`jazzImageStr`), `,`)
	jazzImages := []string{}
	for _, image := range jazzImageStr {
		jazzImages = append(jazzImages, image)
	}

	jazzVideoStr := strings.Split(c.FormValue(`jazzVideoStr`), `,`)
	jazzVideos := []string{}
	for _, video := range jazzVideoStr {
		jazzVideos = append(jazzVideos, video)
	}

	for _, typeID := range types {
		if typeID == 1 {
			bandType := model.BandType{
				BandID: int(band.ID),
				TypeID: int(typeID),
				Detail: c.FormValue(`acousticAbout`),
			}
			db.Create(&bandType)
			for _, image := range acousticImages {
				db.Create(&model.BandImage{
					Image:      image,
					Thumbnail:  image,
					BandtypeID: bandType.ID,
				})
			}
			for _, videoCode := range acousticVideos {
				db.Create(&model.BandVideo{
					Code:       videoCode,
					Video:      `https://www.youtube.com/watch?v=` + videoCode,
					BandtypeID: bandType.ID,
				})
			}
		} else if typeID == 2 {
			bandType := model.BandType{
				BandID: int(band.ID),
				TypeID: int(typeID),
				Detail: c.FormValue(`fullbandAbout`),
			}
			db.Create(&bandType)
			for _, image := range fullbandImages {
				db.Create(&model.BandImage{
					Image:      image,
					Thumbnail:  image,
					BandtypeID: bandType.ID,
				})
			}
			for _, videoCode := range fullbandVideos {
				db.Create(&model.BandVideo{
					Code:       videoCode,
					Video:      `https://www.youtube.com/watch?v=` + videoCode,
					BandtypeID: bandType.ID,
				})
			}
		} else if typeID == 3 {
			bandType := model.BandType{
				BandID: int(band.ID),
				TypeID: int(typeID),
				Detail: c.FormValue(`djAbout`),
			}
			db.Create(&bandType)
			for _, image := range djImages {
				db.Create(&model.BandImage{
					Image:      image,
					Thumbnail:  image,
					BandtypeID: bandType.ID,
				})
			}
			for _, videoCode := range djVideos {
				db.Create(&model.BandVideo{
					Code:       videoCode,
					Video:      `https://www.youtube.com/watch?v=` + videoCode,
					BandtypeID: bandType.ID,
				})
			}
		} else if typeID == 4 {
			bandType := model.BandType{
				BandID: int(band.ID),
				TypeID: int(typeID),
				Detail: c.FormValue(`stringAbout`),
			}
			db.Create(&bandType)
			for _, image := range stringImages {
				db.Create(&model.BandImage{
					Image:      image,
					Thumbnail:  image,
					BandtypeID: bandType.ID,
				})
			}
			for _, videoCode := range stringVideos {
				db.Create(&model.BandVideo{
					Code:       videoCode,
					Video:      `https://www.youtube.com/watch?v=` + videoCode,
					BandtypeID: bandType.ID,
				})
			}
		} else if typeID == 5 {
			bandType := model.BandType{
				BandID: int(band.ID),
				TypeID: int(typeID),
				Detail: c.FormValue(`jazzAbout`),
			}
			db.Create(&bandType)
			for _, image := range jazzImages {
				db.Create(&model.BandImage{
					Image:      image,
					Thumbnail:  image,
					BandtypeID: bandType.ID,
				})
			}
			for _, videoCode := range jazzVideos {
				db.Create(&model.BandVideo{
					Code:       videoCode,
					Video:      `https://www.youtube.com/watch?v=` + videoCode,
					BandtypeID: bandType.ID,
				})
			}
		}

	}

	user := model.User{}
	db.First(&user, userID)
	user.RoleID = 2
	db.Save(&user)

	db.Model(&user).Related(&user.Role)
	favourites := []model.Band{}
	db.Joins("JOIN favourites ON favourites.band_id = bands.id AND favourites.user_id = ?", user.ID).Find(&favourites)
	for i, band := range favourites {
		userBand := model.User{}
		db.First(&userBand, band.UserID)
		favourites[i].User = &userBand
	}
	user.Favourites = favourites

	return c.JSON(http.StatusOK, user)
}

func quickBook(c echo.Context) error {

	catID, err := strconv.ParseUint(c.FormValue(`category_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	typeID, err := strconv.ParseUint(c.FormValue(`type_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	price, err := strconv.ParseFloat(c.FormValue(`price`), 64)
	if err != nil {
		fmt.Println(err)
	}

	s := strings.Split(c.FormValue(`genres`), `,`)
	genres := []int64{}
	for _, gen := range s {
		genreID, err := strconv.ParseInt(gen, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		genres = append(genres, genreID)
	}

	bands := []model.Band{}
	res := Response{Message: `not fount band`}

	if err := db.Table("bands").Select("*").
		Joins(`JOIN band_genres ON bands.id = band_genres.band_id `).
		Joins(`JOIN band_types ON bands.id = band_types.band_id `).
		Joins(`JOIN band_categories ON bands.id = band_categories.band_id `).
		Where("band_genres.genre_id IN (?) AND bands.min_price  > ?  AND band_types.type_id = ? AND band_categories.category_id = ?",
			genres, price, typeID, catID).
		Group("bands.user_id").
		Scan(&bands).Error; gorm.IsRecordNotFoundError(err) {

		return c.JSON(http.StatusOK, res)
	}
	if len(bands) == 0 {
		return c.JSON(http.StatusOK, res)
	}

	userID, err := strconv.ParseUint(c.FormValue(`user_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	lat, err := strconv.ParseFloat(c.FormValue(`latitude`), 64)
	if err != nil {
		fmt.Println(err)
	}
	lon, err := strconv.ParseFloat(c.FormValue(`longitude`), 64)
	if err != nil {
		fmt.Println(err)
	}

	dateStr := c.FormValue(`date`)
	timeStr := c.FormValue(`time`)
	input := dateStr + " " + timeStr
	layout := "2006-01-02 15:04:05"

	t, _ := time.Parse(layout, input)

	booking := model.Booking{
		UserID:     uint(userID),
		CategoryID: uint(catID),
		TypeID:     uint(typeID),
		Location:   c.FormValue(`location`),
		DateTime:   t,
		Latitude:   lat,
		Longitude:  lon,
		Duration:   c.FormValue(`duration`),
		Price:      price,
		Time:       &timeStr,
		Date:       &dateStr,
		IsQuick:    true,
	}
	db.Create(&booking)

	for _, genID := range genres {
		db.Create(&model.BookingGenre{
			GenreID:   int(genID),
			BookingID: int(booking.ID),
		})
	}

	user := model.User{}
	db.First(&user, userID)

	for _, band := range bands {
		db.Create(&model.BookingBand{
			BandID:    int(band.ID),
			BookingID: int(booking.ID),
		})
		message := `New order from ` + user.Name
		db.Create(&model.Notification{
			BookingID: getID(int(booking.ID)),
			UserID:    int(band.UserID),
			Title:     `New order`,
			Detail:    message,
		})
		players := []model.PlayerID{}
		db.Find(&players, `user_id = ?`, int(band.UserID))

		// db.Model(&notifition).Related(&notifition.User)
		// db.Model(&booking).Related(&booking.User)
		// db.Model(&booking).Related(&booking.Category)
		// db.Model(&booking).Related(&booking.Type)
		// notifition.Booking = &booking
		// payloadJSON, _ := json.Marshal(notifition)
		data := `{
			"page": "form_noti",
			"payload": "` + `yoyo` + `"
		}`
		notification.SendPushNotiByPlayerID(players, data, message)
	}

	return c.JSON(http.StatusOK, booking)
}

func booking(c echo.Context) error {

	userID, err := strconv.ParseUint(c.FormValue(`user_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	lat, err := strconv.ParseFloat(c.FormValue(`latitude`), 64)
	if err != nil {
		fmt.Println(err)
	}

	lon, err := strconv.ParseFloat(c.FormValue(`longitude`), 64)
	if err != nil {
		fmt.Println(err)
	}

	catID, err := strconv.ParseUint(c.FormValue(`category_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	typeID, err := strconv.ParseUint(c.FormValue(`type_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	price, err := strconv.ParseFloat(c.FormValue(`price`), 64)
	if err != nil {
		fmt.Println(err)
	}

	dateStr := c.FormValue(`date`)
	timeStr := c.FormValue(`time`)
	input := dateStr + " " + timeStr
	layout := "2006-01-02 15:04:05"

	t, _ := time.Parse(layout, input)

	booking := model.Booking{
		UserID:     uint(userID),
		CategoryID: uint(catID),
		TypeID:     uint(typeID),
		Location:   c.FormValue(`location`),
		DateTime:   t,
		Latitude:   lat,
		Longitude:  lon,
		Duration:   c.FormValue(`duration`),
		Price:      price,
		Time:       &timeStr,
		Date:       &dateStr,
	}
	db.Create(&booking)

	band := model.Band{}
	db.First(&band, c.FormValue(`band_id`))

	db.Create(&model.BookingBand{
		BandID:    int(band.ID),
		BookingID: int(booking.ID),
	})

	user := model.User{}
	db.First(&user, userID)
	message := `New order from ` + user.Name
	notifition := &model.Notification{
		BookingID: getID(int(booking.ID)),
		UserID:    int(band.UserID),
		Title:     `New order`,
		Detail:    message,
	}
	db.Create(&notifition)

	players := []model.PlayerID{}
	db.Find(&players, `user_id = ?`, int(band.UserID))

	db.Model(&notifition).Related(&notifition.User)
	db.Model(&booking).Related(&booking.User)
	db.Model(&booking).Related(&booking.Category)
	db.Model(&booking).Related(&booking.Type)
	notifition.Booking = &booking
	// payloadJSON, _ := json.Marshal(notifition)
	data := `{
		"page": "form_noti",
		"payload": "` + `yoyo` + `"
	}`

	notification.SendPushNotiByPlayerID(players, data, message)

	return c.JSON(http.StatusOK, booking)
}

func selectBandBooking(c echo.Context) error {
	res := Response{}
	booking := model.Booking{}
	if err := db.First(&booking, c.FormValue(`booking_id`)).Error; gorm.IsRecordNotFoundError(err) {
		res.Message = `Not Found Booking`
		return c.JSON(http.StatusOK, res)
	}
	bandID, err := strconv.ParseUint(c.FormValue(`band_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	db.Model(&booking).Where(`id = ?`, c.FormValue(`booking_id`)).Updates(map[string]interface{}{"band_id": bandID, "status": 2})

	res.Message = `Select Band Success`
	return c.JSON(http.StatusOK, res)
}

func readNotification(c echo.Context) error {
	res := Response{}
	noti := model.Notification{}
	db.Model(&noti).Where(`id = ?`, c.FormValue(`id`)).Update(`seen`, true)
	res.Message = `Seen NOti Success`
	return c.JSON(http.StatusOK, res)
}

func bandAcceptBooking(c echo.Context) error {
	res := Response{}
	band := model.Band{}
	noti := model.Notification{}
	db.First(&band, `user_id = ?`, c.FormValue(`user_id`))

	bookingBand := model.BookingBand{}
	if err := db.First(&bookingBand, `booking_id = ? AND band_id = ?`, c.FormValue(`booking_id`), band.ID).Error; gorm.IsRecordNotFoundError(err) {
		res.Message = `Not Found Booking`
		return c.JSON(http.StatusOK, res)
	}

	db.Model(&bookingBand).Update("status", 1)
	res.Message = `Accept Success`

	booking := model.Booking{}
	db.First(&booking, c.FormValue(`booking_id`))
	if booking.IsQuick == true {
		booking.Status = 1
	} else {
		booking.Status = 2
		booking.BandID = &band.ID
	}
	db.Save(&booking)
	db.Model(&noti).Where("user_id = ? AND booking_id = ?", c.FormValue(`user_id`), c.FormValue(`booking_id`)).Update("status", 2)

	user := model.User{}
	db.First(&user, band.UserID)
	message := user.Name + ` accepted the order.`
	players := []model.PlayerID{}
	db.Find(&players, `user_id = ?`, booking.UserID)
	data := `{
			"page": "form_noti",
			"payload": "` + `yoyo` + `"
		}`

	notification.SendPushNotiByPlayerID(players, data, message)

	return c.JSON(http.StatusOK, res)
}
func bandDiscardtBooking(c echo.Context) error {
	res := Response{}
	band := model.Band{}
	db.First(&band, `user_id = ?`, c.FormValue(`user_id`))
	bookingBand := model.BookingBand{}
	if err := db.First(&bookingBand, `booking_id = ? AND band_id = ?`, c.FormValue(`booking_id`), band.ID).Error; gorm.IsRecordNotFoundError(err) {
		res.Message = `Not Found Booking`
		return c.JSON(http.StatusOK, res)
	}

	db.Delete(&bookingBand)
	res.Message = `Discard Success`
	return c.JSON(http.StatusOK, res)
}

func paymentBandBooking(c echo.Context) error {
	res := Response{}
	booking := model.Booking{}
	if err := db.First(&booking, c.FormValue(`booking_id`)).Error; gorm.IsRecordNotFoundError(err) {
		res.Message = `Not Found Booking`
		return c.JSON(http.StatusOK, res)
	}
	db.Model(&booking).Update("status", 3)
	res.Message = `Payment Success`
	return c.JSON(http.StatusOK, res)
}

func confirmBooking(c echo.Context) error {
	res := Response{}
	booking := model.Booking{}

	userID, err := strconv.ParseUint(c.FormValue(`user_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	bandID, err := strconv.ParseUint(c.FormValue(`band_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	bookingID, err := strconv.ParseUint(c.FormValue(`booking_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	rate, err := strconv.ParseFloat(c.FormValue(`rate`), 64)
	if err != nil {
		fmt.Println(err)
	}

	review := model.Review{
		UserID:    uint(userID),
		BandID:    uint(bandID),
		BookingID: uint(bookingID),
		Rate:      float32(rate),
		Detail:    c.FormValue(`detail`),
	}

	db.Create(&review)

	if err := db.First(&booking, c.FormValue(`booking_id`)).Error; gorm.IsRecordNotFoundError(err) {
		res.Message = `Not Found Booking`
		return c.JSON(http.StatusOK, res)
	}
	db.Model(&booking).Update("status", 4)
	res.Message = `Confirm Success`
	return c.JSON(http.StatusOK, res)
}

func reviewBooking(c echo.Context) error {
	res := Response{}
	booking := model.Booking{}

	userID, err := strconv.ParseUint(c.FormValue(`user_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	bandID, err := strconv.ParseUint(c.FormValue(`band_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	bookingID, err := strconv.ParseUint(c.FormValue(`booking_id`), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	rate, err := strconv.ParseFloat(c.FormValue(`rate`), 64)
	if err != nil {
		fmt.Println(err)
	}

	review := model.Review{
		UserID:    uint(userID),
		BandID:    uint(bandID),
		BookingID: uint(bookingID),
		Rate:      float32(rate),
		Detail:    c.FormValue(`detail`),
	}

	db.Create(&review)

	if err := db.First(&booking, c.FormValue(`booking_id`)).Error; gorm.IsRecordNotFoundError(err) {
		res.Message = `Not Found Booking`
		return c.JSON(http.StatusOK, res)
	}
	db.Model(&booking).Update("status", 4)
	res.Message = `Confirm Success`
	return c.JSON(http.StatusOK, res)
}

func getCurrentBooking(c echo.Context) error {
	bookings := []model.Booking{}
	if err := db.Order("created_at desc").Find(&bookings, `user_id = ?`, c.FormValue(`user_id`)).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, `booking`)
	}
	for i := range bookings {
		db.Model(&bookings[i]).Related(&bookings[i].User)
		db.Model(&bookings[i]).Related(&bookings[i].Category)
		db.Model(&bookings[i]).Related(&bookings[i].Type)
		db.Model(&bookings[i]).Related(&bookings[i].Genres, "genres")
		db.Model(&bookings[i]).Related(&bookings[i].BandSelect, "booking_id")
		for j, bandselect := range bookings[i].BandSelect {
			db.Model(&bandselect).Related(&bandselect.Band)
			user := model.User{}
			db.First(&user, bandselect.Band.UserID)

			band := model.Band{}
			band = bandselect.Band
			db.Model(&band).Related(&band.Reviews)
			rateAvg := model.GetRateAVG(band.Reviews)
			band.RateAvg = &rateAvg

			bookings[i].BandSelect[j].Band = band
			bookings[i].BandSelect[j].Band.User = &user
		}

		if bookings[i].BandID != nil {
			band := model.Band{}
			if err := db.First(&band, `id = ?`, bookings[i].BandID).Error; gorm.IsRecordNotFoundError(err) {

			}
			bookings[i].Band = &band
		}
	}
	return c.JSON(http.StatusOK, bookings)
}

func getCurrentBookingBand(c echo.Context) error {
	band := model.Band{}
	db.First(&band, `user_id = ?`, c.FormValue(`user_id`))
	bookings := []model.Booking{}
	if err := db.Order("created_at desc").Find(&bookings, `band_id = ?`, band.ID).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusOK, `booking`)
	}
	for i := range bookings {
		db.Model(&bookings[i]).Related(&bookings[i].User)
		db.Model(&bookings[i]).Related(&bookings[i].Category)
		db.Model(&bookings[i]).Related(&bookings[i].Type)
		db.Model(&bookings[i]).Related(&bookings[i].Genres, "genres")

		if bookings[i].BandID != nil {
			band := model.Band{}
			if err := db.First(&band, `id = ?`, bookings[i].BandID).Error; gorm.IsRecordNotFoundError(err) {

			}
			bookings[i].Band = &band
		}
	}
	return c.JSON(http.StatusOK, bookings)
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
	userChats := []model.Chat{}
	if err := db.Where(`user_id = ? `, c.Param(`uID`)).Or(`to_id = ? `, c.Param(`uID`)).Order("created_at desc").Find(&chats).Error; gorm.IsRecordNotFoundError(err) {

	}
	for _, chat := range chats {

		if len(userChats) > 0 {
			valCheck := 0
			for _, user := range userChats {

				if user.UserID == chat.UserID && user.ToID == chat.ToID {
					valCheck++
				} else if user.UserID == chat.ToID && user.ToID == chat.UserID {
					valCheck++
				}

			}

			if valCheck == 0 {
				db.Model(&chat).Related(&chat.User)
				db.Model(&chat).Related(&chat.ToUser, "ToID")
				userChats = append(userChats, chat)
			}

		} else {
			db.Model(&chat).Related(&chat.User)
			db.Model(&chat).Related(&chat.ToUser, "ToID")
			userChats = append(userChats, chat)
		}

	}
	return c.JSON(http.StatusOK, userChats)
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
	for i := range bands {
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

//////////// ADMIN////////////////

func adminLogin(c echo.Context) error {
	user := model.User{}
	name := c.FormValue("name")
	if err := db.Where("name = ? AND role_id = ?", name, 3).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return c.JSON(http.StatusNotFound, "invalid user")
	}
	password := c.FormValue("password")
	check := comparePasswords(user.Password, []byte(password))
	if check {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusNotFound, "invalid password")
	}
}

func getUsers(c echo.Context) error {
	users := []model.User{}
	db.Find(&users, `role_id != ?`, 3)
	for _, user := range users {
		db.Model(&user).Related(&user.Role)
	}

	return c.JSON(http.StatusOK, users)
}

func getWorks(c echo.Context) error {
	bookings := []model.Booking{}
	db.Find(&bookings)
	for i := range bookings {
		db.Model(&bookings[i]).Related(&bookings[i].User)
		db.Model(&bookings[i]).Related(&bookings[i].Category)
		db.Model(&bookings[i]).Related(&bookings[i].Type)
		db.Model(&bookings[i]).Related(&bookings[i].Genres, "genres")
		if bookings[i].BandID != nil {
			band := model.Band{}
			if err := db.First(&band, `id = ?`, bookings[i].BandID).Error; gorm.IsRecordNotFoundError(err) {

			}
			user := model.User{}
			db.First(&user, band.UserID)
			band.User = &user
			bookings[i].Band = &band
		}
	}

	return c.JSON(http.StatusOK, bookings)
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

func getID(x int) *int {
	return &x
}
func getUID(x uint) *uint {
	return &x
}

func testNoti(c echo.Context) error {
	players := []model.PlayerID{}
	db.Find(&players, `user_id = ?`, 3)
	data := `{
		"page": "chat",
		"user_id": "41",
		"to_id": "3"
	}`
	notification.SendPushNotiByPlayerID(players, data, `สวัสดีครับ`)

	return c.JSON(http.StatusOK, players)
}

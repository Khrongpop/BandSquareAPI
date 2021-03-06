package model

import (
	"math"
	"time"
)

type Band struct {
	// gorm.Model
	ID           uint       `gorm:"primary_key" json:"id"`
	About        string     `json:"about"`
	Member       string     `json:"member"`
	WorkLocation string     `json:"work_location"`
	MinPrice     int64      `json:"min_price"`
	MaxPrice     int64      `json:"max_price"`
	Cover        string     `json:"cover"`
	UserID       uint       `json:"user_id"`
	CreatedAt    time.Time  `json:"created_at"`
	User         *User      `json:"user"`
	Types        []BandType `json:"types"`
	Genres       []Genre    `gorm:"many2many:band_genres;" json:"genres"`
	Categories   []Category `gorm:"many2many:band_categories" json:"categories"`
	CategoryList *string    `json:"categores_list"`
	GenreList    *string    `json:"genres_list"`
	Bookings     []Booking  `json:"bookings"`
	Reviews      []Review   `json:"reviews"`
	RateAvg      *float32   `json:"rate_avg"`
}

type BandType struct {
	ID     uint        `gorm:"primary_key" json:"id"`
	BandID int         `json:"band_id"`
	TypeID int         `json:"type_id"`
	Detail string      `json:"detail"`
	Band   *Band       `json:"band"`
	Type   Type        `json:"type"`
	Images []BandImage `gorm:"foreignkey:BandtypeID" json:"images"`
	Videos []BandVideo `gorm:"foreignkey:BandtypeID" json:"videos"`
}

type BandGenre struct {
	BandID  int `gorm:"primary_key:true" json:"band_id"`
	GenreID int `gorm:"primary_key:true" json:"genre_id"`
}

type BandCategory struct {
	BandID     int `gorm:"primary_key:true" json:"band_id"`
	CategoryID int `gorm:"primary_key:true" json:"category_id"`
}

type BandImage struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	BandtypeID uint   `json:"bandtype_id"`
	Image      string `json:"image"`
	Thumbnail  string `json:"thumbnail"`
}

type BandVideo struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	BandtypeID uint   `json:"bandtype_id"`
	Video      string `json:"video"`
	Code       string `json:"code"`
}

func GetGenreList(band Band) string {
	genresList := ""
	for i, genre := range band.Genres {

		genresList += genre.Name
		if i != len(band.Genres)-1 {
			genresList += " , "
		}
	}
	return genresList
}

func GetCategoryList(band Band) string {
	categoryList := ""
	for i, category := range band.Categories {

		categoryList += category.Name
		if i != len(band.Categories)-1 {
			categoryList += " , "
		}
	}
	return categoryList
}

func GetRateAVG(reviews []Review) float32 {
	rateAvg := float64(0)
	reviewCount := len(reviews)
	if reviewCount > 0 {
		rateSum := float64(0)
		for _, review := range reviews {
			rateSum += float64(review.Rate)
		}
		rateAvg = rateSum / float64(reviewCount)
	}
	rateAvg = math.Round(rateAvg*100) / 100
	return float32(rateAvg)
}

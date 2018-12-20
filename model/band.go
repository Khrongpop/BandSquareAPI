package model

import (
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
	Genres       []Genre    `gorm:"many2many:band_genres;" json:"genres `
	Categories   []Category `gorm:"many2many:band_categories;" json:"categories `
	CategoryList *string    `json:"categores_list"`
	GenreList    *string    `json:"genres_list"`
	Bookings     []Booking  `json:"bookings"`
	Reviews      []Review   `json:"reviews"`
	RateAvg      *float32   `json:"rate_avg"`
}

type BandType struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	BandID int    `json:"band_id"`
	TypeID int    `json:"type_id"`
	Detail string `json:"detail"`
	Band   *Band  `json:"band"`
	Type   Type   `json:"type"`
}

type BandGenre struct {
	BandID  int `gorm:"primary_key:true" json:"band_id"`
	GenreID int `gorm:"primary_key:true" json:"genre_id"`
}

type BandCategory struct {
	BandID     int `gorm:"primary_key:true" json:"band_id"`
	CategoryID int `gorm:"primary_key:true" json:"category_id"`
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
	rateAvg := float32(0)
	reviewCount := len(reviews)
	if reviewCount > 0 {
		rateSum := float32(0)
		for _, review := range reviews {
			rateSum += review.Rate
		}
		rateAvg = rateSum / float32(reviewCount)
	}
	return rateAvg
}

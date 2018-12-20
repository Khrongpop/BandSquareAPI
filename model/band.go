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
	Genres       []Genre    `gorm:"many2many:band_genres;" json:"genres`
	CategoryList *string    `json:"categores_list"`
	GenreList    *string    `json:"genres_list"`
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

package model

type Category struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
type Genre struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
type Type struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// package main

// import (
// 	// _ "github.com/go-sql-driver/mysql"

// 	"fmt"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// var db *gorm.DB
// var err error

// type Person struct {
// 	ID        uint   `json:”id”`
// 	FirstName string `json:”firstname”`
// 	LastName  string `json:”lastname”`
// }

// func main() {
// 	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/boomcrud")
// 	// db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/BandSquare")
// 	if err != nil {
// 		panic("failed to connect database")
// 		log.Fatal(err)
// 	}

// 	db.Create(&Person{FirstName: "“Jane”", LastName: "“Smith”"})
// 	db.Create(&Person{FirstName: "HO", LastName: "He"})

// 	r := gin.Default()
// 	r.GET("/", GetProjects)
// 	r.Run(":8080")
// }

// func GetProjects(c *gin.Context) {
// 	var people []Person
// 	if err := db.Find(&people).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, people)
// 	}
// }

// func main() {

// 	db, err := sql.Open("mysql",
// 		"root:@tcp(127.0.0.1:3306)/BandSquare")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 		// do something here
// 	}

// var (
// 	id   int
// 	name string
// )
// rows, err := db.Query("select id, name from users where id = ?", 1)
// // rows, err := db.Query("select id, name from users ")
// if err != nil {
// 	log.Fatal(err)
// }
// defer rows.Close()
// for rows.Next() {
// 	err := rows.Scan(&id, &name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(id, name)
// }
// err = rows.Err()
// if err != nil {
// 	log.Fatal(err)
// }
// }

// db.CreateTable(&Place{})
// db.CreateTable(&Town{})
// t := Town{
// 	Name: "TestTown",
// }

// p1 := Place{
// 	Name:   "Test",
// 	TownId: 1,
// }

// p2 := Place{
// 	Name:   "Test2",
// 	TownId: 1,
// }

// err := db.Save(&t).Error
// err = db.Save(&p1).Error
// err = db.Save(&p2).Error
// if err != nil {
// 	panic(err)
// }

// places := []Place{}
// err = db.Find(&places).Error
// for i, _ := range places {
// 	db.Model(places[i]).Related(&places[i].Town)
// }
// if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println(places)
// }
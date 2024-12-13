package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//データベースの項目
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:14"`
	Age  int
}

func main() {
	dsn := "root:123qwecc@tcp(db:3306)/ecc_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var user User
	//条件に合ったデータの取得nameがKato
	db.Where("name=?", "Kato").First(&user)
	//SQLを記述することも可能
	// db.Raw("select * from users where name=?", "Kato").Scan(&user)
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{

		//Katoのageを返す
			"age": user.Age,
		})
	})
	router.Run(":8000")
}
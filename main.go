package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

type Comment struct {
	gorm.Model
	Name    string
	Content string
}

func GetComment(c *gin.Context) {
	var commentData []Comment
	arr := make([]gin.H, 0)
	DB.Find(&commentData)
	for _, data := range commentData {
		arr = append(arr, gin.H{
			"time":    data.Model.CreatedAt,
			"name":    data.Name,
			"content": data.Content,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": arr,
	})
}
func AddComment(c *gin.Context) {
	var form Comment
	err := c.ShouldBindJSON(&form)
	log.Println(&form)
	if err != nil {
		log.Println(err)
	}
	DB.Create(&form)
}

func main() {
	r := gin.Default()
	dsn := "test:123456@tcp(localhost:3306)/xsstest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err := DB.AutoMigrate(
		&Comment{},
	)
	if err != nil {
		log.Println(err)
	}
	log.Println("db ok")
	r.GET("/api/", GetComment)
	r.POST("/api/add", AddComment)
	r.Run(":8080")
}

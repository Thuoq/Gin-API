package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func main() {
	dsn := "root:thuong123@tcp(127.0.0.1:3306)/gin-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(":::: db connection", db)
	//now := time.Now().UTC()
	//
	//item := TodoItem{
	//	Id:          1,
	//	Title:       "Task 1",
	//	Description: "Content 1",
	//	Status:      "Doing",
	//	CreatedAt:   &now,
	//	UpdatedAt:   nil,
	//}
	//jsData, err := json.Marshal(item)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(jsData))
	//
	//jsonString := "{\"id\":1,\"title\":\"Task 1\",\"description\":\"Content 1\",\"status\":\"Doing\",\"created_at\":\"2024-02-13T07:44:00.669384Z\",\"updated_at\":null}"
	//
	//var item2 TodoItem
	//if err := json.Unmarshal([]byte(jsonString), &item2); err != nil {
	//	log.Println(err)
	//}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("")
			items.POST("")
			items.DELETE("/:id")
			items.GET("/:id")
			items.PATCH("/:id")
		}
	}
	r.Run(":3000")
}

//cmd: "go work init" tạo file go.work để làm việc với nhiều project trong cùng 1 thư mục
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	restaurantgin "golang200lab-learn/module/restaurant/transport/gin"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalln("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		restaurant := v1.Group("/restaurant")
		{
			restaurant.POST("", restaurantgin.CreateRestaurantHandel(db))
			restaurant.GET("/:restaurant_id", restaurantgin.GetRestaurantHandel(db))
		}

	}

	router.Run(":3000")
}

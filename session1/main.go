//cmd: "go work init" tạo file go.work để làm việc với nhiều project trong cùng 1 thư mục
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
}

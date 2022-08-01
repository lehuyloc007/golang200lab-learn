//cmd: "go work init" tạo file go.work để làm việc với nhiều project trong cùng 1 thư mục
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", createTodoItem(db))
			items.GET("/", getListTodoItem(db))
			items.GET("/:item_id", getTodoItem(db))
		}

	}

	router.Run(":3000")
}

//business struct
type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string {
	return "todo_items"
}

type ToDoItemCreate struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItemCreate) TableName() string {
	return ToDoItem{}.TableName()
}

type ToDoItemUpdate struct {
	//tại sao phải dùng con trỏ ở *string vì gorm không thể đưa string về chuỗi rỗng
	//mà con trỏ thì có 2 trạng thái: 1 là nil không có dữ liệu => bỏ qua, 2 là có chuỗi rỗng thì vẫn là có dữ liệu
	Title  *string `json:"title" gorm:"column:title;"`
	Status *string `json:"status" gorm:"column:status;"`
}

func (ToDoItemUpdate) TableName() string {
	return ToDoItem{}.TableName()
}

func createTodoItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data ToDoItemCreate
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}
		if err := db.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": 1, "data": data.Id})
	}
}
func getListTodoItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}
		var paging DataPaging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}
		if paging.Page <= 0 {
			paging.Page = 1
		}
		if paging.Limit <= 0 {
			paging.Limit = 10
		}
		offset := (paging.Page - 1) * paging.Limit

		var result []ToDoItem

		if err := db.Table(ToDoItem{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id desc").
			Find(&result).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": 1, "paging": paging, "data": result})
	}
}
func getTodoItem(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data ToDoItem
		id, err := strconv.Atoi(ctx.Param("item_id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := db.Where("id= ?", id).First(&data).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, gin.H{"ok": 1, "data": data})
	}
}

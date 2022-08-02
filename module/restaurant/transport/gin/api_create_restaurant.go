package restaurantgi

import (
	"fmt"
	restaurantbiz "golang200lab-learn/module/restaurant/biz"
	restaurantmodel "golang200lab-learn/module/restaurant/model"
	restaurantstorage "golang200lab-learn/module/restaurant/storage"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurantHandel(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		storage := restaurantstorage.NewSQLStore(db)
		fmt.Println(storage)
		biz := restaurantbiz.NewCreateRestaurantBiz(storage)

		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data.Id})

	}
}

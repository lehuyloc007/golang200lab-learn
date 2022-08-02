package restaurantbiz

import (
	"context"
	restaurantmodel "golang200lab-learn/module/restaurant/model"
)

//b2 tạo 1 interface interface này có nhiệm vụ là
type CreateRestaurantStore interface {
	InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

//b7 để sử dụng struct ở b1 từ bên ngoài tạo hàm này
func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

//b1 tạo 1 struct đại diện cho Encapsulation(bao đóng) của biz này
type createRestaurantBiz struct {
	//b3 gọi cái store để tạo
	store CreateRestaurantStore
}

//b4 tạo method của struct
func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	//b5 kiểm tra xem các lỗi
	if err := data.Validate(); err != nil {
		return err
	}
	//b6 tạo dữ liệu
	if err := biz.store.InsertRestaurant(ctx, data); err != nil {
		return err
	}
	return nil
}

package restaurantbiz

import (
	"context"
	restaurantmodel "golang200lab-learn/module/restaurant/model"
)

type CreateRestaurantStore interface {
	InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.InsertRestaurant(ctx, data); err != nil {
		return err
	}
	return nil
}

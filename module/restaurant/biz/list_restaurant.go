package restaurantbiz

import (
	"context"
	"golang200lab-learn/common"
	restaurantmodel "golang200lab-learn/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListRestaurant(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

func NewListRestaurantBiz(store ListRestaurantStore) *lisRestaurantBiz {
	return &lisRestaurantBiz{store: store}
}

type lisRestaurantBiz struct {
	store ListRestaurantStore
}

func (biz *lisRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListRestaurant(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}

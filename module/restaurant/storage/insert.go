package restaurantstorage

import (
	"context"
	restaurantmodel "golang200lab-learn/module/restaurant/model"
)

//method này implement cái hàm InsertRestaurant
func (store *sqlStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := store.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

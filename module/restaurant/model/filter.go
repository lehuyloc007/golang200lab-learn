package restaurantmodel

//filter dùng để lọc
type Filter struct {
	OwnerId int `json:"owner_id" form:"owner_id"`
}

package restaurantstorage

import (
	"gorm.io/gorm"
)

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

//b1 tạo struct chứa  db
type sqlStore struct {
	db *gorm.DB
}

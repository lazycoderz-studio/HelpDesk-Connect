package repository

import "gorm.io/gorm"

type User struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &User{
		db: db,
	}
}

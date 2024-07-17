package repository

import "gorm.io/gorm"

type Repositories struct {
	User UserRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: NewUserRepo(db),
	}
}

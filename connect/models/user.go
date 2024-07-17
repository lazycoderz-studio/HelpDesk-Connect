package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName     string     `gorm:"type:varchar(100);not null"`
	Email        string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone        string     `gorm:"type:varchar(20);not null"`
	Password     string     `gorm:"type:varchar(255);not null"`
	RefreshToken string     `gorm:"type:varchar(255)"`
	LastLogin    *time.Time `gorm:"type:timestamp"`
	IsActive     bool       `gorm:"default:true"`
	CreatedAt    time.Time  `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;default:current_timestamp"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    Id uint `json:"id" gorm:"primarykey"`
    Username  string `json:"username" gorm:"unique"`
    Email string `json:"email" gorm:"unique"`
    Password string `json:"password"`
    CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP"`
    DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
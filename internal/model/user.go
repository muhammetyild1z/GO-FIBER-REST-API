package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         string    `json:"id" gorm:"uniqueIndex;not null"`
	UserName   string    `json:"username" gorm:" not null"`
	Password   string    `json:"password" gorm:" not null"`
	Email      string    `json:"email" gorm:" not null"`
	CreateDate time.Time `json"createdate" gorm:" not null"`
	UpdateDate time.Time `json:"updatedate" gorm:" not null"`
}

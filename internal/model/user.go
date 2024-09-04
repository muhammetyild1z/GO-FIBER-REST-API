package user

import (
	"time"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Id         string    `json:"id" gorm:"uniqueIndex;not null"`
	UserName   string    `json:"username" gorm:" not null"`
	Password   string    `json:"password" gorm:" not null"`
	Email      string    `json:"email" gorm:" not null"`
	CreateDate time.Time `json"createdate" gorm:" not null"`
	UpdateDate time.Time `json:"updatedate" gorm:" not null"`
}

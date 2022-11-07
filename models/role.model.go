package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	ID   int32  `gorm:"primary_key"`
	Role string `json:"role" gorm:"type:varchar(300) not null"`
}

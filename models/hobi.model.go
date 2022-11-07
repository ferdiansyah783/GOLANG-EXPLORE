package models

import "gorm.io/gorm"

type Hobi struct {
	gorm.Model

	ID     int16  `gorm:"primary_key"`
	Hobi   string `json:"hobi" gorm:"type:varchar(200) not null"`
	UserID int
}

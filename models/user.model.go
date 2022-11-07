package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID     int32  `gorm:"primary_key"`
	Name   string `json:"name" gorm:"type:varchar(300) not null"`
	RoleID int    `gorm:"ForeignKey:id"`
	Role   Role
	Hobies []Hobi `gorm:"ForeignKey:UserID"`
}

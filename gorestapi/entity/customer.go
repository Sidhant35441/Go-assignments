package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(100)"`
	Age    int    `json:"age" gorm:"type:int(2)"`
	Gender string `json:"gender" gorm:"type:varchar(6)"`
	Mobile int    `json:"mobile" gorm:"type:int(10)"`
}

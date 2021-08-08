package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `json:"name" gorm:"type:varchar(40)" binding:"min=2,max=40"`
	Age    int    `json:"age" gorm:"type:int(6)" binding:"gte=18,lte=99"`
	Gender string `json:"gender" gorm:"type:varchar(6)" binding:"min=1,max=6"`
	Email  string `json:"email" gorm:"type:varchar(40)" binding:"required,email"`
}

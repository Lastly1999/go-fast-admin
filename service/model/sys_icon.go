package model

import "gorm.io/gorm"

type SysIcon struct {
	gorm.Model
	IconName string `json:"iconName" gorm:"column:icon_name"`
}

package model

import (
	"fast-admin-service/global"
)

type SysIcon struct {
	global.Model
	IconName string `json:"iconName" gorm:"column:icon_name"`
}

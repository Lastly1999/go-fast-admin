package services

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
)

type SysSystemService struct {
}

type ISysSystemService interface {
	GetSystemIcons() (icons []*model.SysIcon, err error)
}

// GetSystemIcons 获取系统图标列表
func (s SysSystemService) GetSystemIcons() (icons []*model.SysIcon, err error) {
	if err = global.GLOBAL_DB.Find(&icons).Error; err != nil {
		return nil, err
	}
	return icons, nil
}

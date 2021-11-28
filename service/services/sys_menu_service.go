package services

import (
	"errors"
	"fast-admin-service/global"
	"fast-admin-service/model"
)

type BaseMenuService struct {
}

type IBaseMenuService interface {
	PutBaseMenu(baseMenu *model.SysBaseMenu) (err error)
	DeleteBaseMenu(id int) (err error)
	GetBaseMenu() (baseMenus []*model.SysBaseMenu, err error)
	UpdateBaseMenu(baseMenu *model.SysBaseMenu) error
	GetBaseInfo(id string) (sysBaseMenu *model.SysBaseMenu, err error)
}

// PutBaseMenu  添加系统菜单
func (baseMenuService *BaseMenuService) PutBaseMenu(baseMenu *model.SysBaseMenu) (err error) {
	if err = global.GLOBAL_DB.Create(&baseMenu).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBaseMenu 删除系统菜单
func (baseMenuService *BaseMenuService) DeleteBaseMenu(id int) (err error) {
	if err = global.GLOBAL_DB.Delete(&model.SysBaseMenu{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetBaseMenu 获取全部系统菜单
func (baseMenuService *BaseMenuService) GetBaseMenu() (baseMenus []*model.SysBaseMenu, err error) {
	if err = global.GLOBAL_DB.Find(&baseMenus).Error; err != nil {
		return nil, err
	}
	return baseMenus, nil
}

// UpdateBaseMenu 修改系统菜单
func (baseMenuService *BaseMenuService) UpdateBaseMenu(baseMenu *model.SysBaseMenu) error {
	if err := global.GLOBAL_DB.Model(&model.SysBaseMenu{}).Where("id = ?", baseMenu.ID).Updates(&baseMenu).Error; err != nil {
		return errors.New("修改失败")
	}
	return nil
}

// GetBaseInfo 获取菜单详情
func (baseMenuService *BaseMenuService) GetBaseInfo(id string) (sysBaseMenu *model.SysBaseMenu, err error) {
	if err = global.GLOBAL_DB.First(&sysBaseMenu, id).Error; err != nil {
		return nil, errors.New("未找到菜单详情")
	}
	return sysBaseMenu, nil
}

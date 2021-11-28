package services

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
)

type RoleService struct {
}

type IRoleService interface {
	GetRoles() (roles []*model.SysRole, err error)
	DeleteRoleById(id int) (err error)
	PutRole(role *model.SysRole) (err error)
	UpdateRole(role *model.SysRole) (err error)
	UpdateRoleMenu(roleId string, menuIds []uint) error
}

// DeleteRoleById 删除角色
func (roleService *RoleService) DeleteRoleById(id int) (err error) {
	if err = global.GLOBAL_DB.Model(&model.SysRole{}).Where("role_id = ?", id).Delete(&model.SysRole{}).Error; err != nil {
		return err
	}
	return nil
}

// GetRoles 查询角色列表
func (roleService *RoleService) GetRoles() (roles []*model.SysRole, err error) {
	if err = global.GLOBAL_DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// PutRole 添加角色
func (roleService *RoleService) PutRole(role *model.SysRole) (err error) {
	if err = global.GLOBAL_DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

// UpdateRole 更新角色
func (roleService *RoleService) UpdateRole(role *model.SysRole) (err error) {
	if err = global.GLOBAL_DB.Model(&model.SysRole{}).Where("role_id = ?", role.RoleId).Updates(&role).Error; err != nil {
		return err
	}
	return nil
}

// UpdateRoleMenu 更新角色菜单
func (roleService *RoleService) UpdateRoleMenu(roleId uint, menuIds []uint) error {
	var sysBaseMenuIds []*model.SysBaseMenu
	for _, v := range menuIds {
		sysBaseMenuIds = append(sysBaseMenuIds, &model.SysBaseMenu{Model: global.Model{ID: v}})
	}
	err := global.GLOBAL_DB.Model(&model.SysRole{RoleId: roleId}).Association("BaseMenu").Replace(sysBaseMenuIds)
	if err != nil {
		return err
	}
	return nil
}

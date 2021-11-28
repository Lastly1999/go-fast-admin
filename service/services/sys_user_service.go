package services

import (
	"errors"
	"fast-admin-service/global"
	"fast-admin-service/model"
	"fast-admin-service/model/request"
)

type UserService struct {
}

type IUserService interface {
	GetUsers(pageInfo request.PageInfo) (user []*model.SysUser, total int64, err error)
	CreateUser(user *model.SysUser) error
	UpdateUserById(user *model.SysUser) error
	DeleteUserById(userId int) error
	GetSystemUserRoles(user *model.SysUser) (roles []*model.SysRole, err error)
	NewUserAssociationRole(params *request.SystemUserRoleParams) (err error)
}

// GetUsers 获取系统用户
func (userService *UserService) GetUsers(userListParams *request.SystemUserListParams) (users []*model.SysUser, total int64, err error) {
	offset := (userListParams.Page - 1) * userListParams.PageSize
	userModel := global.GLOBAL_DB.Model(&model.SysUser{})
	if userListParams.StartTime.String() != "0001-01-01 00:00:00" && userListParams.EndTime.String() != "0001-01-01 00:00:00" {
		userModel.Where("created_at >= ? And created_at <= ?", userListParams.StartTime, userListParams.EndTime)
	}
	err = userModel.Count(&total).Error
	if err != nil {
		return users, total, err
	}
	err = userModel.Where("user_name LIKE ?", "%"+userListParams.KeyWords+"%").Preload("Role").Find(&users).Limit(userListParams.PageSize).Offset(offset).Error
	if err != nil {
		return users, total, err
	}
	return users, total, nil
}

// CreateUser 新增系统用户
func (userService *UserService) CreateUser(user *model.SysUser) error {
	result := global.GLOBAL_DB.Create(&user)
	if result.RowsAffected == 0 || result.Error != nil {
		return errors.New("创建失败")
	}
	// 关联权限
	err := result.Association("Role").Append(&user.Role)
	if err != nil {
		return errors.New("关联权限失败")
	}
	return nil
}

// UpdateUserById 修改系统用户信息
func (userService *UserService) UpdateUserById(user *model.SysUser) error {
	result := global.GLOBAL_DB.Where("id = ? ", user.ID).Updates(&user)
	if result.RowsAffected == 0 || result.Error != nil {
		return errors.New("更新失败")
	}
	// 重新关联权限
	err := global.GLOBAL_DB.Model(&model.SysUser{Model: global.Model{ID: user.ID}}).Association("Role").Replace(&user.Role)
	if err != nil {
		return errors.New("关联权限失败")
	}
	return nil
}

// DeleteUserById 删除系统用户
func (userService *UserService) DeleteUserById(userId int) error {
	if err := global.GLOBAL_DB.Where("id", userId).Delete(&model.SysUser{}).Error; err != nil {
		return err
	}
	return nil
}

// GetSystemUserRoles 获取系统用户角色列表
func (userService *UserService) GetSystemUserRoles(user *model.SysUser) (roles []*model.SysRole, err error) {
	err = global.GLOBAL_DB.Model(&user).Association("Role").Find(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// NewUserAssociationRole 新增用户关联角色
func (userService *UserService) NewUserAssociationRole(params *request.SystemUserRoleParams) (err error) {
	var role []*model.SysRole
	for _, v := range params.RoleIds {
		role = append(role, &model.SysRole{
			RoleId: v,
		})
	}
	user := &model.SysUser{
		Model: global.Model{
			ID: params.UserId,
		},
	}
	err = global.GLOBAL_DB.Model(&user).Association("Role").Replace(&role)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserRole 更新用户默认角色
func (userService *UserService) UpdateUserRole(params *request.SysRoleDefaultParams) (err error) {
	if err = global.GLOBAL_DB.Model(&model.SysUser{}).Where("id = ?", params.UserId).Update("role_id", params.RoleId).Error; err != nil {
		return err
	}
	return nil
}

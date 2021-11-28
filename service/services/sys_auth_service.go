package services

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
	"fast-admin-service/pkg/captcha"
)

type AuthService struct {
}

type IAuthService interface {
	CheckAuth(user *model.SysUser) (userRecord *model.SysUser, err error)
	GenerateVerificode() (string, string, error)
	GetSystemPermissionsMenu(roleId string) (menus []model.SysBaseMenu, err error)
	GetSystemPermissionsMenuIds(roleId string) (sysBaseMenuIds []int, err error)
	GetSystemUserInfoById(userId int) (user *model.SysUser, err error)
}

// CheckAuth 用户授权验证 检验用户是否存在
func (authService *AuthService) CheckAuth(user *model.SysUser) (userRecord *model.SysUser, err error) {
	err = global.GLOBAL_DB.Model(&model.SysUser{}).Where("user_name = ? and pass_word = ? ", user.UserName, user.PassWord).Preload("Role").First(&userRecord).Error
	if err != nil {
		return userRecord, err
	}
	return userRecord, nil
}

// GenerateVerificode 生成图形验证码
func (authService *AuthService) GenerateVerificode() (string, string, error) {
	// 生成图形验证码
	id, base := captcha.GetCaptcha()
	//// 存储redis
	//err := redis.Rdb.Set(id, id, 60*time.Second).Err()
	////  异常
	//if err != nil {
	//	return "", "", err
	//}
	return id, base, nil
}

// GetSystemPermissionsMenu 获取系统权限菜单
func (authService *AuthService) GetSystemPermissionsMenu(roleId string) ([]model.SysBaseMenu, error) {
	var sysRole model.SysRole
	if err := global.GLOBAL_DB.Where("role_id = ?", roleId).Preload("BaseMenu").Find(&sysRole).Error; err != nil {
		return nil, err
	}
	return sysRole.BaseMenu, nil
}

// GetSystemPermissionsMenuIds 获取用户权限id组
func (authService *AuthService) GetSystemPermissionsMenuIds(roleId string) (sysBaseMenuIds []int, err error) {
	var sysRole model.SysRole
	if err = global.GLOBAL_DB.Preload("BaseMenu").Where("role_id = ?", roleId).Find(&sysRole).Error; err != nil {
		return nil, nil
	}
	for _, v := range sysRole.BaseMenu {
		sysBaseMenuIds = append(sysBaseMenuIds, int(v.ID))
	}
	return sysBaseMenuIds, nil
}

// GetSystemUserInfoById 获取系统用户信息
func (authService *AuthService) GetSystemUserInfoById(userId int) (user *model.SysUser, err error) {
	if err = global.GLOBAL_DB.Model(&model.SysUser{}).Preload("Role").Where("id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

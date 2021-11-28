package model

import "fast-admin-service/global"

type SysUser struct {
	global.Model
	UserName   string    `json:"userName" gorm:"column:user_name;comment:用户名"`
	PassWord   string    `json:"passWord" gorm:"column:pass_word;comment:账户密码"`
	UserAvatar string    `json:"userAvatar" gorm:"column:user_avatar;comment:用户头像"`
	NikeName   string    `json:"nikeName" gorm:"column:nike_name;comment:用户昵称"`
	RoleId     uint      `json:"roleId" gorm:"column:role_id;comment:权限id;"`
	Status     *bool     `json:"status" gorm:"column:status;default:true;comment:用户状态"`
	Role       []SysRole `json:"role" gorm:"many2many:sys_user_roles;"` // 多对多 用户角色表
}

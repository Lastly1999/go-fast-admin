package model

import (
	"fast-admin-service/global"
)

type SysBaseMenu struct {
	global.Model
	Name     string    `json:"label" gorm:"column:name;comment:权限菜单名称"`
	Path     string    `json:"path" gorm:"column:path;comment:权限菜单路径"`
	ParentId uint      `json:"pId" gorm:"column:parentId;comment:根菜单id"`
	Icon     string    `json:"icon" gorm:"column:icon;comment:图标名称"`
	Role     []SysRole `gorm:"many2many:sys_authority_menus"` // 多对多 用户角色表
}

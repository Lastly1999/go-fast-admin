package model

import (
	"fast-admin-service/utils"
	"gorm.io/gorm"
)

type SysRole struct {
	RoleId    uint            `json:"roleId" gorm:"column:role_id;primary_key;comment:权限id"`
	RoleName  string          `json:"roleName" gorm:"column:role_name;comment:角色名称"`
	Describe  string          `json:"describe" gorm:"column:describe;comment:角色别名"`
	Status    *bool           `json:"status" gorm:"column:status;default:1;comment:状态"`
	User      []SysUser       `json:"user" gorm:"many2many:sys_user_roles"`                     // 多对多 用户表
	Users     []SysUser       `json:"users" gorm:"ForeignKey:RoleId;AssociationForeignKey:ID" ` // 一对多 用户表
	BaseMenu  []SysBaseMenu   `json:"baseMenu" gorm:"many2many:sys_authority_menus"`            // 多对多 菜单表
	CreatedAt utils.LocalTime `json:"createdAt"`
	UpdatedAt utils.LocalTime `json:"updatedAt"`
	DeletedAt gorm.DeletedAt  `json:"-" gorm:"index"`
}

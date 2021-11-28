package model

type SysUserRole struct {
	RoleId string `json:"roleId" gorm:"column:sys_role_role_id;comment:角色id"`
	UserId uint   `json:"userId" gorm:"column:sys_user_id;comment:用户id"`
}

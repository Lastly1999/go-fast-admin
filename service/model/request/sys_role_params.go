package request

type SysRoleParams struct {
	RoleName string `json:"roleName"`
	Describe string `json:"describe"`
}

type SysPutRoleParams struct {
	RoleId   uint   `json:"roleId" binding:"required"`
	RoleName string `json:"roleName"`
	Describe string `json:"describe"`
}

type SysRoleMenuParams struct {
	RoleId       uint   `json:"roleId"`
	PermissionId []uint `json:"permissionId"`
}

type SysRoleDefaultParams struct {
	RoleId string `json:"roleId"`
	UserId int    `json:"userId"`
}

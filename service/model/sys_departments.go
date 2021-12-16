package model

import "fast-admin-service/global"

type SysDepartments struct {
	global.Model
	DepName string    `json:"depName" gorm:"column:dep_name;comment:部门名称"`
	DepPid  uint      `json:"depPid" gorm:"column:dep_pid;comment:根部门id"`
	Users   []SysUser `gorm:"ForeignKey:DepId;AssociationForeignKey:ID" json:"users"`
}

package request

import (
	"fast-admin-service/utils"
)

type Login struct {
	UserName string `json:"userName" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
	CodeAuth string `json:"codeAuth" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

type SystemUserListParams struct {
	PageInfo
	StartTime utils.LocalTime `json:"startTime"`
	EndTime   utils.LocalTime `json:"endTime"`
}

type SystemUserParams struct {
	Id         uint   `json:"id"`
	UserName   string `json:"userName"`
	PassWord   string `json:"passWord"`
	UserAvatar string `json:"userAvatar"`
	NikeName   string `json:"nikeName"`
	RoleId     uint   `json:"roleId"`
	RoleIds    []uint `json:"roleIds"`
}

type SystemUserRoleParams struct {
	UserId  uint   `json:"id" binding:"required"`
	RoleIds []uint `json:"roleIds" binding:"required"`
}

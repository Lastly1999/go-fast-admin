package services

import "fast-admin-service/global"

type SysCasbinRuleService struct {
}

type ISysCasbinRuleService interface {
	GetPolicys() [][]string
	AddPolicy()
	DelPolicy()
}

func (sysCasbinRuleService *SysCasbinRuleService) GetPolicys() [][]string {
	return global.GLOBAL_Enforcer.GetPolicy()
}

func (sysCasbinRuleService *SysCasbinRuleService) AddPolicy() {
	panic("implement me")
}

func (sysCasbinRuleService *SysCasbinRuleService) DelPolicy() {
	panic("implement me")
}

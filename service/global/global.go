package global

import (
	"fast-admin-service/utils"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB       *gorm.DB
	ZAP_LOG         *zap.Logger
	GLOBAL_Enforcer *casbin.Enforcer
)

type Model struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt utils.LocalTime
	UpdatedAt utils.LocalTime
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

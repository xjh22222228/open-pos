package models

import (
	"time"
)

// ErpTenant 租户表
type ErpTenant struct {
	CommonModel
	DeletedAt

	TenantCode string `gorm:"type:varchar(30);not null;uniqueIndex:uk_tenant_code;comment:租户编码"`
	TenantName string `gorm:"type:varchar(100);not null;index:idx_tenant_name;comment:租户名称"`

	TenantType uint8 `gorm:"type:tinyint;default:1;comment:1单店 2总部 3门店"`
	ParentID   *uint `gorm:"type:int;default:null;comment:父租户"`

	ContactPerson string `gorm:"type:varchar(50);default:'';comment:联系人"`
	ContactPhone  string `gorm:"type:varchar(20);default:'';comment:联系电话"`
	Address       string `gorm:"type:varchar(255);default:'';comment:地址"`

	Status     uint8      `gorm:"type:tinyint;not null;default:1;comment:1正常 0禁用"`
	ExpireTime *time.Time `gorm:"type:datetime;default:null;comment:租户有效期"`
	PlanID     *uint      `gorm:"type:int;default:null;comment:订阅套餐"`
}

func (ErpTenant) TableName() string { return "erp_tenants" }

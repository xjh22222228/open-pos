package models

import "time"

// ErpUser 用户表
type ErpUser struct {
	CommonModel
	DeletedAt

	StoreId uint64 `gorm:"type:bigint;comment:门店ID"`
	Username string `gorm:"type:varchar(50);not null;comment:登录账号"`
	Password string `gorm:"type:varchar(100);not null;comment:登录密码"`
	RealName string `gorm:"type:varchar(50);comment:真实姓名"`
	Phone         string         `gorm:"size:20;default:'';comment:手机号"`
	Email         string         `gorm:"size:100;default:'';comment:邮箱"`
	RoleID        uint           `gorm:"default:0;comment:关联角色ID"`
	Status        int8           `gorm:"not null;default:1;comment:用户状态：1=启用，0=禁用"`
	LastLoginTime *time.Time     `gorm:"comment:最后登录时间"`
}

func (ErpUser) TableName() string { return "erp_users" }

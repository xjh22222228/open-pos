package models

// ErpUser 用户表
type ErpUser struct {
	BaseCommonModel
	DeletedAt

	Username string `gorm:"type:varchar(50);not null;uniqueIndex:uk_tenant_username;comment:登录账号"`
	Password string `gorm:"type:varchar(100);not null;comment:登录密码"`
	RealName string `gorm:"type:varchar(50);not null;comment:真实姓名"`
	Status   uint8  `gorm:"type:tinyint;not null;default:1;comment:状态"`
}

func (ErpUser) TableName() string { return "erp_users" }

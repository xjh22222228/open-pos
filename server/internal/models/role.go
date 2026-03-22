package models

// ErpRole 角色表
type ErpRole struct {
	CommonModel
	DeletedAt

	RoleId      uint64 `gorm:"type:bigint;not null;index:idx_role_id;comment:角色唯一标识"`
	RoleName    string `gorm:"type:varchar(50);not null;comment:角色名称"`
	Description string `gorm:"type:varchar(255);default:'';comment:角色描述"`
	Status      uint8  `gorm:"type:tinyint;not null;default:1;comment:状态：1=启用，0=禁用"`
}

func (ErpRole) TableName() string { return "erp_roles" }

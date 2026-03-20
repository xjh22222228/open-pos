package models

// ErpStore 门店表
type ErpStore struct {
	BaseCommonModel
	DeletedAt

	StoreName string `gorm:"type:varchar(100);not null;comment:门店名称"`
	Address   string `gorm:"type:varchar(255);default:'';comment:门店地址"`
	Status    uint8  `gorm:"type:tinyint;not null;default:1;comment:1正常 0禁用"`
}

func (ErpStore) TableName() string { return "erp_stores" }

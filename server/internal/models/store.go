package models

import "time"

// ErpStore 门店表
type ErpStore struct {
	BaseCommonModel
	DeletedAt

	StoreName string `gorm:"type:varchar(100);comment:门店名称"`
	Address   string `gorm:"type:varchar(255);default:'';comment:门店地址"`
	Status    uint8  `gorm:"type:tinyint;not null;default:1;comment:1正常 0禁用"`

	OpenTime  *time.Time `gorm:"type:time;default:null;comment:营业开始时间"`
	CloseTime *time.Time `gorm:"type:time;default:null;comment:营业结束时间"`

	ContactPerson string `gorm:"size:50;default:'';comment:门店联系人"`
	ContactPhone  string `gorm:"size:20;default:'';comment:门店联系电话"`
}

func (ErpStore) TableName() string { return "erp_stores" }

package models

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint      `gorm:"primarykey"`
	TenantId  uint64    `gorm:"type:bigint;not null;comment:租户唯一标识（雪花算法ID）"`
	CreatedAt time.Time `gorm:"autoCreateTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoCreateTime;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
}

// BaseCommonModel 除了租户表，用户表，其他所有表都应该继承
type BaseCommonModel struct {
	CommonModel
	StoreId uint64 `gorm:"type:bigint;not null;comment:门店ID"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除时间"`
}

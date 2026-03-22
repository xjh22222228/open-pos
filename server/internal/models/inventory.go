package models

// ErpInventory 库存表（门店级独立）
type ErpInventory struct {
	BaseCommonModel
	DeletedAt

	GoodsId         uint64 `gorm:"type:bigint;not null;index:idx_goods_id;comment:商品唯一标识"`
	Quantity        int64  `gorm:"type:bigint;not null;default:0;comment:当前库存数量"`
	LockedQuantity  int64  `gorm:"type:bigint;not null;default:0;comment:锁定库存数量"`
	WarningQuantity int64  `gorm:"type:bigint;not null;default:0;comment:预警库存数量"`
}

func (ErpInventory) TableName() string { return "erp_inventories" }

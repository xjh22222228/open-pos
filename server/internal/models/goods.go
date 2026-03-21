package models

// ErpGoods 商品表
type ErpGoods struct {
	BaseCommonModel
	DeletedAt

	GoodsId    uint64 `gorm:"type:bigint;not null;uniqueIndex:uk_goods_id;comment:商品唯一标识(雪花ID)"`
	CategoryId uint64 `gorm:"type:bigint;not null;index:idx_category_id;comment:商品分类ID"`

	GoodsName string `gorm:"type:varchar(100);not null;index:idx_goods_name;comment:商品名称"`
	Barcode   string `gorm:"type:varchar(50);default:'';index:idx_barcode;comment:商品条码"`

	SalePrice     float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:零售价"`
	PurchasePrice float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:进价"`
	StockQuantity int64   `gorm:"type:bigint;not null;default:0;comment:当前库存(整数)"`

	Status uint8  `gorm:"type:tinyint;not null;default:1;comment:状态: 1上架 0下架"`
	Remark string `gorm:"type:varchar(255);default:'';comment:备注"`
}

func (ErpGoods) TableName() string { return "erp_goods" }

package models

// ErpCategory 商品类目表，使用 parent_id 自关联支持无限级分类
type ErpCategory struct {
	BaseCommonModel
	DeletedAt

	CategoryId   uint64 `gorm:"type:bigint;not null;uniqueIndex:uk_category_id;comment:类目ID"`
	CategoryName string `gorm:"type:varchar(100);not null;comment:类目名称"`
	ParentId     uint64 `gorm:"type:bigint;not null;default:0;index:idx_parent_id;comment:父类目ID，0表示根节点"`
	Sort         int32  `gorm:"type:int;not null;default:0;comment:排序值，越小越靠前"`
	Status       uint8  `gorm:"type:tinyint;not null;default:1;comment:1正常 0禁用"`
	Remark       string `gorm:"type:varchar(255);default:'';comment:备注"`
}

func (ErpCategory) TableName() string { return "erp_categories" }

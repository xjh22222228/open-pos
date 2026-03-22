package goods

import (
	"errors"

	"github.com/xjh22222228/open-erp/server/database"
	"github.com/xjh22222228/open-erp/server/internal/cryptoutils"
	"github.com/xjh22222228/open-erp/server/internal/models"
	"gorm.io/gorm"
)

type GoodsService struct {
	db *gorm.DB
}

func NewGoodsService() *GoodsService {
	return &GoodsService{db: database.SqlDB}
}

type CreateGoodsInput struct {
	TenantId      uint64
	CategoryId    uint64
	GoodsName     string
	Barcode       string
	SalePrice     float64
	PurchasePrice float64
	Status        uint8
	Remark        string
}

type UpdateGoodsInput struct {
	TenantId      uint64
	GoodsId       uint64
	CategoryId    uint64
	GoodsName     string
	Barcode       string
	SalePrice     float64
	PurchasePrice float64
	Status        uint8
	Remark        string
}

func (s *GoodsService) Create(in CreateGoodsInput) (*models.ErpGoods, error) {
	goods := models.ErpGoods{
		CommonModel: models.CommonModel{
			TenantId: in.TenantId,
		},
		GoodsId:       uint64(cryptoutils.RandomSonyflake()),
		CategoryId:    in.CategoryId,
		GoodsName:     in.GoodsName,
		Barcode:       in.Barcode,
		SalePrice:     in.SalePrice,
		PurchasePrice: in.PurchasePrice,
		Status:        in.Status,
		Remark:        in.Remark,
	}
	if err := s.db.Create(&goods).Error; err != nil {
		return nil, err
	}
	return &goods, nil
}

func (s *GoodsService) Update(in UpdateGoodsInput) error {
	res := s.db.Model(&models.ErpGoods{}).
		Where("tenant_id = ? AND goods_id = ?", in.TenantId, in.GoodsId).
		Updates(map[string]any{
			"category_id":    in.CategoryId,
			"goods_name":     in.GoodsName,
			"barcode":        in.Barcode,
			"sale_price":     in.SalePrice,
			"purchase_price": in.PurchasePrice,
			"status":         in.Status,
			"remark":         in.Remark,
		})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("商品不存在")
	}
	return nil
}

func (s *GoodsService) Delete(tenantId, goodsId uint64) error {
	res := s.db.Where("tenant_id = ? AND goods_id = ?", tenantId, goodsId).
		Delete(&models.ErpGoods{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("商品不存在")
	}
	return nil
}

type ListParams struct {
	TenantId   uint64
	GoodsName  string
	CategoryId uint64
	Page       int
	PageSize   int
}

type ListResult struct {
	Goods    []models.ErpGoods `json:"goods"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"pageSize"`
}

func (s *GoodsService) List(p ListParams) (*ListResult, error) {
	var goods []models.ErpGoods
	var total int64

	db := s.db.Model(&models.ErpGoods{}).Where("tenant_id = ?", p.TenantId)

	if p.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+p.GoodsName+"%")
	}
	if p.CategoryId > 0 {
		db = db.Where("category_id = ?", p.CategoryId)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (p.Page - 1) * p.PageSize
	if err := db.Offset(offset).Limit(p.PageSize).Order("id desc").Find(&goods).Error; err != nil {
		return nil, err
	}

	return &ListResult{
		Goods:    goods,
		Total:    total,
		Page:     p.Page,
		PageSize: p.PageSize,
	}, nil
}

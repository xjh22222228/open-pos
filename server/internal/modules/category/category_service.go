package category

import (
	"errors"

	"github.com/xjh22222228/open-erp/server/database"
	"github.com/xjh22222228/open-erp/server/internal/cryptoutils"
	"github.com/xjh22222228/open-erp/server/internal/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{db: database.SqlDB}
}

type CreateCategoryInput struct {
	TenantId     uint64
	StoreId      uint64
	CategoryName string
	ParentId     uint64
	Sort         int32
	Status       uint8
	Remark       string
}

type UpdateCategoryInput struct {
	TenantId     uint64
	StoreId      uint64
	CategoryId   uint64
	CategoryName string
	ParentId     uint64
	Sort         int32
	Status       uint8
	Remark       string
}

type CategoryTreeNode struct {
	ID           uint64             `json:"id"`
	CategoryId   uint64             `json:"categoryId"`
	CategoryName string             `json:"categoryName"`
	ParentId     uint64             `json:"parentId"`
	Sort         int32              `json:"sort"`
	Status       uint8              `json:"status"`
	Remark       string             `json:"remark"`
	Children     []CategoryTreeNode `json:"children"`
}

func (s *CategoryService) Create(input CreateCategoryInput) (*models.ErpCategory, error) {
	if input.ParentId > 0 {
		ok, err := s.existsCategory(input.TenantId, input.StoreId, input.ParentId)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("父类目不存在")
		}
	}

	item := models.ErpCategory{
		BaseCommonModel: models.BaseCommonModel{
			CommonModel: models.CommonModel{
				TenantId: input.TenantId,
			},
			StoreId: input.StoreId,
		},
		CategoryId:   uint64(cryptoutils.RandomSonyflake()),
		CategoryName: input.CategoryName,
		ParentId:     input.ParentId,
		Sort:         input.Sort,
		Status:       input.Status,
		Remark:       input.Remark,
	}

	if err := s.db.Create(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *CategoryService) Update(input UpdateCategoryInput) error {
	var current models.ErpCategory
	if err := s.db.Where("tenant_id = ? AND store_id = ? AND category_id = ?", input.TenantId, input.StoreId, input.CategoryId).First(&current).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("类目不存在")
		}
		return err
	}

	if input.ParentId == input.CategoryId {
		return errors.New("父类目不能是自身")
	}
	if input.ParentId > 0 {
		ok, err := s.existsCategory(input.TenantId, input.StoreId, input.ParentId)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("父类目不存在")
		}
		isDescendant, err := s.isDescendant(input.TenantId, input.StoreId, input.ParentId, input.CategoryId)
		if err != nil {
			return err
		}
		if isDescendant {
			return errors.New("不能把类目移动到自己的子类目下")
		}
	}

	update := map[string]any{
		"category_name": input.CategoryName,
		"parent_id":     input.ParentId,
		"sort":          input.Sort,
		"status":        input.Status,
		"remark":        input.Remark,
	}

	return s.db.Model(&models.ErpCategory{}).
		Where("tenant_id = ? AND store_id = ? AND category_id = ?", input.TenantId, input.StoreId, input.CategoryId).
		Updates(update).Error
}

func (s *CategoryService) Delete(tenantId, storeId, categoryId uint64) error {
	if categoryId == 0 {
		return errors.New("类目ID不能为空")
	}

	var count int64
	if err := s.db.Model(&models.ErpCategory{}).
		Where("tenant_id = ? AND store_id = ? AND parent_id = ?", tenantId, storeId, categoryId).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("请先删除子类目")
	}

	// 2. 校验该分类下是否存在商品
	var goodsCount int64
	if err := s.db.Model(&models.ErpGoods{}).
		Where("tenant_id = ? AND store_id = ? AND category_id = ?", tenantId, storeId, categoryId).
		Count(&goodsCount).Error; err != nil {
		return err
	}
	if goodsCount > 0 {
		return errors.New("该分类下存在商品，不允许删除")
	}

	res := s.db.Where("tenant_id = ? AND store_id = ? AND category_id = ?", tenantId, storeId, categoryId).Delete(&models.ErpCategory{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("类目不存在")
	}
	return nil
}

func (s *CategoryService) Tree(tenantId, storeId uint64) ([]CategoryTreeNode, error) {
	var rows []models.ErpCategory
	if err := s.db.Where("tenant_id = ? AND store_id = ?", tenantId, storeId).
		Order("sort asc, id asc").
		Find(&rows).Error; err != nil {
		return nil, err
	}

	nodesByParent := make(map[uint64][]CategoryTreeNode)
	for _, row := range rows {
		node := CategoryTreeNode{
			ID:           uint64(row.ID),
			CategoryId:   row.CategoryId,
			CategoryName: row.CategoryName,
			ParentId:     row.ParentId,
			Sort:         row.Sort,
			Status:       row.Status,
			Remark:       row.Remark,
			Children:     []CategoryTreeNode{},
		}
		nodesByParent[row.ParentId] = append(nodesByParent[row.ParentId], node)
	}

	var build func(parentId uint64) []CategoryTreeNode
	build = func(parentId uint64) []CategoryTreeNode {
		children := nodesByParent[parentId]
		for i := range children {
			children[i].Children = build(children[i].CategoryId)
		}
		return children
	}

	return build(0), nil
}

func (s *CategoryService) existsCategory(tenantId, storeId, categoryId uint64) (bool, error) {
	var count int64
	if err := s.db.Model(&models.ErpCategory{}).
		Where("tenant_id = ? AND store_id = ? AND category_id = ?", tenantId, storeId, categoryId).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *CategoryService) isDescendant(tenantId, storeId, maybeChildId, ancestorId uint64) (bool, error) {
	if maybeChildId == 0 {
		return false, nil
	}
	cursor := maybeChildId
	for cursor != 0 {
		if cursor == ancestorId {
			return true, nil
		}
		var c models.ErpCategory
		err := s.db.Select("parent_id").
			Where("tenant_id = ? AND store_id = ? AND category_id = ?", tenantId, storeId, cursor).
			First(&c).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false, nil
			}
			return false, err
		}
		cursor = c.ParentId
	}
	return false, nil
}

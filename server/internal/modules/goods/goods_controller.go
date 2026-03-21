package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/xjh22222228/open-erp/server/internal/httputils"
	"github.com/xjh22222228/open-erp/server/internal/middleware"
)

type SaveGoodsRequest struct {
	GoodsId       uint64  `json:"goodsId"`
	CategoryId    uint64  `json:"categoryId" binding:"required"`
	GoodsName     string  `json:"goodsName" binding:"required"`
	Barcode       string  `json:"barcode"`
	SalePrice     float64 `json:"salePrice"`
	PurchasePrice float64 `json:"purchasePrice"`
	StockQuantity int64   `json:"stockQuantity"`
	Status        uint8   `json:"status"`
	Remark        string  `json:"remark"`
}

type ListGoodsRequest struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	GoodsName  string `json:"goodsName"`
	CategoryId uint64 `json:"categoryId"`
}

type DeleteGoodsRequest struct {
	GoodsId uint64 `json:"goodsId" binding:"required"`
}

func CreateController(c *gin.Context) {
	resp := httputils.NewResponse(c)
	currentUser := middleware.GetCurrentUser(c)

	var req SaveGoodsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	service := NewGoodsService()
	goods, err := service.Create(CreateGoodsInput{
		TenantId:      currentUser.TenantId,
		StoreId:       currentUser.StoreId,
		CategoryId:    req.CategoryId,
		GoodsName:     req.GoodsName,
		Barcode:       req.Barcode,
		SalePrice:     req.SalePrice,
		PurchasePrice: req.PurchasePrice,
		StockQuantity: req.StockQuantity,
		Status:        req.Status,
		Remark:        req.Remark,
	})
	if err != nil {
		resp.Error(500, err.Error())
		return
	}

	resp.Success(goods, "创建成功")
}

func UpdateController(c *gin.Context) {
	resp := httputils.NewResponse(c)
	currentUser := middleware.GetCurrentUser(c)

	var req SaveGoodsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	if req.GoodsId == 0 {
		resp.BadRequest("商品ID不能为空")
		return
	}

	service := NewGoodsService()
	err := service.Update(UpdateGoodsInput{
		TenantId:      currentUser.TenantId,
		StoreId:       currentUser.StoreId,
		GoodsId:       req.GoodsId,
		CategoryId:    req.CategoryId,
		GoodsName:     req.GoodsName,
		Barcode:       req.Barcode,
		SalePrice:     req.SalePrice,
		PurchasePrice: req.PurchasePrice,
		StockQuantity: req.StockQuantity,
		Status:        req.Status,
		Remark:        req.Remark,
	})
	if err != nil {
		resp.Error(500, err.Error())
		return
	}

	resp.Success(nil, "更新成功")
}

func DeleteController(c *gin.Context) {
	resp := httputils.NewResponse(c)
	currentUser := middleware.GetCurrentUser(c)

	var req DeleteGoodsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	service := NewGoodsService()
	if err := service.Delete(currentUser.TenantId, currentUser.StoreId, req.GoodsId); err != nil {
		resp.Error(500, err.Error())
		return
	}

	resp.Success(nil, "删除成功")
}

func ListController(c *gin.Context) {
	resp := httputils.NewResponse(c)
	currentUser := middleware.GetCurrentUser(c)

	var req ListGoodsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	service := NewGoodsService()
	result, err := service.List(ListParams{
		TenantId:   currentUser.TenantId,
		StoreId:    currentUser.StoreId,
		GoodsName:  req.GoodsName,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		resp.Error(500, err.Error())
		return
	}

	resp.Success(result)
}

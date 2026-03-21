package category

import (
	"github.com/gin-gonic/gin"
	"github.com/xjh22222228/open-erp/server/internal/httputils"
	"github.com/xjh22222228/open-erp/server/internal/middleware"
)

type SaveCategoryRequest struct {
	CategoryId   uint64 `json:"categoryId"`
	CategoryName string `json:"categoryName" binding:"required"`
	ParentId     uint64 `json:"parentId"`
	Sort         int32  `json:"sort"`
	Status       uint8  `json:"status"`
	Remark       string `json:"remark"`
}

type DeleteCategoryRequest struct {
	CategoryId uint64 `json:"categoryId" binding:"required"`
}

func CreateController(c *gin.Context) {
	resp := httputils.NewResponse(c)
	currentUser := middleware.GetCurrentUser(c)

	var req SaveCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}
	if req.Status != 0 && req.Status != 1 {
		req.Status = 1
	}

	service := NewCategoryService()
	item, err := service.Create(CreateCategoryInput{
		TenantId:     currentUser.TenantId,
		StoreId:      currentUser.StoreId,
		CategoryName: req.CategoryName,
		ParentId:     req.ParentId,
		Sort:         req.Sort,
		Status:       req.Status,
		Remark:       req.Remark,
	})
	if err != nil {
		resp.Error(500, err.Error())
		return
	}

	resp.Success(item, "创建成功")
}

func UpdateController(c *gin.Context) {
	resp := httputils.NewResponse(c)
	currentUser := middleware.GetCurrentUser(c)

	var req SaveCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	if req.CategoryId == 0 {
		resp.BadRequest("类目ID不能为空")
		return
	}

	if req.Status != 0 && req.Status != 1 {
		req.Status = 1
	}

	service := NewCategoryService()
	err := service.Update(UpdateCategoryInput{
		TenantId:     currentUser.TenantId,
		StoreId:      currentUser.StoreId,
		CategoryId:   req.CategoryId,
		CategoryName: req.CategoryName,
		ParentId:     req.ParentId,
		Sort:         req.Sort,
		Status:       req.Status,
		Remark:       req.Remark,
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

	var req DeleteCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(err.Error())
		return
	}

	service := NewCategoryService()
	if err := service.Delete(currentUser.TenantId, currentUser.StoreId, req.CategoryId); err != nil {
		resp.Error(500, err.Error())
		return
	}

	resp.Success(nil, "删除成功")
}

func TreeController(c *gin.Context) {
	resp := httputils.NewResponse(c)
	currentUser := middleware.GetCurrentUser(c)

	service := NewCategoryService()
	tree, err := service.Tree(currentUser.TenantId, currentUser.StoreId)
	if err != nil {
		resp.Error(500, err.Error())
		return
	}
	resp.Success(tree)
}

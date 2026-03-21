package tenant

import (
	"errors"
	"fmt"

	"github.com/xjh22222228/open-erp/server/database"
	"github.com/xjh22222228/open-erp/server/internal/cryptoutils"
	"github.com/xjh22222228/open-erp/server/internal/models"
	"gorm.io/gorm"
)

type TenantService struct {
	db *gorm.DB
}

func NewTenantService() *TenantService {
	return &TenantService{
		db: database.SqlDB,
	}
}

// Register 注册新租户：包含创建租户、默认门店、管理员用户
func (s *TenantService) Register(tenantCode, username, password string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 1. 校验租户编号
		var count int64
		tx.Model(&models.ErpTenant{}).Where("tenant_code = ?", tenantCode).Count(&count)
		if count > 0 {
			return errors.New("租户编号已存在")
		}

		// 2. 创建租户
		tenantId := uint64(cryptoutils.RandomSonyflake())
		newTenant := models.ErpTenant{
			CommonModel: models.CommonModel{
				TenantId: tenantId,
			},
			TenantCode: tenantCode,
			TenantName: fmt.Sprintf("租户_%s", tenantCode), // 默认名称
			Status:     1,                                  // 正常
		}
		if err := tx.Create(&newTenant).Error; err != nil {
			return err
		}
// 3. 创建默认门店
storeId := uint64(cryptoutils.RandomSonyflake())
newStore := models.ErpStore{
	BaseCommonModel: models.BaseCommonModel{
		CommonModel: models.CommonModel{
			TenantId: tenantId,
		},
		StoreId: storeId,
	},
	StoreName: "默认门店",
	Status:    1,
}
if err := tx.Create(&newStore).Error; err != nil {
	return err
}

		// 4. 创建管理员用户
		hashedPassword, _ := cryptoutils.HashPassword(password)
		adminUser := models.ErpUser{
			CommonModel: models.CommonModel{
				TenantId: tenantId,
			},
			StoreId:  storeId,
			Username: username,
			Password: hashedPassword,
			RealName: "管理员",
			Status:   1,
		}
		if err := tx.Create(&adminUser).Error; err != nil {
			return err
		}

		return nil
	})
}

package login

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/xjh22222228/open-erp/server/database"
	"github.com/xjh22222228/open-erp/server/internal/cryptoutils"
	"github.com/xjh22222228/open-erp/server/internal/models"
	"gorm.io/gorm"
)

// CreateSession 创建会话：存入 Redis，有效期 30 天
func (s *LoginService) CreateSession(token string, u *models.ErpUser) error {
	ctx := context.Background()
	userData, _ := json.Marshal(u)

	// 30 天 = 30 * 24h
	expiration := time.Hour * 24 * 30

	key := fmt.Sprintf("session:%s", token)
	err := database.RedisClient.Set(ctx, key, userData, expiration).Err()
	return err
}
type LoginService struct {
	db *gorm.DB
}

func NewLoginService() *LoginService {
	return &LoginService{
		db: database.SqlDB,
	}
}

// Authenticate 核心登录验证逻辑
func (s *LoginService) Authenticate(tenantCode, username, password string) (*models.ErpUser, error) {
	// 1. 校验租户是否存在
	var t models.ErpTenant
	if err := s.db.Where("tenant_code = ? AND status = 1", tenantCode).First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("租户不存在或已禁用")
		}
		return nil, err
	}

	// 2. 校验用户是否存在
	var u models.ErpUser
	if err := s.db.Where("tenant_id = ? AND username = ?", t.TenantId, username).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	// 3. 校验密码 
	if !cryptoutils.CheckPassword(u.Password, password) {
		return nil, errors.New("用户名或密码错误.")
	}

	if u.Status != 1 {
		return nil, errors.New("该账号已被禁用")
	}

	return &u, nil
}

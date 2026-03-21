package database

import (
	"fmt"
	"log"

	"github.com/xjh22222228/open-erp/server/config"
	"github.com/xjh22222228/open-erp/server/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func MySqlStart() {
	var err error
	c := config.GlobalConfig.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)

	SqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("mysql 连接数据库失败: %v", err)
	}
	fmt.Printf("数据库连接成功: %s:%d/%s\n", c.Host, c.Port, c.Database)

	// 自动建表（从统一 models 导入）
	err = SqlDB.AutoMigrate(
		&models.ErpTenant{},
		&models.ErpStore{},
		&models.ErpUser{},
		&models.ErpCategory{},
		&models.ErpGoods{},
	)
	if err != nil {
		log.Panicf("数据库迁移失败: %v", err)
	}
}

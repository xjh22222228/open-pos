package server

import (
	"github.com/xjh22222228/open-erp/server/config"
	"github.com/xjh22222228/open-erp/server/database"
)

func Run() {
	config.LoadConfig()
	database.MySqlStart()
	database.RedisStart()
	RouterRun()
}

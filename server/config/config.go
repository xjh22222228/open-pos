package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var GlobalConfig *Config

func LoadConfig() {
	// 默认使用 dev 环境
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	configName := fmt.Sprintf("config_%s", env)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 支持环境变量
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Failed to read config file (%s.yaml): %v", configName, err)
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Panicf("Failed to unmarshal config: %v", err)
	}

	fmt.Printf("Config [%s] loaded successfully, server port: %d\n", env, GlobalConfig.Server.Port)
}

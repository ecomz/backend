package config

import (
	"fmt"
	"time"

	"github.com/ecomz/backend/libs/utils"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Name string
	Port string
}

func getAppConfig() AppConfig {
	return AppConfig{
		Name: utils.GetStringOrPanic("APP_NAME"),
		Port: utils.GetStringOrPanic("HTTP_PORT"),
	}
}

type DatabaseConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifeTime time.Duration
	MaxConnIdleTime time.Duration
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DSN:             utils.GetStringOrPanic("DSN"),
		MaxOpenConns:    utils.GetIntOrDefault("DB_MAX_OPEN_CONNS", 20),
		MaxIdleConns:    utils.GetIntOrDefault("DB_MAX_IDLE_CONNS", 100),
		MaxConnLifeTime: time.Duration(utils.GetIntOrDefault("MAX_CONN_LIFE_TIME", 60)) * time.Minute,
		MaxConnIdleTime: time.Duration(utils.GetIntOrDefault("MAX_CONN_IDLE_TIME", 100)) * time.Minute,
	}
}

type JWTConfig struct {
	SecretKey  string
	LoginExp   time.Duration
	RefreshExp time.Duration
}

func getJWTConfig() JWTConfig {
	return JWTConfig{
		SecretKey:  utils.GetStringOrPanic("JWT_SECRET_KEY"),
		LoginExp:   time.Duration(utils.GetIntOrDefault("JWT_LOGIN_EXP", 24)) * time.Hour,
		RefreshExp: time.Duration(utils.GetIntOrDefault("JWT_REFRESH_EXP", 7)) * 24 * time.Hour,
	}
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	MaxIdle  int
}

func getRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     utils.GetStringOrPanic("REDIS_HOST"),
		Port:     utils.GetIntOrPanic("REDIS_PORT"),
		Password: utils.GetStringOrPanic("REDIS_PASSWORD"),
		MaxIdle:  utils.GetIntOrDefault("REDIS_MAX_IDLE", 10),
	}
}

func LoadConfigFromFile(path, fileName, ext string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(ext)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("failed reading config file: %v", err)
	}

	return &Config{
		App:      getAppConfig(),
		Database: getDatabaseConfig(),
		JWT:      getJWTConfig(),
		Redis:    getRedisConfig(),
	}
}

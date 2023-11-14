package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Key  string `mapstructure:"APP_KEY"`
	Port string `mapstructure:"APP_PORT"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_DATABASE"`
}

type RedisConfig struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     string `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	Db       int    `mapstructure:"REDIS_DB"`
}

type MailConfig struct {
	Host     string `mapstructure:"MAIL_HOST"`
	Port     int    `mapstructure:"MAIL_PORT"`
	Username string `mapstructure:"MAIL_USERNAME"`
	Password string `mapstructure:"MAIL_PASSWORD"`
	MailFrom string `mapstructure:"MAIL_FROM"`
	ReplyTo  string `mapstructure:"MAIL_REPLY_TO"`
}

type JwtConfig struct {
	Secret string `mapstructure:"JWT_SECRET"`
}

type TokenConfig struct {
	AccessTokenExpiresIn int `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
	RefeshTokenExpiresIn int `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

var App *AppConfig
var Db *DatabaseConfig
var Redis *RedisConfig
var Mail *MailConfig
var Jwt *JwtConfig
var Token *TokenConfig

func LoadConfig() error {
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return unmarshal(&App, &Db, &Redis, &Mail, &Jwt, &Token)
}

func unmarshal(configs ...interface{}) error {
	for _, config := range configs {
		err := viper.Unmarshal(config)
		if err != nil {
			return err
		}
	}

	return nil
}

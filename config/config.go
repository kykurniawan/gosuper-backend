package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Key  string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	MailFrom string
	ReplyTo  string
}

type JwtConfig struct {
	Secret string
}

type TokenConfig struct {
	AccessTokenExpiresIn int
	RefeshTokenExpiresIn int
}

type RabbitMQConfig struct {
	Dsn string
}

type QueueConfig struct {
	Exchange string
	Mail     struct {
		QueueName    string
		RoutingKey   string
		ConsumerName string
	}
}

var App *AppConfig
var Db *DatabaseConfig
var Redis *RedisConfig
var Mail *MailConfig
var Jwt *JwtConfig
var Token *TokenConfig
var RabbitMQ *RabbitMQConfig
var Queue *QueueConfig

func LoadConfig() error {
	viper.AddConfigPath("../")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	App = &AppConfig{}
	App.Port = viper.GetString("APP_PORT")
	App.Key = viper.GetString("APP_KEY")

	Db = &DatabaseConfig{}
	Db.Host = viper.GetString("DB_HOST")
	Db.Port = viper.GetString("DB_PORT")
	Db.Username = viper.GetString("DB_USERNAME")
	Db.Password = viper.GetString("DB_PASSWORD")
	Db.Database = viper.GetString("DB_DATABASE")

	Redis = &RedisConfig{}
	Redis.Password = viper.GetString("REDIS_PASSWORD")
	Redis.Host = viper.GetString("REDIS_HOST")
	Redis.Port = viper.GetString("REDIS_PORT")
	Redis.Db = viper.GetInt("REDIS_DB")

	Mail = &MailConfig{}
	Mail.Host = viper.GetString("MAIL_HOST")
	Mail.Port = viper.GetInt("MAIL_PORT")
	Mail.Username = viper.GetString("MAIL_USERNAME")
	Mail.Password = viper.GetString("MAIL_PASSWORD")
	Mail.MailFrom = viper.GetString("MAIL_FROM")
	Mail.ReplyTo = viper.GetString("MAIL_REPLY_TO")

	Jwt = &JwtConfig{}
	Jwt.Secret = App.Key

	Token = &TokenConfig{}
	Token.AccessTokenExpiresIn = viper.GetInt("ACCESS_TOKEN_EXPIRES_IN")
	Token.RefeshTokenExpiresIn = viper.GetInt("REFRESH_TOKEN_EXPIRES_IN")

	RabbitMQ = &RabbitMQConfig{}
	RabbitMQ.Dsn = viper.GetString("RABBITMQ_DSN")

	Queue = &QueueConfig{}
	Queue.Exchange = viper.GetString("QUEUE_EXCHANGE")
	Queue.Mail.QueueName = "mail-queue"
	Queue.Mail.RoutingKey = "mail-routing-key"
	Queue.Mail.ConsumerName = "mail-consumer"

	return nil
}

package database

import (
	"fmt"
	"gosuper/app/models"
	"gosuper/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Db.Username,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Otp{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

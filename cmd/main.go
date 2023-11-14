package main

import (
	"gosuper/app"
	"gosuper/app/integrations/database"
	"gosuper/config"
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	err = config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err = database.CreateConnection()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := app.InitializeApp(db)
	app.Run()
}

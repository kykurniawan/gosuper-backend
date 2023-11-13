package main

import (
	"gosuper/app"
	"gosuper/app/integrations/database"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	godotenv.Load()

	db, err = database.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := app.InitializeApp(db)
	app.Run()
}

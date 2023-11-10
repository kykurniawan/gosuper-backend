package main

import (
	"gosuper/app"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	db, err := app.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}

	app := app.InitializeApp(db)

	app.Run()
}

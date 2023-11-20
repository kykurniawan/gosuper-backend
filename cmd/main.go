package main

import (
	"gosuper/app"
	"gosuper/app/integrations/database"
	"gosuper/app/integrations/rabbitmq"
	"gosuper/config"
	"log"

	"github.com/rabbitmq/amqp091-go"

	"gorm.io/gorm"
)

var db *gorm.DB
var amqp *amqp091.Connection

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

	amqp, err = rabbitmq.CreateConnection()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := app.InitializeApp(db, amqp)
	app.Run()
}

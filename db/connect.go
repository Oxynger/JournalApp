package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Client инстанс для подключеной базы данных
	client mongo.Client
)

const ()

// Client godoc
func Client() mongo.Client {
	return client
}

// Connect Получает инстанс подключения к базе данных
func Connect(uri string) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	auth := options.Credential{
		Username:    "root",
		Password:    "t0p5ecret",
		PasswordSet: true,
	}

	connectOptions := options.Client().ApplyURI(uri).SetAuth(auth)

	connect, err := mongo.Connect(timeout, connectOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := connect.Ping(timeout, nil); err != nil {
		log.Fatal(err)
	}

	client = *connect

}

package db

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client инстанс для подключеной базы данных
var client mongo.Client

// Client godoc
func Client() mongo.Client {
	return client
}

// Connect к БД
func Connect(uri string) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	auth := options.Credential{
		Username:    viper.GetString("mongodb_root_username"),
		Password:    viper.GetString("mongodb_root_password"),
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

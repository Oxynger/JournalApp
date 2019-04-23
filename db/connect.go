package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Oxynger/JournalApp/config"
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

func authorizeMongoUser(username, password string) options.Credential {
	return options.Credential{
		Username:    username,
		Password:    password,
		PasswordSet: true,
	}
}

func isEmptyMongoUser(conf config.Config) bool {
	return conf.MongoUsername == ""
}

func authOptions(conf config.Config) (connectOptions *options.ClientOptions) {

	if isEmptyMongoUser(conf) {
		connectOptions = options.Client().ApplyURI(conf.MongoURI)
	} else {
		auth := authorizeMongoUser(conf.MongoUsername, conf.MongoPassword)
		connectOptions = options.Client().ApplyURI(conf.MongoURI).SetAuth(auth)
	}

	return connectOptions

}

// Connect Получает инстанс подключения к базе данных
func Connect(conf config.Config) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	connectOptions := authOptions(conf)

	connect, err := mongo.Connect(timeout, connectOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := connect.Ping(timeout, nil); err != nil {
		log.Fatal(err)
	}

	client = *connect

}

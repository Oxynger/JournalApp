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

func authorizeMongoUser(username, password string) options.Credential {
	return options.Credential{
		Username:    username,
		Password:    password,
		PasswordSet: true,
	}
}

func isEmptyMongoUser(mognoUser options.Credential) bool {
	return mognoUser.Username == ""
}

func authOptions() (connectOptions *options.ClientOptions) {
	auth := authorizeMongoUser(viper.Get("MONGODB_ROOT_USERNAME").(string), viper.Get("MONGODB_ROOT_PASSWORD").(string))

	if isEmptyMongoUser(auth) {
		connectOptions = options.Client().ApplyURI(viper.Get("mongodb_uri").(string))
	} else {
		connectOptions = options.Client().ApplyURI(viper.Get("mongodb_uri").(string)).SetAuth(auth)
	}

	return connectOptions

}

// Connect Получает инстанс подключения к базе данных
func Connect() {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	connectOptions := authOptions()

	connect, err := mongo.Connect(timeout, connectOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := connect.Ping(timeout, nil); err != nil {
		log.Fatal(err)
	}

	client = *connect
}

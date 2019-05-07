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

func authOptions(mongoUri string, auth options.Credential) (connectOptions *options.ClientOptions) {
	if isEmptyMongoUser(auth) {
		connectOptions = options.Client().ApplyURI(mongoUri)
	} else {
		connectOptions = options.Client().ApplyURI(mongoUri).SetAuth(auth)
	}

	return connectOptions

}

// Connect Получает инстанс подключения к базе данных
func Connect(mongoUri string) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	auth := authorizeMongoUser(
		viper.GetString("mongodb_root_username"),
		viper.GetString("mongodb_root_password"),
	)

	connectOptions := authOptions(mongoUri, auth)

	connect, err := mongo.Connect(timeout, connectOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := connect.Ping(timeout, nil); err != nil {
		log.Fatal(err)
	}

	client = *connect
}

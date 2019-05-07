package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Oxynger/JournalApp/db"
	"github.com/Oxynger/JournalApp/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
	sessions   *SessionService
}

func NewUserService() *UserService {
	u := UserService{
		collection: userCollection(),
		sessions:   NewSessionService(),
	}
	_, err := u.collection.Indexes().CreateOne(context.Background(), user.UserIndexModel())
	if err != nil {
		log.Fatal(err)
	}
	return &u
}

func userCollection() *mongo.Collection {
	client := db.Client()
	return client.Database("test").Collection("Users")
}

func (srv *UserService) Create(u user.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = srv.collection.InsertOne(timeout, u)
	return err
}

func (srv *UserService) findByUsername(username string) (*user.User, bool) {
	var result *user.User
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{Key: "username", Value: username}}
	withoutFields := bson.D{
		{Key: "_id", Value: 0},
		{Key: "deleted_at", Value: 0},
		{Key: "Password", Value: 0},
	}
	findOneOptions := options.FindOne().SetProjection(withoutFields)

	if err := srv.collection.FindOne(timeout, filter, findOneOptions).Decode(&result); err != nil {
		return nil, false
	}
	return result, true
}

func (srv *UserService) Authenticate(creds user.Credentials) (*user.User, error) {
	usr, ok := srv.findByUsername(creds.Username)
	if !ok || comparePasswords(usr.Password, creds.Password) {
		return nil, errors.New("Wrong username or password")
	}
	return usr, nil
}

func comparePasswords(expectedPassword, givenPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(givenPassword)); err != nil {
		return false
	}
	return true
}

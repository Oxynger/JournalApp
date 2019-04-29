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
}

func NewUserService() *UserService {
	u := UserService{userCollection()}
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
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := srv.collection.InsertOne(timeout, u)
	return err
}

func (srv *UserService) FindByUsername(username string) (result *user.User, err error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{{Key: "username", Value: username}}
	withoutFields := bson.D{
		{Key: "_id", Value: 0},
		{Key: "deleted_at", Value: 0},
		{Key: "Password", Value: 0},
	}
	findOneOptions := options.FindOne().SetProjection(withoutFields)

	if err = srv.collection.FindOne(timeout, filter, findOneOptions).Decode(&result); err != nil {
		return nil, err
	}
	return
}

func (srv *UserService) Authenticate(cred user.Credentials) (*string, error) {
	incorrectCreds := errors.New("incorrect login or password")
	user, err := srv.FindByUsername(cred.Username)
	if err != nil {
		return nil, incorrectCreds
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		return nil, incorrectCreds
	}
}

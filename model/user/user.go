package user

import (
	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Role string

type User struct {
	db.Model `bson:"inline"`
	ID       *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Username string              `bson:"username" json:"username" binding:"required"`
	Password string              `bson:"password" json:"password" binding:"required"`
	Role     Role
}

func UserIndexModel() mongo.IndexModel {
	return mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
}

package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"../db"
)

// ScaleType godoc
type ScaleType struct {
	Name      string `bson:"name" json:"name" example:"Имя весов"`
	MaxWeight int    `bson:"max_W" json:"max_W" example:"30"`
	MinWeight int    `bson:"min_W" json:"min_W" example:"100"`
	Deviation int    `bson:"deviation" json:"deviation" example:"5"`
	Passport  string `bson:"passport,omitempty" json:"passport"`
}

// Scales godoc
type Scales struct {
	InventoryNumber string    `bson:"_id,omitempty" json:"inventory_number" example:"5c8a9fa2371e1c3d98756ffa" format:"string"`
	Type            ScaleType `bson:"scaletype" json:"scaletype"`
	SerialNumber    int64     `bson:"serial_number,minsize" json:"serial_number" example:"468844"`

	// VerificationDate c часовым поясом сервера
	VerificationDate time.Time `bson:"verification_date" json:"verification_date" example:"2019-03-14T23:08:14.586Z"`

	// NextVerificationDate c часовым поясом сервера
	NextVerificationDate time.Time `bson:"next_verification_date" json:"next_verification_date" example:"2019-06-15T23:08:14.586Z"`

	Bailee string `bson:"bailee" json:"bailee" example:"Толкунова А.А."`
}

func collection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("scale")

	return coll
}

// ScalesOne Возвращает весы с заданым id
func ScalesOne(id string) (Scales, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	var scales Scales
	err := collection().FindOne(context.TODO(), filter).Decode(&scales)

	if err != nil {
		return Scales{}, err
	}

	return scales, nil
}

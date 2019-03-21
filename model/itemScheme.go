package model

import (
	"context"
	"log"

	"../db"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemField struct {
	Name  string `bson:"name" json:"name" example:"serial_number"`
	Title string `bson:"title" json:"title" example:"Серийный номер"`
	Type  string `bson:"type" json:"type" example:"String"`
}

type ItemScheme struct {
	Name    string      `bson:"name" json:"name" example:"scale"`
	Title   string      `bson:"title" json:"title" example:"Весы"`
	Fields  []ItemField `bson:"fields" json:"fields"`
	Deleted bool        `bson:"deleted" json:"-"`
}

func ItemSchemeCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("itemScheme")

	return coll
}

func SomeAdd() ItemScheme {
	scaleScheme := ItemScheme{
		Name:  "scale",
		Title: "Весы",
		Fields: []ItemField{
			{
				Name:  "name",
				Title: "Название",
				Type:  "String",
			},
			{
				Name:  "serial_number",
				Title: "Серийный номер",
				Type:  "String",
			},
			{
				Name:  "min_w",
				Title: "Минимальный вес",
				Type:  "String",
			},
		},
	}

	insertResault, err := ItemSchemeCollection().InsertOne(context.Background(), scaleScheme)

	if err != nil {
		log.Println(err)
		return ItemScheme{}
	}

	log.Println("Inserted documents: ", insertResault.InsertedID)

	return scaleScheme
}

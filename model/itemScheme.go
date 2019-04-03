package model

import (
	"context"
	"errors"
	"log"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ItemInfo godoc
type ItemInfo struct {
	Name   string   `bson:"name" json:"name" example:"scale"`
	Fields []string `bson:"fields" json:"fields"`
}

//Errors godoc
var (
	ErrNameInvalid = errors.New("name is empty")
)

// ItemField godoc
type ItemField struct {
	Name  string `bson:"name" json:"name" example:"serial_number"`
	Title string `bson:"title" json:"title" example:"Серийный номер"`
	Type  string `bson:"type" json:"type" example:"String"`
}

// ItemScheme godoc
type ItemScheme struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id" example:"5ca10d9d015c736a72b7b3ba"`
	Name    string             `bson:"name" json:"name" example:"scale"`
	Title   string             `bson:"title" json:"title" example:"Весы"`
	Fields  []ItemField        `bson:"fields" json:"fields"`
	Deleted bool               `bson:"deleted" json:"-"`
}

// NewItemScheme godoc
type NewItemScheme struct {
	Name    string      `bson:"name" json:"name" example:"scale"`
	Title   string      `bson:"title" json:"title" example:"Весы"`
	Fields  []ItemField `bson:"fields" json:"fields"`
	Deleted bool        `bson:"deleted" json:"-"`
}

// UpdateItemScheme godoc
type UpdateItemScheme struct {
	Name    string      `bson:"name" json:"name" example:"scale"`
	Title   string      `bson:"title" json:"title" example:"Весы"`
	Fields  []ItemField `bson:"fields" json:"fields"`
	Deleted bool        `bson:"deleted" json:"-"`
}

// Insert godoc
func (s NewItemScheme) Insert() error {
	insertResault, err := ItemSchemeCollection().InsertOne(context.Background(), s)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Inserted documents: ", insertResault.InsertedID)
	return err
}

// Update godoc
func (s UpdateItemScheme) Update(id string) error {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := ItemSchemeCollection().UpdateOne(context.Background(), bson.D{{"_id", ojectID}}, bson.D{{"$set", s}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("updated documents: ", updateResault.UpsertedID)
	return err
}

// Validation godoc
func (s NewItemScheme) Validation() error {
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// Validation godoc
func (s UpdateItemScheme) Validation() error {
	switch {
	case len(s.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// ItemSchemeCollection godoc
func ItemSchemeCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("itemScheme")

	return coll
}

// SomeAdd godoc
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

//ItemSchemeAll get list item schemes godoc
func ItemSchemeAll() ([]ItemScheme, error) {
	cur, err := ItemSchemeCollection().Find(context.Background(), bson.D{{"deleted", false}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())
	listSchemes := []ItemScheme{}
	row := new(ItemScheme)
	for cur.Next(context.Background()) {
		err := cur.Decode(&row)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listSchemes = append(listSchemes, *row)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return listSchemes, err
}

//ItemSchemeOne get list item schemes with id godoc
func ItemSchemeOne(id string) (ItemScheme, error) {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return ItemScheme{}, err
	}
	row := new(ItemScheme)
	err = ItemSchemeCollection().FindOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", ojectID}}, bson.D{{"deleted", false}}}}}).Decode(&row)
	if err != nil {
		log.Println(err)
		return ItemScheme{}, err
	}

	return *row, err
}

// DeleteSchemeOne godoc
func DeleteSchemeOne(id string) error {
	ojectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	updateResault, err := ItemSchemeCollection().UpdateOne(context.Background(), bson.D{{"$and", bson.A{bson.D{{"_id", ojectID}}, bson.D{{"deleted", false}}}}}, bson.D{{"$set", bson.D{{"deleted", true}}}})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("deleted documents: ", updateResault.UpsertedID)
	return err
}

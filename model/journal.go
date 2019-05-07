package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Journal godoc
type Journal struct {
	db.Model `bson:",inline"`
	Daily    bool `bson:"daily" json:"daily" binding:"required"`
	Fixed    bool `bson:"fixed" json:"fixed" binding:"required"`
	Values   map[string]interface{}
}

// journalCollection godoc
func journalCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("Journal")

	return coll
}

// JournalsAll godoc
func JournalsAll() (list []Journal, err error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{
		{
			Key:   "deleted",
			Value: false,
		},
	}

	withoutFields := bson.D{
		{Key: "deleted", Value: 0},
	}

	findOptions := options.Find()
	findOptions.SetProjection(withoutFields)

	cur, err := journalCollection().Find(timeout, filter, findOptions)

	defer cur.Close(timeout)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for cur.Next(timeout) {
		var resault Journal
		err := cur.Decode(&resault)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		list = append(list, resault)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return list, nil
}

func journalFindOne(id primitive.ObjectID) (journal *Journal, err error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{
		{
			Key:   "deleted",
			Value: false,
		},
		{
			Key:   "_id",
			Value: id,
		},
	}

	withoutFields := bson.D{
		{Key: "deleted", Value: 0},
	}

	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(withoutFields)

	err = journalCollection().FindOne(timeout, filter, findOneOptions).Decode(&journal)

	if err != nil {
		return nil, err
	}

	return journal, nil
}

// JournalOne godoc
func JournalOne(id string) (journal *Journal, err error) {
	journalID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	journal, err = journalFindOne(journalID)

	if err != nil {
		return nil, err
	}

	return journal, nil
}

// JournalDelete godoc
func JournalDelete(id string) (journal *Journal, err error) {
	journalID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{
		{
			Key:   "deleted",
			Value: false,
		},
		{
			Key:   "_id",
			Value: journalID,
		},
	}

	deleteSet := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "deleted",
					Value: true,
				},
			},
		},
	}

	journal, err = JournalOne(id)

	if err != nil {
		return nil, err
	}

	_, err = journalCollection().UpdateOne(timeout, filter, deleteSet)

	if err != nil {
		return nil, err
	}

	return journal, nil
}

// AddJournal godoc
func AddJournal(journal Journal) (*Journal, error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	insertedResault, err := journalCollection().InsertOne(timeout, journal)

	if err != nil {
		return nil, err
	}

	resaultJournal, err := journalFindOne(insertedResault.InsertedID.(primitive.ObjectID))

	if err != nil {
		return nil, err
	}

	return resaultJournal, nil
}

// JournalUpdate godoc
func JournalUpdate(id string, journal Journal) (*Journal, error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	journalID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: journalID,
		},
	}

	update := bson.D{
		{
			Key:   "$set",
			Value: journal,
		},
	}

	journalCollection().UpdateOne(timeout, filter, update)

	resaultJournal, err := journalFindOne(journalID)

	if err != nil {
		return nil, err
	}

	return resaultJournal, nil
}

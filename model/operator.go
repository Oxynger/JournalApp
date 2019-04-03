package model

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Operator соотвествует сущности controller.
type Operator struct {
	db.Model
	FirstName  string `bson:"first_name" json:"first_name" example:"Олег"`
	MiddleName string `bson:"middle_name" json:"middle_name" example:"Олеговичь"`
	LastName   string `bson:"last_name" json:"last_name" example:"Олегов"`
	Password   []byte `bson:"password" json:"password" example:"qwert"`
}

// HashPassword encrypts operator password
func (o *Operator) HashPassword() error {
	var pwd interface{}

	pwd = o.Password
	convertPwd, ok := pwd.([]byte)

	if ok == false {
		return errors.New("Not correct password")
	}

	password, err := bcrypt.GenerateFromPassword(convertPwd, bcrypt.DefaultCost)
	o.Password = password

	if err != nil {
		return err
	}

	return nil
}

func operatorCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("Operator")

	return coll
}

// OperatorsAll godoc
func OperatorsAll() (list []Operator, err error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{
		{Key: "deleted_at", Value: nil},
	}

	withoutFields := bson.D{
		{Key: "deleted_at", Value: 0},
		{Key: "Password", Value: 0},
	}

	findOptions := options.Find()
	findOptions.SetProjection(withoutFields)

	cur, err := operatorCollection().Find(timeout, filter, findOptions)
	defer cur.Close(timeout)

	if err != nil {
		return nil, err
	}

	for cur.Next(timeout) {
		var resault Operator
		err := cur.Decode(&resault)

		if err != nil {
			return nil, err
		}

		list = append(list, resault)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func operatorFindOne(id primitive.ObjectID) (operator *Operator, err error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.D{
		{Key: "deleted_at", Value: nil},
		{Key: "_id", Value: id},
	}
	withoutFields := bson.D{
		{Key: "deleted_at", Value: 0},
		{Key: "Password", Value: 0},
	}

	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(withoutFields)

	err = operatorCollection().FindOne(timeout, filter, findOneOptions).Decode(&operator)

	if err != nil {
		return nil, err
	}

	return operator, nil
}

// OperatorOne godoc
func OperatorOne(id string) (operator *Operator, err error) {
	operatorID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	operator, err = operatorFindOne(operatorID)

	if err != nil {
		return nil, err
	}

	return operator, nil
}

// OperatorDelete godoc
func OperatorDelete(id string) (operator *Operator, err error) {
	operatorID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{
		{
			Key:   "deleted_at",
			Value: nil,
		},
		{
			Key:   "_id",
			Value: operatorID,
		},
	}

	deleteSet := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "deleted_at",
					Value: time.Now(),
				},
			},
		},
	}

	operator, err = OperatorOne(id)

	if err != nil {
		return nil, err
	}

	_, err = operatorCollection().UpdateOne(timeout, filter, deleteSet)

	if err != nil {
		return nil, err
	}

	return operator, nil
}

// AddOperator godoc
func AddOperator(operator Operator) (*Operator, error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	operator.DeletedAt = nil

	insertedResault, err := operatorCollection().InsertOne(timeout, operator)

	if err != nil {
		return nil, err
	}

	resaultOperator, err := operatorFindOne(insertedResault.InsertedID.(primitive.ObjectID))

	if err != nil {
		return nil, err
	}

	return resaultOperator, nil
}

// OperatorUpdate godoc
func OperatorUpdate(id string, operator Operator) (*Operator, error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	operatorID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: operatorID,
		},
	}

	update := bson.D{
		{
			Key:   "$set",
			Value: operator,
		},
	}

	operatorCollection().UpdateOne(timeout, filter, update)

	resaultOperator, err := operatorFindOne(operatorID)

	if err != nil {
		return nil, err
	}

	return resaultOperator, nil
}

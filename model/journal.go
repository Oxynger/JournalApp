package model

import (
	"context"
	"log"
	"time"

	"github.com/Oxynger/JournalApp/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// BlockArray godoc
	BlockArray = []Block

	// Journal godoc
	Journal = map[string]interface{}
)

// JournalInfo описание обекта журнала
type JournalInfo struct {
	// Name название журнала
	Name string `bson:"name" json:"name" example:"scale_repair"`

	// ID идентификатор журнала
	ID interface{} `bson:"journal_id" json:"journal_id" example:"5c93e5621f23834a97aba93b"`

	// Daily является ли журнал ежедневным
	Daily bool `bson:"daily" json:"daily" example:"true"`

	// Accepted было ли завершено заполнение журнала сегодня
	Accepted bool `bson:"accepted" json:"accepted" example:"true"`
}

// ListJournalItems список объектов принадлежащих группе
type ListJournalItems struct {
	// Возможность добавления в журнал новых позиций
	Addition bool        `bson:"addition" json:"addition" example:"true"`
	Group    []ItemGroup `bson:"Group" json:"Group"`
}

// JournalItem объект принадлежащий группе
type JournalItem struct {
	Header CurrentItem `bson:"Header" json:"Header"`
	Blocks []Block
	Errors []Error
}

// Block godoc
type Block struct {
	ID        string     `bson:"id" json:"id" example:"1234"`
	Type      string     `bson:"type" json:"type" example:"text"`
	Name      string     `bson:"name" json:"name"`
	Value     *string    `bson:"value,omitempty" json:"value,omitempty"`
	Check     *Check     `bson:"check,omitempty" json:"check,omitempty"`
	Statement *string    `bson:"statement,omitempty" json:"statement,omitempty"`
	Image     *string    `bson:"image,omitempty" json:"image,omitempty"`
	Buttons   *[2]string `bson:"buttons,omitempty" json:"buttons,omitempty"`
	AID       *string    `bson:"a_id,omitempty" json:"a_id,omitempty"`
	BID       *string    `bson:"b_id,omitempty" json:"b_id,omitempty"`
	Enum      *[]string  `bson:"enum,omitempty" json:"enum,omitempty"`
}

// Error godoc
type Error struct {
	Block
}

// Check godoc
type Check struct {
	Type    string `bson:"type" json:"type" example:"deviation"`
	ShowNow bool   `bson:"show_now" json:"show_now" example:"true"`

	// Norm Если type deviation
	Norm      *float32 `bson:"norm,omitempty" json:"norm,omitempty" example:"10"`          // Нормальный вес
	Devetions *float32 `bson:"devetions,omitempty" json:"devetions,omitempty" example:"2"` // Допустимое отклонение

	// Range Если type range
	Range *[2]float32 `bson:"range,omitempty" json:"range,omitempty"` // Допустимы предел

	// Value Если type equals
	Value *float32 `bson:"value,omitempty" json:"value,omitempty" example:"0.2"` // Значение, которому должно быть равно

	// Max Если type less
	Max *float32 `bson:"max,omitempty" json:"max,omitempty" example:"1.2"` // Допустимое максимальное значение

	// Min Если type more
	Min *float32 `bson:"min,omitempty" json:"min,omitempty" example:"0.2"` // Допустимое минимальное значение

	//ID Если type more_than
	ID *string  `bson:"id,omitempty" json:"id,omitempty" example:"5c9396471f23834a97aba93a"` // Идентификатор блока
	On *float32 `bson:"on,omitempty" json:"on,omitempty" example:"3"`                        // Значение, на которое должен быть больше

	// Enum Если type enum
	// Массив допустимых значений. Если выбранное
	// значение не принадлежит данному массиву, то
	// check = false
	Enum *[]string `bson:"enum,omitempty" json:"enum,omitempty"`
}

// journalCollection godoc
func journalCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("Journal")

	return coll
}

// ListJournals godoc
func ListJournals() ([]Journal, error) {
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var List []Journal

	cur, err := journalCollection().Find(timeout, bson.D{})

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

		List = append(List, resault)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return List, nil
}

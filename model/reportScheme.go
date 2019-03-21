package model

import (
	"context"
	"log"

	"../db"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportField struct {
	Title string `bson:"title" json:"title" example:"Дата"`
	Value string `bson:"value" json:"value" example:"{journal.date}"`
}

type ReportScheme struct {
	Name    string        `bson:"name" json:"name" example:"scales_calibration"`
	Title   string        `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Journal string        `bson:"journal" json:"journal" example:"scales_calibration"`
	Fields  []ReportField `bson:"fields" json:"fields"`
}

func ReportSchemeCollection() *mongo.Collection {
	client := db.Client()
	coll := client.Database("test").Collection("reportScheme")

	return coll
}

func SomeReportSchemeAdd() ReportScheme {
	scalesCalibrationScheme := ReportScheme{
		Name:    "scale",
		Title:   "Весы",
		Journal: "scales_calibration",
		Fields: []ReportField{
			{
				Title: "Дата",
				Value: "{journal.date}",
			},
			{
				Title: "Инвентарный номер",
				Value: "{object.inventory_number}",
			},
			{
				Title: "Тип весов",
				Value: "{object.name}\nmin={object.min_w}, max={object.max_w}, e={object.deviation}",
			},
			{
				Title: "Заводской номер",
				Value: "{object.serial_number}",
			},
		},
	}

	insertResault, err := ReportSchemeCollection().InsertOne(context.Background(), scalesCalibrationScheme)

	if err != nil {
		log.Println(err)
		return ReportScheme{}
	}

	log.Println("Inserted documents: ", insertResault.InsertedID)

	return scalesCalibrationScheme
}

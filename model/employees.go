package model

import "time"

type Employees struct {
	MidlleName        string    `bson:"midname" json:"midname" example:"Сидоров"`
	Name              string    `bson:"name" json:"name" example:"Петр"`
	Surname           string    `bson:"surname" json:"surname" example:"Петровичь"`
	Function          string    `bson:"function" json:"function" example:"Повар"`
	CertificationDate time.Time `bson:"certification_date" json:"certification_date" example:"2019-03-14T23:08:14.586Z"`
	MedexamDate       time.Time `bson:"medexam_date" json:"medexam_date" example:"2019-03-14T23:08:14.586Z"`
	BirthYear         int       `bson:"birth_year" json:"birth_year" example:"1970"`
	Address           string    `bson:"address" json:"address" example:"пр. Ленина 78, кв. 13"`
	Lmk               int64     `bson:"lmk" json:"lmk" example:"882391949" format:"int64"`
	RoundHologram     string    `bson:"round_hologram" json:"round_hologram" example:"МП96654213"`
	SquareHologram    string    `bson:"square_hologram" json:"square_hologram" example:"А21546879"`
}

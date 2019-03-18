package model

type Operators struct {
	Surname    string `bson:"surname" json:"surname" examlpe:"Иванова"`
	Name       string `bson:"name" json:"name" example:"Александра"`
	MiddleName string `bson:"middlename" json:"middlename" example:"Анатольева"`
}

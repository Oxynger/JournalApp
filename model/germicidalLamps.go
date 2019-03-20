package model

// GermicidalLamps godoc
type GermicidalLamps struct {
	Name         string `bson:"name" json:"name" example:"Термометр стекл. жидк."`
	SerialNumber string `bson:"_id,omitempty" json:"serial_number" example:"Облучатель - рециркулятор воздуха ультрафиолетовый бактерицидный" format:"string"`
}

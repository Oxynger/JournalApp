package model

import "time"

// Hydrometers godoc
type Hydrometers struct {
	Name                 string    `bson:"name" json:"name" example:"Термометр стекл. жидк."`
	SerialNumber         string    `bson:"_id,omitempty" json:"serial_number" example:"ТСЖ-Х № 1879з" format:"string"`
	VerificationDate     time.Time `bson:"verifications_date" json:"verifications_date" example:"2019-03-14T23:08:14.586Z"`
	NextVerificationDate time.Time `bson:"next_nerification_date" json:"next_nerification_date" example:"2019-03-14T23:08:14.586Z"`
	Bailee               string    `bson:"bailee" json:"bailee" example:"Стрельникова А.А."`
}

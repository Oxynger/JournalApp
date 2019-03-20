package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ResaultInTimeDuration godoc
type ResaultInTimeDuration struct {
	Payload gin.H

	Start time.Time `bson:"start" json:"start" example:"2019-03-14T23:08:14.586Z"`
	End   time.Time `bson:"end" json:"end" example:"2019-03-14T23:08:14.586Z"`
}

// PlacesMeasuring godoc
type PlacesMeasuring struct {
	Payload gin.H

	Places   Places                  `bson:"Places" json:"Places"`
	Operator Operators               `bson:"Operators" json:"Operators"`
	Sinature Singatures              `bson:"Singatures" json:"Singatures"`
	Date     time.Time               `bson:"date" json:"date" example:"2019-03-14T23:08:14.586Z"`
	Resaults []ResaultInTimeDuration `bson:"resaults" json:"resaults"`
}

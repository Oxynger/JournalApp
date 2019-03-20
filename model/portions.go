package model

// Portions godoc
type Portions struct {
	// В граммах
	Weidth     int `bson:"weidth" json:"weidth" example:"100"`
	MinWeidth  int `bson:"min_w" json:"min_w" example:"100"`
	MaxWeidth  int `bson:"max_w" json:"max_w" example:"150"`
	PackWeidth int `bson:"pack_w" json:"pack_w" example:"16.34"`

	// В °C

	MinTemperature int `bson:"min_t" json:"min_t" example:"1.2"`
	MaxTemperature int `bson:"max_t" json:"max_t" example:"5.2"`
}

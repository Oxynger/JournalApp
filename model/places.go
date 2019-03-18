package model

type Places struct {
	Number string `bson:"_id,omitempty" json:"number" example:"1" format:"string"`
	Name   string `bson:"name" json:"name" example:"Холодильник"`

	// MinHumidity влажность в %
	MinHumidity int `bson:"min_h" json:"min_h" example:"60"`

	// MaxHumidity влажность в %
	MaxHumidity int `bson:"max_h" json:"max_h" example:"65"`

	// NormalTemperature температура в °C
	NormalTemperature int `bson:"normal_t" json:"normal_t" example:"4"`

	// DeviationTemperature погрешность ± температуры в °C
	DeviationTemperature int `bson:"deviations_t" json:"deviations_t" example:"2"`
}

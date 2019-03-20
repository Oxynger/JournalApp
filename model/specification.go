package model

// Specification godoc
type Specification struct {
	Code                   string `bson:"_id,omitempty" json:"code" example:"22" format:"string"`
	Name                   string `bson:"name" json:"name" example:"Огурцы маринованные «Sun Feel»"`
	Provider               string `bson:"provider" json:"provider" example:"Индиан групп"`
	Brand                  string `bson:"brand" json:"brand" example:"Индиан тропикан Агро Продактс"`
	SecondaryControlMethod string `bson:"secondary_control_method" json:"secondary_control_method" example:"Полное описание метода вторичного контроля"`
}

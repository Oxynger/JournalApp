package model

// Item описание объекта
type Item struct {
	// Идентификатор позиции
	ItemID string `bson:"_id" json:"item_id" example:"5c93e5621f23834a97aba93b"`

	// Название позиции
	Name string `bson:"name" json:"name" example:"scale"`

	// Было ли завершено заполнение позиции сегодня -1 возвращается если было завершено, но с корректирующими действиями (может отсутствовать)
	Accepted *int `bson:"accepted,omitempty" json:"accepted,omitempty" example:"-1"`
}

// ItemGroup описание группы объектов
type ItemGroup struct {
	Name  string `bson:"name" json:"name" example:"Салатный цех"`
	Items []Item `bson:"items" json:"items"`
}

// VarItem godoc
type VarItem struct {
	Name  string `bson:"name" json:"name" example:"max_w"`
	Value string `bson:"value" json:"value" example:"2"`
}

// CurrentItem godoc
type CurrentItem struct {
	Name   string    `bson:"name" json:"name" example:"scale"`
	Fields []VarItem `bson:"fields" json:"fields"`
}

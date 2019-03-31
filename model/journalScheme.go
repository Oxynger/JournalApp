package model

// JournalComputed godoc
type JournalComputed struct {
	Type  string `bson:"type" json:"type" example:"deviation"`
	Field string `bson:"field" json:"field" example:"result"`

	// Norm Если type deviation
	Norm      *float32 `bson:"norm,omitempty" json:"norm,omitempty" example:"10"`          // Нормальный вес
	Devetions *float32 `bson:"devetions,omitempty" json:"devetions,omitempty" example:"2"` // Допустимое отклонение

	// Range Если type range
	Range *[2]float32 `bson:"range,omitempty" json:"range,omitempty" example:"[ 1.1 , 4 ]"` // Допустимы предел

	// Value Если type equals
	Value *float32 `bson:"value,omitempty" json:"value,omitempty" example:"0.2"` // Значение, которому должно быть равно

	// Max Если type less
	Max *float32 `bson:"max,omitempty" json:"max,omitempty" example:"1.2"` // Допустимое максимальное значение

	// Min Если type more
	Min *float32 `bson:"min,omitempty" json:"min,omitempty" example:"0.2"` // Допустимое минимальное значение

	// ID Если type more_than
	ID *string  `bson:"id,omitempty" json:"id,omitempty" example:"5c9396471f23834a97aba93a"` // Идентификатор блока
	On *float32 `bson:"on,omitempty" json:"on,omitempty" example:"3"`                        // Значение, на которое должен быть больше

	// Enum Если type enum
	// Массив допустимых значений. Если выбранное
	// значение не принадлежит данному массиву, то
	// check = false
	Enum *[]string `bson:"enum,omitempty" json:"enum,omitempty" example:"0.2"`
}

// JournalField godoc
type JournalField struct {
	Name  string `bson:"name" json:"name" example:"serial_number"`
	Title string `bson:"title" json:"title" example:"Серийный номер"`
	Type  string `bson:"type" json:"type" example:"String"`

	// Computed вычесляемое поле с переменным количеством полей
	Computed *JournalComputed `bson:"computed,omitempty" json:"computed,omitempty"`

	// If непонятно какое условие

}

// JournalScheme godoc
type JournalScheme struct {
	Name    string         `bson:"name" json:"name" example:"scales_calibration"`
	Title   string         `bson:"title" json:"title" example:"Учет и калибровка весов"`
	Daily   bool           `bson:"daily" json:"daily" example:"true"`
	Fixed   bool           `bson:"fixed" json:"fixed" example:"true"`
	Items   []ItemInfo     `bson:"items" json:"items"`
	Fields  []JournalField `bson:"fields" json:"fields"`
	Deleted bool           `bson:"deleted" json:"-"`
}

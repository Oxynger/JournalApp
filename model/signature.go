package model

// Signature Хранит изображение росписи
type Signature struct {
	// Code Закодированное изображение
	Code int64 `bson:"signature" json:"signature"`
}

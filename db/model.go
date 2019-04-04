package db

import (
	"time"
)

// Model базовая модел для всех структур который будут записанны в бд.
// Пример:
// type Example struct {
// 	db.Model
// 	someField some
// }
type Model struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	// DeletedAt может быть nil. В таком случае считается что объект не удален
	DeletedAt *time.Time `bson:"deleted_at" json:"deleted_at"`
}

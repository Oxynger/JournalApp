package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model базовая модел для всех структур который будут записанны в бд.
// Пример:
// type Example struct {
// 	db.Model
// 	someField some
// }
type Model struct {
	ID        primitive.ObjectID `bson:"_id" json:"ID" example:"5ca10d9d015c736a72b7b3ba"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`

	// DeletedAt может быть nil. В таком случае считается что объект не удален
	DeletedAt *time.Time `bson:"deleted_at" json:"deleted_at"`
}

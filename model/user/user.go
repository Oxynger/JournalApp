package user

//go:generate stringer -type Role

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role int

const (
	Operator Role = iota
	Administrator
	Helpdesk
)

type User struct {
	ID       *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Username string              `bson:"username" json:"username"`
	Password string              `bson:"password" json:"password"`
	Role     Role                `bson:"role" json:"role"`
}

func (usr User) MarshalJSON() ([]byte, error) {
	temp := struct {
		Username string `json:"username"`
		Role     Role   `json:"role"`
	}{
		usr.Username,
		usr.Role,
	}

	return json.Marshal(&temp)
}

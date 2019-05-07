package user

//go:generate stringer -type Role

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

type Role int

const (
	Anonymous Role = iota
	Administrator
	Constructor
	Operator
	Helpdesk
)

type User struct {
	ID       uuid.UUID `bson:"userid,omitempty" json:"ID" `
	Username string    `bson:"username" json:"username"`
	Password string    `bson:"password" json:"password"`
	Role     Role      `bson:"role" json:"role"`
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

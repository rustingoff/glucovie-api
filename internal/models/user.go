package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" binding:"email,required"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty" binding:"required"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" binding:"required"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("invalid email address")
	}
	return nil
}

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
	Gender   string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Age      uint8              `json:"age,omitempty" bson:"age,omitempty"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("invalid email address")
	}
	return nil
}

type UserCredentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u UserCredentials) GetMail() string {
	return u.Email
}

func (u UserCredentials) GetPassword() string {
	return u.Password
}

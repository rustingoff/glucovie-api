package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EventModel struct {
	Title  string `json:"title"`
	From   string `json:"from"`
	To     string `json:"to"`
	UserID string `json:"user_id"`
}

type EventResponse struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title"`
	From  string             `json:"from"`
	To    string             `json:"to"`
}

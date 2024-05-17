package models

import "time"

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username"`
}

type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Expense struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"roomId"`
	DaddyID   string    `json:"daddyId"`
	Currency  string    `json:"currency"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
}

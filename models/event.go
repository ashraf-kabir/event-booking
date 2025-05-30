package models

import "time"

type Event struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

var events []Event

func (e Event) SaveEvent() {
	// Todo: add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

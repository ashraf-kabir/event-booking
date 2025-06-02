package models

import (
	"event-booking/db"
	"time"
)

type Event struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var events []Event

func (e Event) SaveEvent() error {
	query := `
		INSERT INTO events (user_id, name, description, location)
			VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.UserId, e.Name, e.Description, e.Location)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.UserId, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UpdatedAt)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int) (*Event, error) {
	query := `SELECT * FROM events WHERE id=?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.UserId, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

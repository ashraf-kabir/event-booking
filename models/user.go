package models

import (
	"event-booking/db"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) SaveUser() error {
	query := `INSERT INTO USERS(first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.Id = id
	return err
}

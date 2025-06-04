package models

import (
	"errors"
	"event-booking/db"
	"event-booking/utils"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) SaveUser() error {
	query := `INSERT INTO USERS(email, password, created_at, updated_at) VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.Id = id
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string
	err := row.Scan(&u.Id, &retrievePassword)

	if err != nil {
		return err
	}

	isValidPassword := utils.CheckPasswordHash(u.Password, retrievePassword)

	if !isValidPassword {
		return errors.New("invalid credentials")
	}

	return nil
}

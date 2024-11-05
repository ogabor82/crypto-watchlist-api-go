package models

import (
	"crypto-watchlist-api/db"
	"time"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
}

func CreateUser(email, password, name string) (*User, error) {
	// Generate a simple ID (in real app, use UUID or similar)
	id := "user_" + email

	// Insert user into database
	query := `
		INSERT INTO users (id, email, password, name)
		VALUES (?, ?, ?, ?)
	`

	_, err := db.DB.Exec(query, id, email, password, name)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       id,
		Email:    email,
		Password: password,
		Name:     name,
	}

	return user, nil
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, password, name, created_at FROM users WHERE email = ?`

	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

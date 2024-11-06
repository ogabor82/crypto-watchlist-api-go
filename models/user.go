package models

import (
	"crypto-watchlist-api/db"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
}

func CreateUser(email, password, name string) (*User, error) {
	// Generate UUID for the user
	id := uuid.New().String()

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Insert user into database
	query := `
		INSERT INTO users (id, email, password, name)
		VALUES (?, ?, ?, ?)
	`

	_, err = db.DB.Exec(query, id, email, string(hashedPassword), name)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       id,
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	return user, nil
}

func VerifyCredentials(email, password string) (*User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
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

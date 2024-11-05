package models

type User struct {
	ID       string
	Email    string
	Password string // In a real app, this would be hashed
	Name     string
}

// CreateUser creates a new user (in-memory for now)
func CreateUser(email, password, name string) (*User, error) {
	// Generate a simple ID (in real app, use UUID or similar)
	user := &User{
		ID:       "user_" + email, // simplified ID generation
		Email:    email,
		Password: password, // In real app, hash the password
		Name:     name,
	}
	return user, nil
} 
package models

import (
	"database/sql"
	"time"
)

// User is a model that's used to interact with the 'users' table
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// UserModel is a wrapper around the database to provide encapsulation.
// Encapsulation here means 'to not spread raw sql db access everywhere'.
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new user record to the database
func (um *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate method verifies whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (um *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exists checks wether user is prevalent in the database
func (um *UserModel) Exists(id int) (bool, error) {
	return false, nil
}

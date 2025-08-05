package models

import "errors"

var (
	// This error should be returned when no matching record has been found in the database.
	// It exists for full model encapsulation, so the handlers aren't concerned with the
	// underlying datastore or reliant on storage specific errors (like sql.ErrNoRows).
	// It's value is "models: no matching record found".
	ErrNoRecord = errors.New("models: no matching record found")

	// Use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// Use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

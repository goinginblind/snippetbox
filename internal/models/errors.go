package models

import "errors"

// This error should be returned when no matching record has been found in the database.
// It exists for full model encapsulation, so the handlers aren't concerned with the
// underlying datastore or reliant on storage specific errors (like sql.ErrNoRows).
// It's value is "models: no matching record found".
var ErrNoRecord = errors.New("models: no matching record found")

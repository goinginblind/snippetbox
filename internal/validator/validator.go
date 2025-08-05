package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

// EmailRX stores regular expression for email sanity checks.
// It's defined in the validator package to save perfomance: compiled only once.
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Define a new Validator struct which contains a map of validation error messages
// for our form fields
type Validator struct {
	// FieldErrors holds a specific field name as a map key,
	// and a message to be displayed for the user as value.
	FieldErrors map[string]string
	// NonFieldErrors holds errors that are not specific for
	// specific fields, i.e. an incorrect email-password
	// combination
	NonFieldErrors []string
}

// Valid() returns true if the FieldErrors map doesn't contain any entries
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

// AddFieldError() adds an error message to the FieldErrors map (so long as no
// entry already exists for the given key).
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, ok := v.FieldErrors[key]; !ok {
		v.FieldErrors[key] = message
	}
}

// AddNonFieldError() adds an error message to the NonFiledErrors slice
// of strings with a given message.
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

// CheckField() adds an error message to the FieldErrors map only if a
// validation check is not 'ok'.
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank() returns true if a value is not an empty string.
func NotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}

// MaxChars() returns true if a value contains no more (<=) than n characters.
func MaxChars(s string, n int) bool {
	return utf8.RuneCountInString(s) <= n
}

// MinChars() returns true if a value contains no less (>=) than n characters.
func MinChars(s string, n int) bool {
	return utf8.RuneCountInString(s) >= n
}

// PermittedValue() returns true if a value is in a list of specific permitted
// values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

// Matches() returns true if a value matches a provided compiled regular
// expression pattern.
func Matches(s string, rx *regexp.Regexp) bool {
	return rx.MatchString(s)
}

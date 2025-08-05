package main

type contextKey string

// isAuthenticatedContextKey is a constant variable of a contextKey
// type string, it prevents naming collisions with other packages in case
// they have 'isAuthenticated' fields.
const isAuthenticatedContextKey = contextKey("isAuthenticated")

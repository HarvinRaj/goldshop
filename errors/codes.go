package errors

const (
	// ErrNotFound starting 1001 for application
	ErrNotFound = 1000

	// ErrNotFoundQuery starting 2001 for database
	ErrNotFoundQuery = 2000

	// IsFoundQuery starting 3000 for exist defaults
	IsFoundQuery = 3000
)

var ErrorMessages = map[int]string{
	ErrNotFound:      "Resource not found",
	ErrNotFoundQuery: "Item not found in database",
	IsFoundQuery:     "Item found in database",
}

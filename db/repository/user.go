package repository

// User represents a specific user account that may have a landing page with multiple links
// I make the assumption that a user will have a set of links that represent a single landing page.
// This may be a bit unrealistic as a user may have multiple landing pages and will want to organise links under that
// but it was done to keep the task simple.
type User struct {
	ID    int64 `db:"id"`
	Links []*Link
}

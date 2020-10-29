package request

// UserSearch specify search user information
type UserSearch struct {
	UserID      string
	FirstName   string
	LastName    string
	CountryCode string
	PaginationInfo
}

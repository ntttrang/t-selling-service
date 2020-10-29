package request

// PaginationInfo pagination information
type PaginationInfo struct {
	Skip      int
	PageSize  int
	SortBy    string
	Direction string
}

package entity

type Pagination struct {
	CurrentPage     int64 `json:"current_page"`
	CurrentElements int64 `json:"current_elements"`
	TotalPages      int64 `json:"total_pages"`
	TotalElements   int64 `json:"total_elements"`
}

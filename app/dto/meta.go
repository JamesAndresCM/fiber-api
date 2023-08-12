package dto

type Meta struct {
	CurrentPage    int     `json:"current_page"`
	TotalElements  int     `json:"total_elements"`
	TotalPages     float64 `json:"total_pages"`
	ObjectsPerPage int     `json:"objects_per_page"`
}

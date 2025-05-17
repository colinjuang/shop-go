package request

type CategoryRequest struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	ParentID  uint64 `json:"parentId"`
	Level     int    `json:"level"`
	ImageUrl  string `json:"imageUrl"`
	SortOrder int    `json:"sortOrder"`
}

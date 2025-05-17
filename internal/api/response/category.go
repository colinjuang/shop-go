package response

import "time"

type CategoryResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	ParentID  uint64    `json:"parentId"`
	ImageUrl  string    `json:"imageUrl"`
	SortOrder int       `json:"sortOrder"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CategoryTreeResponse struct {
	ID        uint64             `json:"id"`
	Name      string             `json:"name"`
	ImageUrl  string             `json:"imageUrl"`
	SortOrder int                `json:"sortOrder"`
	Children  []CategoryResponse `json:"children"`
}

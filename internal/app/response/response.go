package response

// Response represents a standard API response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse returns a success response
func SuccessResponse(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// ErrorResponse returns an error response
func ErrorResponse(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}

// TokenExpiredResponse returns a token expired response
func TokenExpiredResponse() Response {
	return Response{
		Code:    208,
		Message: "Token expired",
	}
}

// Pagination represents pagination information
type Pagination struct {
	Total       int64       `json:"total"`
	PageSize    int         `json:"page_size"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	Data        interface{} `json:"data"`
}

// NewPagination creates a new pagination response
func NewPagination(total int64, page, pageSize int, data interface{}) Pagination {
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return Pagination{
		Total:       total,
		PageSize:    pageSize,
		CurrentPage: page,
		TotalPages:  totalPages,
		Data:        data,
	}
}

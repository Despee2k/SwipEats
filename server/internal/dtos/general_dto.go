package dtos

type APISuccessResponse[T any] struct {
	Message string `json:"message"`
	Data   	T      `json:"data,omitempty"`
}

type APIErrorResponse struct {
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

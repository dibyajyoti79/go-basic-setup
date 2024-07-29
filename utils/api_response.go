package utils

type ApiResponse struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     []string    `json:"errors,omitempty"`
	Stack      string      `json:"stack,omitempty"`
}

func NewApiResponse(statusCode int, data interface{}, message string) *ApiResponse {
	if message == "" {
		message = "Success"
	}
	return &ApiResponse{
		StatusCode: statusCode,
		Success:    statusCode < 400,
		Message:    message,
		Data:       data,
	}
}

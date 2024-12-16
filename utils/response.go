package utils

type Response struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      any    `json:"error"`
}

func ApiResponse(StatusCode int, Data interface{}) Response {
	response := Response{
		StatusCode: StatusCode,
		Data:       Data,
	}

	return response
}

func ApiErrorResponse(StatusCode int, Message string, Error interface{}) ErrorResponse {
	errorResponse := ErrorResponse{
		StatusCode: StatusCode,
		Message:    Message,
		Error:      Error,
	}

	return errorResponse
}

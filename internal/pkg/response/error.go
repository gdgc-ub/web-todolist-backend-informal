package response

type ErrorResponse struct {
	Code    int    `json:"-"`
	Message string `json:"error"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{Code: code, Message: message}
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

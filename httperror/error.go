package httperror

// New Конструктор ошибки
func New(status int, err error) HTTPError {
	return HTTPError{
		Code:    status,
		Message: err.Error(),
	}
}

// HTTPError Объект ошибки
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

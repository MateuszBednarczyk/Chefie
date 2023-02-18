package services

type ServiceResponse struct {
	Message string
	Code    int
	Content []interface{}
}

func NewServiceResponse(message string, code int, content []interface{}) *ServiceResponse {
	return &ServiceResponse{
		Message: message,
		Code:    code,
		Content: content,
	}
}

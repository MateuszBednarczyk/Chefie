package handlers

import "back/src/pkg/services"

type HandlerResponse struct {
	Message string
	Content []interface{}
}

func NewHandlerResponse(serviceResponse *services.ServiceResponse) *HandlerResponse {
	return &HandlerResponse{
		Message: serviceResponse.Message,
		Content: serviceResponse.Content,
	}
}

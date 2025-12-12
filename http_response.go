package sgwf

import "net/http"

type HttpResponse struct {
	StatusCode int
	Body       interface{}
	Error      error
}

type Message struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewErrorMessage(status int, err error, msg string) *HttpResponse {
	return &HttpResponse{
		StatusCode: status,
		Body: &Message{
			Error: msg,
		},
		Error: err,
	}
}

func NewOkResponse() *HttpResponse {
	return &HttpResponse{
		StatusCode: http.StatusOK,
		Body:       nil,
		Error:      nil,
	}
}

func NewOkBodyResponse(body interface{}) *HttpResponse {
	return &HttpResponse{
		StatusCode: http.StatusOK,
		Body:       body,
		Error:      nil,
	}
}

func NewErrorResponse(status int, err error) *HttpResponse {
	return &HttpResponse{
		StatusCode: status,
		Body:       nil,
		Error:      err,
	}
}

func NewInternalServerErrorResponse(err error) *HttpResponse {
	return &HttpResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       nil,
		Error:      err,
	}
}

func NewBadRequestErrorResponse(err error) *HttpResponse {
	return &HttpResponse{
		StatusCode: http.StatusBadRequest,
		Body:       nil,
		Error:      err,
	}
}

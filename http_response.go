package sgwf

import "net/http"

type HttpResponse struct {
	StatusCode int
	Body       interface{}
	Error      error
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

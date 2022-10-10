package adapter

import (
	"net/http"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/adapter/rest"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/entity"
)

type RESTAdapter interface {
	ParseRequestBody(r *http.Request, v interface{}) *entity.AppError

	ResponseAppError(w http.ResponseWriter, r *http.Request, appErr *entity.AppError)
	ResponseAccepted(w http.ResponseWriter, r *http.Request, v interface{})
	ResponseCreated(w http.ResponseWriter, r *http.Request, v interface{})
	ResponseNoContent(w http.ResponseWriter, r *http.Request)
	ResponseOK(w http.ResponseWriter, r *http.Request, v interface{})
}

func NewRESTAdapter() RESTAdapter {
	return rest.NewAdapter()
}

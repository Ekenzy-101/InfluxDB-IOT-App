package rest

import (
	"log"
	"net/http"
	"strings"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/entity"
	"github.com/go-chi/render"
)

func NewAdapter() *adapter {
	return &adapter{}
}

func (a *adapter) ResponseAccepted(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, http.StatusAccepted)
	render.JSON(w, r, v)
}

func (a *adapter) ResponseCreated(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, v)
}

func (a *adapter) ResponseAppError(w http.ResponseWriter, r *http.Request, appErr *entity.AppError) {
	render.Status(r, appErr.HTTPCode)
	render.JSON(w, r, render.M{
		"error": appErr.Message,
		"code":  appErr.InternalCode,
	})
	if appErr.Error != nil {
		log.Println(appErr.Error)
	}
}

func (a *adapter) ResponseNoContent(w http.ResponseWriter, r *http.Request) {
	render.NoContent(w, r)
}

func (a *adapter) ResponseOK(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, http.StatusOK)
	if strings.EqualFold(r.Header.Get("Accept"), "text/csv") {
		w.Header().Add("Content-Type", "text/csv")
		w.Write([]byte(v.(string)))
	} else {
		render.JSON(w, r, v)
	}
}

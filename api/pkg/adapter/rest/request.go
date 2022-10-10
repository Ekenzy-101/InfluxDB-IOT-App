package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/entity"
)

func (a *adapter) ParseRequestBody(r *http.Request, v interface{}) *entity.AppError {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return entity.NewAppError(entity.ErrorWhenParsingRequestBody).
			SetError(err).
			SetHTTPCode(http.StatusBadRequest).
			SetMessage("Request body is malformed")
	}

	return nil
}

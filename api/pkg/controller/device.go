package controller

import (
	"net/http"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/adapter"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/entity"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/usecase"
	"github.com/go-chi/chi/v5"
)

type deviceController struct {
	deviceUc    usecase.DeviceUsecase
	restAdapter adapter.RESTAdapter
}

func NewDeviceController(deviceUc usecase.DeviceUsecase) Controller {
	return &deviceController{
		deviceUc:    deviceUc,
		restAdapter: adapter.NewRESTAdapter(),
	}
}

func (ctrl *deviceController) Register(router chi.Router) {
	router.Post("/devices", ctrl.CreateDevice)
	router.Get("/devices", ctrl.GetDevices)
	router.Get("/devices/{deviceID}", ctrl.GetDevice)
	router.Post("/devices/generate", ctrl.GenerateDevicesMeasurements)
	router.Post("/devices/{deviceID}/measurements", ctrl.GetDeviceMeasurements)
}

func (ctrl *deviceController) CreateDevice(w http.ResponseWriter, r *http.Request) {
	body := map[string]string{}
	if appErr := ctrl.restAdapter.ParseRequestBody(r, &body); appErr != nil {
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	if appErr := ctrl.deviceUc.CreateDevice(r.Context(), body["deviceId"]); appErr != nil {
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	ctrl.restAdapter.ResponseNoContent(w, r)
}

func (ctrl *deviceController) GetDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "deviceID")
	if deviceID == "" {
		appErr := entity.NewAppError(entity.ErrorWhenParsingRequestPathParams).
			SetHTTPCode(http.StatusBadRequest).SetMessage("Path param 'deviceID' is required")
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	device, appErr := ctrl.deviceUc.GetDevice(r.Context(), deviceID)
	if appErr != nil {
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	ctrl.restAdapter.ResponseOK(w, r, device)
}

func (ctrl *deviceController) GetDevices(w http.ResponseWriter, r *http.Request) {
	devices, appErr := ctrl.deviceUc.GetDevices(r.Context())
	if appErr != nil {
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	ctrl.restAdapter.ResponseOK(w, r, devices)
}

func (ctrl *deviceController) GetDeviceMeasurements(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "deviceID")
	if deviceID == "" {
		appErr := entity.NewAppError(entity.ErrorWhenParsingRequestPathParams).
			SetHTTPCode(http.StatusBadRequest).SetMessage("Path param 'deviceID' is required")
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	body := map[string]string{}
	if appErr := ctrl.restAdapter.ParseRequestBody(r, &body); appErr != nil {
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	devices, appErr := ctrl.deviceUc.GetDeviceMeasurements(r.Context(), deviceID, body["query"])
	if appErr != nil {
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	ctrl.restAdapter.ResponseOK(w, r, devices)
}

func (ctrl *deviceController) GenerateDevicesMeasurements(w http.ResponseWriter, r *http.Request) {
	body := map[string][]string{}
	if appErr := ctrl.restAdapter.ParseRequestBody(r, &body); appErr != nil {
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	deviceIDs, ok := body["deviceIds"]
	if !ok || len(deviceIDs) == 0 {
		appErr := entity.NewAppError(entity.ErrorWhenParsingRequestBody).
			SetHTTPCode(http.StatusUnprocessableEntity).SetMessage("Field 'deviceIds' is required")
		ctrl.restAdapter.ResponseAppError(w, r, appErr)
		return
	}

	for _, deviceID := range deviceIDs {
		if appErr := ctrl.deviceUc.CreateDeviceMeasurement(r.Context(), deviceID); appErr != nil {
			ctrl.restAdapter.ResponseAppError(w, r, appErr)
			return
		}
	}

	ctrl.restAdapter.ResponseNoContent(w, r)
}

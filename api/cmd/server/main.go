package main

import (
	"net/http"
	"os"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/config"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/controller"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/repository"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/usecase"
	"github.com/go-chi/chi/v5"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	router := chi.NewRouter()
	client := influxdb2.NewClient(config.InfluxAPIURL(), config.InfluxAPIToken())

	deviceRepo := repository.NewDeviceRepository(client)
	measurementRepo := repository.NewMeasurementRepository(client)

	deviceUc := usecase.NewDeviceUsecase(deviceRepo, measurementRepo)

	controller.NewDeviceController(deviceUc).Register(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	http.ListenAndServe(":"+port, router)
}

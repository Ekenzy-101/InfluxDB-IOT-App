package repository

import (
	"context"
	"time"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/config"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/entity"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type MeasurementRepository interface {
	Create(ctx context.Context, deviceID string) error
	List(ctx context.Context, query string) (string, error)
}

type measurementRepository struct {
	client influxdb2.Client
}

func NewMeasurementRepository(client influxdb2.Client) MeasurementRepository {
	return &measurementRepository{
		client: client,
	}
}

func (repo *measurementRepository) Create(ctx context.Context, deviceID string) error {
	virtualDevice := entity.NewSensor()
	coordinate := virtualDevice.Geo()
	point := write.NewPointWithMeasurement("environment").
		AddTag("deviceId", deviceID).
		AddTag("temperatureSensor", virtualDevice.Name()).
		AddTag("humiditySensor", virtualDevice.Name()).
		AddTag("pressureSensor", virtualDevice.Name()).
		AddField("temperature", virtualDevice.Temperature()).
		AddField("humidity", virtualDevice.Humidity()).
		AddField("pressure", virtualDevice.Pressure()).
		AddField("latitude", coordinate.Latitude).
		AddField("longitude", coordinate.Longitude).
		SetTime(time.Now().UTC())
	return repo.client.
		WriteAPIBlocking(config.InfluxOrganizationID(), config.InfluxBucketID()).
		WritePoint(ctx, point)
}

func (repo *measurementRepository) List(ctx context.Context, query string) (string, error) {
	result, err := repo.client.QueryAPI(config.InfluxOrganizationID()).QueryRaw(ctx, query, influxdb2.DefaultDialect())
	if err != nil {
		return "", err
	}
	return result, nil
}

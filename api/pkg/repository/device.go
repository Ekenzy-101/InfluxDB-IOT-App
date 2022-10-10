package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/config"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/entity"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2/domain"
)

type DeviceRepository interface {
	Create(ctx context.Context, deviceID string) error
	Get(ctx context.Context, deviceID string) (*entity.Device, error)
	List(ctx context.Context) ([]*entity.Device, error)
}

type deviceRepository struct {
	client influxdb2.Client
}

func NewDeviceRepository(client influxdb2.Client) DeviceRepository {
	return &deviceRepository{
		client: client,
	}
}

func (repo *deviceRepository) Create(ctx context.Context, deviceID string) error {
	organizationID := config.InfluxOrganizationID()
	bucketID := config.InfluxBucketID()
	organizationResource := domain.Resource{
		Id:    &bucketID,
		Type:  domain.ResourceTypeBuckets,
		OrgID: &organizationID,
	}
	permissions := []domain.Permission{
		{
			Action:   domain.PermissionActionRead,
			Resource: organizationResource,
		},
		{
			Action:   domain.PermissionActionWrite,
			Resource: organizationResource,
		},
	}
	auth, err := repo.client.AuthorizationsAPI().CreateAuthorizationWithOrgID(ctx, organizationID, permissions)
	if err != nil {
		return err
	}

	point := write.NewPointWithMeasurement("deviceauth").
		AddTag("deviceId", deviceID).
		AddField("key", *auth.Id).
		AddField("token", *auth.Token).
		SetTime(time.Now().UTC())
	return repo.client.
		WriteAPIBlocking(organizationID, config.InfluxAuthBucketName()).
		WritePoint(ctx, point)
}

func (repo *deviceRepository) Get(ctx context.Context, deviceID string) (*entity.Device, error) {
	query := `from(bucket:"%v")
      |> range(start: 0)
      |> filter(fn: (r) => r._measurement == "deviceauth" and r.deviceId == "%v")
			|> last()
      `
	result, err := repo.client.
		QueryAPI(config.InfluxOrganizationID()).
		Query(ctx, fmt.Sprintf(query, config.InfluxAuthBucketName(), deviceID))
	if err != nil {
		return nil, err
	}
	if err := result.Err(); err != nil {
		return nil, err
	}

	var device *entity.Device
	for result.Next() {
		if device == nil {
			device = &entity.Device{ID: deviceID}
		}
		switch field := result.Record().Field(); field {
		case "key":
			device.Key = result.Record().Value().(string)
		case "token":
			// device.Token = result.Record().Value().(string)
		case "updatedAt":
			device.UpdatedAt = result.Record().Value().(time.Time)
		default:
			fmt.Println("Unexpected", field)
		}
		if device.UpdatedAt.IsZero() {
			device.UpdatedAt = result.Record().Time()
		}
	}
	return device, nil
}

func (repo *deviceRepository) List(ctx context.Context) ([]*entity.Device, error) {
	queryClient := repo.client.QueryAPI(config.InfluxOrganizationID())
	query := `from(bucket:"%v")
      |> range(start: 0)
      |> filter(fn: (r) => r._measurement == "deviceauth")
      |> last()`
	result, err := queryClient.Query(ctx, fmt.Sprintf(query, config.InfluxAuthBucketName()))
	if err != nil {
		return nil, err
	}

	devices := []*entity.Device{}
	deviceIndexes := map[string]int{}

	for result.Next() {
		deviceID := result.Record().ValueByKey("deviceId").(string)
		index, ok := deviceIndexes[deviceID]
		var device *entity.Device
		if ok {
			device = devices[index]
		} else {
			device = &entity.Device{ID: deviceID}
			deviceIndexes[deviceID] = len(devices)
			devices = append(devices, device)
		}

		switch field := result.Record().Field(); field {
		case "key":
			device.Key = result.Record().Value().(string)
		case "token":
			// device.Token = result.Record().Value().(string)
		case "updatedAt":
			device.UpdatedAt = result.Record().Value().(time.Time)
		default:
			fmt.Println("Unexpected", field)
		}
		if device.UpdatedAt.IsZero() {
			device.UpdatedAt = result.Record().Time()
		}
	}

	return devices, nil
}

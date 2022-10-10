package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ekenzy-101/Go-IOT-App/pkg/entity"
	"github.com/Ekenzy-101/Go-IOT-App/pkg/repository"
)

type DeviceUsecase interface {
	CreateDevice(ctx context.Context, deviceId string) *entity.AppError
	GetDevice(ctx context.Context, deviceId string) (*entity.Device, *entity.AppError)
	GetDevices(ctx context.Context) ([]*entity.Device, *entity.AppError)
	GetDeviceMeasurements(ctx context.Context, deviceID string, query string) (string, *entity.AppError)
	CreateDeviceMeasurement(ctx context.Context, deviceID string) *entity.AppError
}

type deviceUsecase struct {
	deviceRepo      repository.DeviceRepository
	measurementRepo repository.MeasurementRepository
}

func NewDeviceUsecase(
	deviceRepo repository.DeviceRepository,
	measurementRepo repository.MeasurementRepository,
) DeviceUsecase {
	return &deviceUsecase{
		deviceRepo:      deviceRepo,
		measurementRepo: measurementRepo,
	}
}

func (uc *deviceUsecase) CreateDevice(ctx context.Context, deviceId string) *entity.AppError {
	device, err := uc.deviceRepo.Get(ctx, deviceId)
	if err != nil {
		return entity.NewAppError(entity.ErrorWhenGettingDevice).SetError(err)
	}
	if device == nil {
		if err := uc.deviceRepo.Create(ctx, deviceId); err != nil {
			return entity.NewAppError(entity.ErrorWhenCreatingDevice).SetError(err)
		}
	}

	return nil
}

func (uc *deviceUsecase) GetDevice(ctx context.Context, deviceID string) (*entity.Device, *entity.AppError) {
	device, err := uc.deviceRepo.Get(ctx, deviceID)
	if err != nil {
		return nil, entity.NewAppError(entity.ErrorWhenGettingDevice).SetError(err)
	}
	if device == nil {
		return nil, entity.NewAppError(entity.ErrorWhenGettingDevice).SetHTTPCode(http.StatusNotFound).
			SetMessage(fmt.Sprintf("Device '%v' not found", deviceID))
	}

	return device, nil
}

func (uc *deviceUsecase) GetDevices(ctx context.Context) ([]*entity.Device, *entity.AppError) {
	devices, err := uc.deviceRepo.List(ctx)
	if err != nil {
		return nil, entity.NewAppError(entity.ErrorWhenGettingDevices).SetError(err)
	}

	return devices, nil
}

func (uc *deviceUsecase) CreateDeviceMeasurement(ctx context.Context, deviceID string) *entity.AppError {
	device, err := uc.deviceRepo.Get(ctx, deviceID)
	if err != nil {
		return entity.NewAppError(entity.ErrorWhenGettingDevice).SetError(err)
	}
	if device == nil {
		return entity.NewAppError(entity.ErrorWhenGettingDevice).SetHTTPCode(http.StatusUnprocessableEntity).
			SetMessage(fmt.Sprintf("Device '%v' not found", deviceID))
	}

	if err := uc.measurementRepo.Create(ctx, deviceID); err != nil {
		return entity.NewAppError(entity.ErrorWhenCreatingMeasurement).SetError(err)
	}

	return nil
}

func (uc *deviceUsecase) GetDeviceMeasurements(ctx context.Context, deviceID string, query string) (string, *entity.AppError) {
	device, err := uc.deviceRepo.Get(ctx, deviceID)
	if err != nil {
		return "", entity.NewAppError(entity.ErrorWhenGettingDevice).SetError(err)
	}
	if device == nil {
		return "", entity.NewAppError(entity.ErrorWhenGettingDevice).SetHTTPCode(http.StatusUnprocessableEntity).
			SetMessage(fmt.Sprintf("Device '%v' not found", deviceID))
	}

	measurements, err := uc.measurementRepo.List(ctx, query)
	if err != nil {
		return "", entity.NewAppError(entity.ErrorWhenGettingMeasurements).SetError(err)
	}

	return measurements, nil
}

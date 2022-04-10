package repository

import (
	"context"
	"smart_house_backend/internal/domain"
)

type UsersRepository interface {
	Get(ctx context.Context, id string) (user domain.User, err error)
	Create(ctx context.Context, user domain.User) (id string, err error)
	Update(ctx context.Context, user domain.User) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type ControllersRepository interface {
	Get(ctx context.Context, id string) (controller domain.Controller, err error)
	Create(ctx context.Context, controller domain.Controller) (id string, err error)
	Update(ctx context.Context, controller domain.Controller) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type ControllerTypes interface {
	Get(ctx context.Context, id string) (controller domain.Controller, err error)
	Create(ctx context.Context, controllerType domain.ControllerType) (id string, err error)
	Update(ctx context.Context, controllerType domain.ControllerType) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type HouseParts interface {
	Get(ctx context.Context, id string) (housePart domain.HousePart, err error)
	Create(ctx context.Context, housePart domain.HousePart) (id string, err error)
	Update(ctx context.Context, housePart domain.HousePart) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type Houses interface {
	Get(ctx context.Context, id string) (house domain.House, err error)
	Create(ctx context.Context, house domain.House) (id string, err error)
	Update(ctx context.Context, house domain.House) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type DeviceTypes interface {
	Get(ctx context.Context, id string) (deviceType domain.DeviceType, err error)
	Create(ctx context.Context, deviceType domain.DeviceType) (id string, err error)
	Update(ctx context.Context, deviceType domain.DeviceType) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type Devices interface {
	Get(ctx context.Context, id string) (deviceType domain.Device, err error)
	Create(ctx context.Context, deviceType domain.Device) (id string, err error)
	Update(ctx context.Context, deviceType domain.Device) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type Pins interface {
	Get(ctx context.Context, id string) (pin domain.Pin, err error)
	Create(ctx context.Context, pin domain.Pin) (id string, err error)
	Update(ctx context.Context, pin domain.Pin) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type Repository struct {
	Users           UsersRepository
	Controllers     ControllersRepository
	ControllerTypes ControllerTypes
	HouseParts      HouseParts
	Houses          Houses
	DeviceTypes     DeviceTypes
	Devices         Devices
	Pins            Pins
}

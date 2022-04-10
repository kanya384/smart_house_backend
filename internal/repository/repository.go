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

type ControllerTypesRepository interface {
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

type Repository struct {
	Users           UsersRepository
	Controllers     ControllersRepository
	ControllerTypes ControllerTypesRepository
	HouseParts      HouseParts
}

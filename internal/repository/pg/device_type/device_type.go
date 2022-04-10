package device_type

import (
	"context"
	"errors"
	"smart_house_backend/internal/domain"
	repo "smart_house_backend/internal/repository"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	Db *pgxpool.Pool
}

func NewRepository(Db *pgxpool.Pool) repo.DeviceTypes {
	return &repository{
		Db: Db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (deviceType domain.DeviceType, err error) {
	query, args, err := prepareGet(id)
	if err != nil {
		return deviceType, err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&deviceType.ID, &deviceType.Name)
	return
}

func (r *repository) Create(ctx context.Context, deviceType domain.DeviceType) (id string, err error) {
	query, args, err := prepeareCreate(deviceType)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

func (r *repository) Update(ctx context.Context, deviceType domain.DeviceType) (err error) {
	query, args, err := prepareUpdate(deviceType)
	if err != nil {
		return err
	}
	rows, err := r.Db.Exec(ctx, query, args...)
	if rows.RowsAffected() == 0 {
		err = errors.New(domain.ErrNoFiledsUpdated)
	}
	return
}

func (r *repository) Delete(ctx context.Context, id string) (err error) {
	query, args, err := prepareDelete(id)
	if err != nil {
		return err
	}
	rows, err := r.Db.Exec(ctx, query, args)
	if rows.RowsAffected() == 0 {
		err = errors.New(domain.ErrNoFiledsDeleted)
	}
	return
}

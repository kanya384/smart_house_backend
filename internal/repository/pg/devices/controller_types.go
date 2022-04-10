package devices

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

func NewRepository(Db *pgxpool.Pool) repo.Devices {
	return &repository{
		Db: Db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (device domain.Device, err error) {
	query, args, err := prepareGet(id)
	if err != nil {
		return
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&device.ID, &device.DeviceTypeId, &device.HousePartId)
	return
}

func (r *repository) Create(ctx context.Context, device domain.Device) (id string, err error) {
	query, args, err := prepeareCreate(device)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

func (r *repository) Update(ctx context.Context, device domain.Device) (err error) {
	query, args, err := prepareUpdate(device)
	if err != nil {
		return err
	}
	rows, err := r.Db.Exec(ctx, query, args...)
	if rows.RowsAffected() == 0 {
		err = errors.New(domain.NoFiledsUpdated)
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
		err = errors.New(domain.NoFiledsDeleted)
	}
	return
}

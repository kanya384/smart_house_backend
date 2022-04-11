package devices

import (
	"context"
	"errors"
	"smart_house_backend/internal/domain"
	repo "smart_house_backend/internal/repository"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	Db *pgxpool.Pool
	qb *queryBuilder
}

func NewRepository(Db *pgxpool.Pool, prefix string) repo.Devices {
	return &repository{
		Db: Db,
		qb: NewQueryBuilder(prefix),
	}
}

func (r *repository) Get(ctx context.Context, id string) (device domain.Device, err error) {
	query, args, err := r.qb.prepareGet(id)
	if err != nil {
		return
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&device.ID, &device.DeviceTypeId, &device.HousePartId)
	if err != nil && err.Error() == "no rows in result set" {
		return domain.Device{}, errors.New(domain.ErrNotFounded)
	}
	return
}

func (r *repository) Create(ctx context.Context, device domain.Device) (id string, err error) {
	query, args, err := r.qb.prepeareCreate(device)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		switch true {
		case strings.Contains(err.Error(), "devices_constraint"):
			err = errors.New(domain.ErrNoSuchDeviceType)
		case strings.Contains(err.Error(), "house_part_constraint"):
			err = errors.New(domain.ErrNoSuchHousePart)
		case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
			err = errors.New(domain.ErrDuplicateKey)
		default:
		}
	}
	return id, err
}

func (r *repository) Update(ctx context.Context, device domain.Device) (err error) {
	query, args, err := r.qb.prepareUpdate(device)
	if err != nil {
		return err
	}
	_, err = r.Db.Exec(ctx, query, args...)
	if err != nil {
		switch true {
		case strings.Contains(err.Error(), "devices_constraint"):
			err = errors.New(domain.ErrNoSuchDeviceType)
		case strings.Contains(err.Error(), "house_part_constraint"):
			err = errors.New(domain.ErrNoSuchHousePart)
		case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
			err = errors.New(domain.ErrDuplicateKey)
		default:
		}
	}
	return
}

func (r *repository) Delete(ctx context.Context, id string) (err error) {
	query, args, err := r.qb.prepareDelete(id)
	if err != nil {
		return err
	}
	rows, err := r.Db.Exec(ctx, query, args[0])
	if rows.RowsAffected() == 0 {
		return errors.New(domain.ErrNoFiledsDeleted)
	}
	return
}

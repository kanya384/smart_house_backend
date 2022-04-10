package house_part

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

func NewRepository(Db *pgxpool.Pool) repo.Houses {
	return &repository{
		Db: Db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (house domain.House, err error) {
	query, args, err := prepareGet(id)
	if err != nil {
		return house, err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&house.ID, &house.Name, &house.HouseID)
	return
}

func (r *repository) Create(ctx context.Context, house domain.House) (id string, err error) {
	query, args, err := prepeareCreate(house)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

func (r *repository) Update(ctx context.Context, house domain.House) (err error) {
	query, args, err := prepareUpdate(house)
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

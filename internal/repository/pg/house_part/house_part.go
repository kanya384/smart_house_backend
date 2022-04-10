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

func NewRepository(Db *pgxpool.Pool) repo.HouseParts {
	return &repository{
		Db: Db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (housePart domain.HousePart, err error) {
	query, args, err := prepareGet(id)
	if err != nil {
		return housePart, err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&housePart.ID, &housePart.Name, &housePart.HouseID)
	return
}

func (r *repository) Create(ctx context.Context, housePartType domain.HousePart) (id string, err error) {
	query, args, err := prepeareCreate(housePartType)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

func (r *repository) Update(ctx context.Context, housePartType domain.HousePart) (err error) {
	query, args, err := prepareUpdate(housePartType)
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

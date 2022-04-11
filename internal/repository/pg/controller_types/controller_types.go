package controller_types

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

func NewRepository(Db *pgxpool.Pool, prefix string) repo.ControllerTypes {
	return &repository{
		Db: Db,
		qb: NewQueryBuilder(prefix),
	}
}

func (r *repository) Get(ctx context.Context, id string) (controllerType domain.ControllerType, err error) {
	query, args, err := r.qb.prepareGet(id)
	if err != nil {
		return controllerType, err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&controllerType.ID, &controllerType.Name, &controllerType.Photo, &controllerType.DigitalPinCnt, &controllerType.AnalogPinCnt)
	if err != nil && err.Error() == "no rows in result set" {
		return domain.ControllerType{}, errors.New(domain.ErrNotFounded)
	}
	return
}

func (r *repository) Create(ctx context.Context, controllerType domain.ControllerType) (id string, err error) {
	query, args, err := r.qb.prepeareCreate(controllerType)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		switch true {
		case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
			err = errors.New(domain.ErrDuplicateKey)
		default:
		}
	}
	return id, err
}

func (r *repository) Update(ctx context.Context, controllerType domain.ControllerType) (err error) {
	query, args, err := r.qb.prepareUpdate(controllerType)
	if err != nil {
		return err
	}
	rows, err := r.Db.Exec(ctx, query, args...)
	if err != nil && rows.RowsAffected() == 0 {
		err = errors.New(domain.ErrNoFiledsUpdated)
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
		err = errors.New(domain.ErrNoFiledsDeleted)
	}
	return
}

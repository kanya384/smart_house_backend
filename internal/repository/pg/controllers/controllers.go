package controllers

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

func NewRepository(Db *pgxpool.Pool) repo.ControllersRepository {
	return &repository{
		Db: Db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (controller domain.Controller, err error) {
	query, args, err := prepareGet(id)
	if err != nil {
		return controller, err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&controller.ID, &controller.ControllerTypeId, &controller.Ip)
	return
}

func (r *repository) Create(ctx context.Context, controller domain.Controller) (id string, err error) {
	query, args, err := prepeareCreate(controller)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

func (r *repository) Update(ctx context.Context, controller domain.Controller) (err error) {
	query, args, err := prepareUpdate(controller)
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

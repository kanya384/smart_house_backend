package users

import (
	"context"
	"errors"
	"fmt"
	"smart_house_backend/internal/domain"
	repo "smart_house_backend/internal/repository"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	Db *pgxpool.Pool
}

func NewRepository(Db *pgxpool.Pool) repo.UsersRepository {
	return &repository{
		Db: Db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (user domain.User, err error) {
	query, args, err := prepareGet(id)
	if err != nil {
		return user, err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.Surname)
	return
}

func (r *repository) Create(ctx context.Context, user domain.User) (id string, err error) {
	query, args, err := prepeareCreate(user)
	if err != nil {
		return "", err
	}
	err = r.Db.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

func (r *repository) Update(ctx context.Context, user domain.User) (err error) {
	query, args, err := prepareUpdate(user)
	if err != nil {
		return err
	}
	rows, err := r.Db.Exec(ctx, query, args...)
	fmt.Println(rows.RowsAffected())
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

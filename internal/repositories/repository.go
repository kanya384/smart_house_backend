package repository

import (
	"smart_house_backend/internal/repositories/users"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Users users.Repository
}

func Setup(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Users: users.NewRepository(pool),
	}
}

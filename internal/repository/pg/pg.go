package pg

import (
	repo "smart_house_backend/internal/repository"
	"smart_house_backend/internal/repository/pg/controller_types"
	"smart_house_backend/internal/repository/pg/controllers"
	"smart_house_backend/internal/repository/pg/house_part"
	"smart_house_backend/internal/repository/pg/users"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Setup(pool *pgxpool.Pool) *repo.Repository {
	return &repo.Repository{
		Users:           users.NewRepository(pool),
		Controllers:     controllers.NewRepository(pool),
		ControllerTypes: controller_types.NewRepository(pool),
		HouseParts:      house_part.NewRepository(pool),
	}
}

package pg

import (
	repo "smart_house_backend/internal/repository"
	"smart_house_backend/internal/repository/pg/controller_types"
	"smart_house_backend/internal/repository/pg/controllers"
	"smart_house_backend/internal/repository/pg/device_type"
	"smart_house_backend/internal/repository/pg/devices"
	"smart_house_backend/internal/repository/pg/house_part"
	"smart_house_backend/internal/repository/pg/houses"
	"smart_house_backend/internal/repository/pg/pins"
	"smart_house_backend/internal/repository/pg/users"

	"github.com/jackc/pgx/v4/pgxpool"
)

const PREFIX = "public"

func Setup(pool *pgxpool.Pool) *repo.Repository {
	return &repo.Repository{
		Users:           users.NewRepository(pool),
		Controllers:     controllers.NewRepository(pool),
		ControllerTypes: controller_types.NewRepository(pool),
		HouseParts:      house_part.NewRepository(pool),
		Houses:          houses.NewRepository(pool),
		DeviceTypes:     device_type.NewRepository(pool, PREFIX),
		Devices:         devices.NewRepository(pool, PREFIX),
		Pins:            pins.NewRepository(pool),
	}
}

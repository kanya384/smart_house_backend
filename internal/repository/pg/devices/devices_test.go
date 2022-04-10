package devices_test

import (
	"context"
	"errors"
	"testing"

	"smart_house_backend/internal/config"
	"smart_house_backend/internal/domain"
	"smart_house_backend/internal/repository"
	"smart_house_backend/internal/repository/pg/controller_types"
	"smart_house_backend/internal/repository/pg/controllers"
	"smart_house_backend/internal/repository/pg/device_type"
	"smart_house_backend/internal/repository/pg/devices"
	"smart_house_backend/internal/repository/pg/house_part"
	"smart_house_backend/internal/repository/pg/houses"
	"smart_house_backend/internal/repository/pg/pins"
	"smart_house_backend/internal/repository/pg/users"
	helpers "smart_house_backend/pkg/helpers/pg"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/suite"
)

func TestRepositoryDevicesTestSuite(t *testing.T) {
	suite.Run(t, &DevicesTestSuite{})
}

type DevicesTestSuite struct {
	suite.Suite
	pool       *pgxpool.Pool
	repository *repository.Repository
}

func (s *DevicesTestSuite) SetupSuite() {
	err := s.buildRepository()
	if err != nil {
		s.FailNow("Failed to create Postgres client: %s", err)
	}
}

func (s *DevicesTestSuite) TestGet() {
	ctx := context.Background()
	device, err := s.createNecessaryData(ctx, true)
	s.Require().NoError(err, "failed to create necessary data")

	cases := map[string]struct {
		input string
		want  domain.Device
		err   error
	}{
		"success": {
			input: device.ID,
			want:  device,
			err:   nil,
		},
		"not founded": {
			input: helpers.CreateID(),
			want:  domain.Device{},
			err:   errors.New(domain.ErrNotFounded),
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			device, err := s.repository.Devices.Get(ctx, cs.input)
			s.Equal(err, cs.err)
			s.Equal(device, cs.want)
		})
	}

}

func (s *DevicesTestSuite) TestCreate() {
	ctx := context.Background()
	device, err := s.createNecessaryData(ctx, false)
	s.Require().NoError(err, "failed to create necessary data")

	cases := map[string]struct {
		input domain.Device
		want  string
		err   error
	}{
		"create device success": {
			input: device,
			want:  device.ID,
			err:   nil,
		},
		"no device type error": {
			input: domain.Device{ID: helpers.CreateID(), DeviceTypeId: helpers.CreateID(), HousePartId: device.HousePartId},
			want:  "",
			err:   errors.New(domain.ErrNoSuchDeviceType),
		},
		"no house part error": {
			input: domain.Device{ID: helpers.CreateID(), DeviceTypeId: device.DeviceTypeId, HousePartId: helpers.CreateID()},
			want:  "",
			err:   errors.New(domain.ErrNoSuchHousePart),
		},
		"duplicate key": {
			input: device,
			want:  "",
			err:   errors.New(domain.ErrDuplicateKey),
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			id, err := s.repository.Devices.Create(ctx, cs.input)
			s.Equal(err, cs.err)
			s.Equal(id, cs.want)
		})
	}
}

func (s *DevicesTestSuite) createNecessaryData(ctx context.Context, create bool) (domain.Device, error) {

	user := domain.User{
		ID:      helpers.CreateID(),
		Name:    "Test Name",
		Surname: "Test Surname",
	}
	_, err := s.repository.Users.Create(ctx, user)
	s.Require().NoError(err, "Failed to create user")

	house := domain.House{
		ID:      helpers.CreateID(),
		Name:    "Test House",
		OwnerID: user.ID,
	}
	_, err = s.repository.Houses.Create(ctx, house)
	s.Require().NoError(err, "Failed to create house")

	house_part := domain.HousePart{
		ID:      helpers.CreateID(),
		Name:    "Test House Part",
		HouseID: house.ID,
	}
	_, err = s.repository.HouseParts.Create(ctx, house_part)
	s.Require().NoError(err, "Failed to create house part")

	device_type := domain.DeviceType{
		ID:   helpers.CreateID(),
		Name: "Test_Type",
	}
	_, err = s.repository.DeviceTypes.Create(ctx, device_type)
	s.Require().NoError(err, "Failed to create device type")

	device := domain.Device{
		ID:           helpers.CreateID(),
		DeviceTypeId: device_type.ID,
		HousePartId:  house_part.ID,
	}

	if create {
		_, err = s.repository.Devices.Create(ctx, device)
		s.Require().NoError(err, "Failed to create entity")

		entity, err := s.repository.Devices.Get(ctx, device.ID)
		s.Require().NoError(err, "It should load entity without errors")
		s.Require().Equal(device, entity, "It should return all entity data")
	}
	return device, err
}

func (s *DevicesTestSuite) buildRepository() (err error) {
	config := config.Config{
		PostgresHost:            "localhost",
		PostgresPort:            "5432",
		PostgresUsername:        "admin",
		PostgresPass:            "admin",
		PostgresDbName:          "smart_house",
		PostgresPoolConnections: 10,
	}

	pgConfig := &helpers.Config{
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		Username: config.PostgresUsername,
		Password: config.PostgresPass,
		DbName:   config.PostgresDbName,
		Timeout:  5,
	}

	poolConfig, err := helpers.NewPoolConfig(pgConfig)

	if err != nil {
		s.FailNow("error creating pg pool config", err.Error(), map[string]interface{}{})
	}
	poolConfig.MaxConns = config.PostgresPoolConnections

	pool, err := helpers.NewConnection(poolConfig)
	if err != nil {
		s.FailNow("connect to database failed", err, map[string]interface{}{})
	}

	s.pool = pool

	s.repository = &repository.Repository{
		Users:           users.NewRepository(pool),
		Controllers:     controllers.NewRepository(pool),
		ControllerTypes: controller_types.NewRepository(pool),
		HouseParts:      house_part.NewRepository(pool),
		Houses:          houses.NewRepository(pool),
		DeviceTypes:     device_type.NewRepository(pool),
		Devices:         devices.NewRepository(pool, repository.DEVICES_TABLE+"_test"),
		Pins:            pins.NewRepository(pool),
	}

	return
}

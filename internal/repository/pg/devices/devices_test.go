package devices_test

import (
	"context"
	"errors"
	"testing"

	"smart_house_backend/internal/config"
	"smart_house_backend/internal/domain"
	"smart_house_backend/internal/repository"
	"smart_house_backend/internal/repository/pg/devices"
	helpers "smart_house_backend/pkg/helpers/pg"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

const SCHEMA_NAME = "TEST"

func TestRepositoryDevicesTestSuite(t *testing.T) {
	suite.Run(t, &DevicesTestSuite{})
}

type DevicesTestSuite struct {
	suite.Suite
	pool       *pgxpool.Pool
	repository repository.Devices
}

func (s *DevicesTestSuite) SetupSuite() {
	err := s.buildRepository()
	if err != nil {
		s.FailNow("Failed to create Postgres client: %s", err)
	}
}

func (s *DevicesTestSuite) TestGet() {
	ctx := context.Background()
	device := domain.Device{ID: "70d3d531-4041-4d74-8306-bf8e7319b74b", DeviceTypeId: "4fba07cb-7c5e-4a18-a62f-2e9044a50c1b", HousePartId: "8120f91d-17b9-405a-ae74-797c4c9e0117"}

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
			device, err := s.repository.Get(ctx, cs.input)
			s.Equal(err, cs.err)
			s.Equal(device, cs.want)
		})
	}

}

func (s *DevicesTestSuite) TestCreate() {
	ctx := context.Background()
	device := domain.Device{ID: helpers.CreateID(), DeviceTypeId: "4fba07cb-7c5e-4a18-a62f-2e9044a50c1b", HousePartId: "8120f91d-17b9-405a-ae74-797c4c9e0117"}

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
			id, err := s.repository.Create(ctx, cs.input)
			s.Equal(err, cs.err)
			s.Equal(id, cs.want)
		})
	}
}

func (s *DevicesTestSuite) buildRepository() (err error) {
	config, err := config.InitConfig("APP")
	if err != nil {
		logrus.Panic("error initializing config: %w", err)
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

	s.repository = devices.NewRepository(pool, SCHEMA_NAME)

	return
}

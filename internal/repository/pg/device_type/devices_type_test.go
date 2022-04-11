package device_type_test

import (
	"context"
	"errors"
	"testing"

	"smart_house_backend/internal/config"
	"smart_house_backend/internal/domain"
	"smart_house_backend/internal/repository"
	"smart_house_backend/internal/repository/pg/device_type"
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
	repository repository.DeviceTypes
}

func (s *DevicesTestSuite) SetupSuite() {
	err := s.buildRepository()
	if err != nil {
		s.FailNow("Failed to create Postgres client: %s", err)
	}
}

func (s *DevicesTestSuite) TestGet() {
	ctx := context.Background()
	deviceType := domain.DeviceType{ID: "4fba07cb-7c5e-4a18-a62f-2e9044a50c1b", Name: "Выключатель", Photo: "https://avselectro.ru/uploads/gallery/44/max/1a2a3ef554d7cd9f32fc6895a6f13d86.jpg"}
	cases := map[string]struct {
		input string
		want  domain.DeviceType
		err   error
	}{
		"success": {
			input: deviceType.ID,
			want:  deviceType,
			err:   nil,
		},
		"not founded": {
			input: helpers.CreateID(),
			want:  domain.DeviceType{},
			err:   errors.New(domain.ErrNotFounded),
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			device, err := s.repository.Get(ctx, cs.input)
			s.Equal(cs.err, err)
			s.Equal(cs.want, device)
		})
	}

}

func (s *DevicesTestSuite) TestCreate() {
	ctx := context.Background()
	deviceType := domain.DeviceType{ID: helpers.CreateID(), Name: "Выключатель", Photo: "https://avselectro.ru/uploads/gallery/44/max/1a2a3ef554d7cd9f32fc6895a6f13d86.jpg"}

	cases := map[string]struct {
		input domain.DeviceType
		want  string
		err   error
	}{
		"create device success": {
			input: deviceType,
			want:  deviceType.ID,
			err:   nil,
		},
		"duplicate key": {
			input: deviceType,
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

func (s *DevicesTestSuite) TestUpdate() {
	ctx := context.Background()
	deviceType := domain.DeviceType{ID: "4fba07cb-7c5e-4a18-a62f-2e9044a50c1b", Name: "Выключатель", Photo: "https://avselectro.ru/uploads/gallery/44/max/1a2a3ef554d7cd9f32fc6895a6f13d86.jpg"}

	cases := map[string]struct {
		input domain.DeviceType
		err   error
	}{
		"update device success": {
			input: deviceType,
			err:   nil,
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			err := s.repository.Update(ctx, cs.input)
			s.Equal(cs.err, err)
		})
	}
}

func (s *DevicesTestSuite) TestDelete() {
	ctx := context.Background()
	deviceTypeID := "2e9044a50c1b-7c5e-4a18-a62f-4fba07cb"

	cases := map[string]struct {
		input string
		err   error
	}{
		"delete device type success": {
			input: deviceTypeID,
			err:   nil,
		},
		"no device to delete": {
			input: helpers.CreateID(),
			err:   errors.New(domain.ErrNoFiledsDeleted),
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			err := s.repository.Delete(ctx, cs.input)
			s.Equal(err, cs.err)
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

	s.repository = device_type.NewRepository(pool, SCHEMA_NAME)

	return
}

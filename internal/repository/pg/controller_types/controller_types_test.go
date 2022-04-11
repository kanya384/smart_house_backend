package controller_types_test

import (
	"context"
	"errors"
	"testing"

	"smart_house_backend/internal/config"
	"smart_house_backend/internal/domain"
	"smart_house_backend/internal/repository"
	"smart_house_backend/internal/repository/pg/controller_types"
	helpers "smart_house_backend/pkg/helpers/pg"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

const SCHEMA_NAME = "TEST"

func TestRepositoryControllerTypes(t *testing.T) {
	suite.Run(t, &ControllerTypes{})
}

type ControllerTypes struct {
	suite.Suite
	pool       *pgxpool.Pool
	repository repository.ControllerTypes
}

func (s *ControllerTypes) SetupSuite() {
	err := s.buildRepository()
	if err != nil {
		s.FailNow("Failed to create Postgres client: %s", err)
	}
}

func (s *ControllerTypes) TestGet() {
	ctx := context.Background()
	controllerType := domain.ControllerType{ID: "39248a56-18d7-46c1-bbd9-a8139b6bf1fa", Name: "Orange Pi One", Photo: "https://static.chipdip.ru/lib/736/DOC002736925.jpg", DigitalPinCnt: 11, AnalogPinCnt: 5}
	cases := map[string]struct {
		input string
		want  domain.ControllerType
		err   error
	}{
		"success": {
			input: controllerType.ID,
			want:  controllerType,
			err:   nil,
		},
		"not founded": {
			input: helpers.CreateID(),
			want:  domain.ControllerType{},
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

func (s *ControllerTypes) TestCreate() {
	ctx := context.Background()
	controllerType := domain.ControllerType{ID: "39248a56-18d7-46c1-bbd9-a8139b6bf1fa", Name: "Orange Pi One", Photo: "https://static.chipdip.ru/lib/736/DOC002736925.jpg", DigitalPinCnt: 11, AnalogPinCnt: 5}
	newID := helpers.CreateID()
	cases := map[string]struct {
		input domain.ControllerType
		want  string
		err   error
	}{
		"create device success": {
			input: domain.ControllerType{ID: newID, Name: "Name"},
			want:  newID,
			err:   nil,
		},
		"duplicate key": {
			input: controllerType,
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

func (s *ControllerTypes) TestUpdate() {
	ctx := context.Background()
	controllerType := domain.ControllerType{ID: "39248a56-18d7-46c1-bbd9-a8139b6bf1fa", Name: "Orange Pi One2", Photo: "https://static.chipdip.ru/lib/736/DOC002736925.jpg", DigitalPinCnt: 11, AnalogPinCnt: 5}

	cases := map[string]struct {
		input domain.ControllerType
		err   error
	}{
		"update success": {
			input: controllerType,
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

func (s *ControllerTypes) TestDelete() {
	ctx := context.Background()
	controllerTypeID := "a8139b6bf1fa-18d7-46c1-bbd9-39248a56"

	cases := map[string]struct {
		input string
		err   error
	}{
		"delete device type success": {
			input: controllerTypeID,
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

func (s *ControllerTypes) buildRepository() (err error) {
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

	s.repository = controller_types.NewRepository(pool, SCHEMA_NAME)

	return
}

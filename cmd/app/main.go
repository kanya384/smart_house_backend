package main

import (
	"context"
	"smart_house_backend/internal/config"
	"smart_house_backend/internal/domain"
	repository "smart_house_backend/internal/repositories"
	"smart_house_backend/pkg/helpers/pg"
	"smart_house_backend/pkg/logger"

	"github.com/sirupsen/logrus"
)

const LOG_FILE = "var/logs/main.log"
const SERVICE_NAME = "smart_house_backend"

func main() {
	cfg, err := config.InitConfig("APP")
	if err != nil {
		logrus.Panic("error initializing config: %w", err)
	}
	logger, err := logger.NewLogger("leadgen_backend", cfg.LogLevel, LOG_FILE)
	if err != nil {
		logrus.Panic("error initializing logger: %w", err)
	}

	//postgress
	pgConfig := &pg.Config{
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		Username: cfg.PostgresUsername,
		Password: cfg.PostgresPass,
		DbName:   cfg.PostgresDbName,
		Timeout:  5,
	}

	poolConfig, err := pg.NewPoolConfig(pgConfig)
	if err != nil {
		logger.Panic("error creating pg pool config", err, map[string]interface{}{})
	}
	poolConfig.MaxConns = cfg.PostgresPoolConnections

	pool, err := pg.NewConnection(poolConfig)
	if err != nil {
		logger.Panic("connect to database failed", err, map[string]interface{}{})
	}
	repository := repository.Setup(pool)
	//id, err := repository.Users.CreateUser(context.Background(), domain.User{ID: pg.CreateID(), Name: "Мулиат", Surname: "Кушу"})
	//user, err := repository.Users.GetUser(context.Background(), "670f40d2-3c3e-4ce3-9c78-9d924bbd5e89")
	err = repository.Users.UpdateUser(context.Background(), domain.User{ID: "a7364b9e-2989-4da6-b22c-6d48b77bf4dc_", Name: "Рамазан_", Surname: "Кушу"})
}

package main

import (
	"fmt"
	"smart_house_backend/internal/config"
	"smart_house_backend/pkg/helpers/pg"
	"smart_house_backend/pkg/logger"

	"github.com/sirupsen/logrus"
)

const LOG_FILE = "/var/logs/main.log"
const SERVICE_NAME = "smart_house_backend"

func main() {
	cfg, err := config.InitConfig("APP")
	if err != nil {
		logrus.Panic("error initializing config: %w", err)
	}
	logger, err := logger.NewLogger("leadgen_backend", cfg.LogLevel, LOG_FILE)
	if err != nil {
		logger.Panic(fmt.Sprintf("error initi"), err, nil)
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

	connPg, err := pg.NewConnection(poolConfig)
	if err != nil {
		logger.Panic("connect to database failed", err, map[string]interface{}{})
	}
}

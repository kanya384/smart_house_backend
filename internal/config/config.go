package config

import (
	"time"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port                    string
	LogLevel                uint32
	Salt                    string
	TokenTTL                time.Duration
	RefreshTokenTTL         time.Duration
	TokenSecret             string
	PostgresHost            string
	PostgresPort            string
	PostgresUsername        string
	PostgresPass            string
	PostgresDbName          string
	PostgresPoolConnections int32
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return conf, err
	}
	return conf, nil
}

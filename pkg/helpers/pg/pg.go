package pg

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

//struct includes all required pg params
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Timeout  int
}

//NewPoolConfig returns connection string
func NewPoolConfig(cfg *Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Timeout)
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	return poolConfig, nil
}

//NewConnection creates connection with pool
func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

//convert bool to bit
func ConvertBoolToBit(boolValue bool) string {
	if boolValue {
		return "1"
	}
	return "0"
}

func ConvertBitToBool(bit pgtype.Bit) bool {
	bitVal, _ := bit.Value()
	return bitVal == "1"
}

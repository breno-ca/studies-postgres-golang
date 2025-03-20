package config

import (
	"fmt"
)

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (c *postgresConfig) DSN() string {
	dsn := fmt.Sprintf(
		// postgres://user:password@host:port/dbname?sslmode=disable
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	return dsn
}

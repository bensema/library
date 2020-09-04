package sql

import _ "github.com/go-sql-driver/mysql"

type Config struct {
	DSN         string
	ReadDSN     []string
	Active      int
	Idle        int
	IdleTimeout int
}

func New(c *Config) (db *DB) {
	db, err := Open(c)
	if err != nil {
	}
	return
}

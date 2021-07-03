package db

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type linkDBKey string

var (
	noDBError = errors.New("context does not contain DB")

	key = linkDBKey("linkdb")
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func (c *Config) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
}

func NewDB(config Config) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", config.String())
}

func WithDB(ctx context.Context, db *sqlx.DB) context.Context {
	return context.WithValue(ctx, key, db)
}

func DB(ctx context.Context) (*sqlx.DB, error) {
	if db, ok := ctx.Value(key).(*sqlx.DB); ok {
		return db, nil
	}

	return nil, noDBError
}

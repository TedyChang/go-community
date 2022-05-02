package model_test

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go-community/model"
	"testing"
)

func setup(f func()) func(t *testing.T) {

	dsn := "postgres://postgres:postgres@localhost:3999/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	model.Db = bun.NewDB(sqldb, pgdialect.New())

	return func(t *testing.T) {
		f()
	}
}

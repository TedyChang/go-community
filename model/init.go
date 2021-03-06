package model

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var Db *bun.DB

func RunDb() {
	dsn := "postgres://postgres:postgres@localhost:3999/postgres"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	Db = bun.NewDB(sqldb, pgdialect.New())

	Db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.WithEnabled(true),
	))
}

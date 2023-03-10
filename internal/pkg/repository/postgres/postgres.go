package postgres

import (
	"context"
	"database/sql"
	"github.com/Abdullayev65/online_test/internal/entity"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var (
	debugging  = true
	resetModel = false
)

func New() *bun.DB {
	dsn := "postgres:root123//postgres:localhost:5432/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(pgdb, pgdialect.New())

	if debugging {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	ctx := context.Background()

	ms := []interface{}{
		(*entity.User)(nil),
		(*entity.Topic)(nil),
		(*entity.Question)(nil),
		(*entity.Answer)(nil),
	}
	if resetModel {
		db.ResetModel(ctx, ms...)
	} else {
		db.RegisterModel(ms...)
	}

	return db
}

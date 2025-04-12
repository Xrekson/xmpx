package cmd

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/xrekson/auction/pkg/model"
)

var DB *bun.DB

func InitDB() error {
	// Connect to PostgreSQL
	sqldb := pgdriver.NewConnector(
		pgdriver.WithUser("app"),
		pgdriver.WithPassword("Thundera@190"),
		pgdriver.WithDatabase("app"),
		pgdriver.WithAddr("localhost:5432"),
	)

	db := bun.NewDB(sql.OpenDB(sqldb), pgdialect.New())

	// Add query hook
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	// Context set
	ctx := context.Background()

	// Create tables
	_, err := db.NewCreateTable().Model((*model.User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return err
	}

	_, err = db.NewCreateTable().Model((*model.Listing)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

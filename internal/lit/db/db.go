package db

import (
	"context"

	"github.com/m-nny/go-lit/internal/lit"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	gDb    *gorm.DB
	ctx    context.Context // background context
	cancel func()          // cancel background context

	// Datasource name.
	DSN string

	// // Destination for events to be published.
	// EventService lit.EventService

}

func NewDB(dsn string) *DB {
	ctx, cancel := context.WithCancel(context.Background())
	return &DB{
		ctx:    ctx,
		cancel: cancel,

		DSN: dsn,
	}
}

func (db *DB) Open() (err error) {
	if db.DSN == "" {
		return lit.Errorf(lit.EINVALID, "dsn required")
	}
	if db.gDb, err = gorm.Open(sqlite.Open(db.DSN), &gorm.Config{}); err != nil {
		return err
	}

	db.gDb.AutoMigrate(&UserModel{})

	return nil
}

func (db *DB) Close() (err error) {
	db.cancel()

	if db.gDb != nil {
		sqlDb, err := db.gDb.DB()
		if err != nil {
			return err
		}
		return sqlDb.Close()
	}

	return nil
}

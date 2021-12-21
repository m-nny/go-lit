package gorm

import (
	"github.com/m-nny/go-lit/internal/lit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	gorm *gorm.DB
	// Datasource name.
	DSN string

	// // Destination for events to be published.
	// EventService lit.EventService

}

func NewDB(dsn string) *DB {
	return &DB{
		DSN: dsn,
	}
}

func (db *DB) Open() (err error) {
	if db.DSN == "" {
		return lit.Errorf(lit.EINVALID, "dsn required")
	}
	if db.gorm, err = gorm.Open(postgres.Open(db.DSN), &gorm.Config{}); err != nil {
		return err
	}

	db.gorm.AutoMigrate(&UserModel{})

	return nil
}

func (db *DB) Close() (err error) {

	if db.gorm != nil {
		sqlDb, err := db.gorm.DB()
		if err != nil {
			return err
		}
		return sqlDb.Close()
	}

	return nil
}

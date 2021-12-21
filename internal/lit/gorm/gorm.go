package gorm

import (
	"github.com/m-nny/go-lit/internal/lit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormService struct {
	GormDb *gorm.DB

	config *lit.DbConfig
}

func NewGormService(config *lit.DbConfig) *GormService {
	return &GormService{
		config: config,
	}
}

func (d *GormService) Open() (err error) {
	dsn := d.config.DSN()
	if dsn == "" {
		return lit.Errorf(lit.EINVALID, "dsn required")
	}
	if d.GormDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}

	err = d.GormDb.AutoMigrate(&UserModel{})
	if err != nil {
		return err
	}

	return nil
}

func (db *GormService) Close() (err error) {

	if db.GormDb != nil {
		sqlDb, err := db.GormDb.DB()
		if err != nil {
			return err
		}
		return sqlDb.Close()
	}

	return nil
}

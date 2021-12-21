package main

import (
	"context"
	"log"

	"github.com/m-nny/go-lit/internal/lit"
	"github.com/m-nny/go-lit/internal/lit/db"
)

type App struct {
	Config *lit.AppConfig
	Db     *db.DB
}

func NewApp() *App {
	return &App{}
}
func (app *App) Load() (err error) {
	if app.Config, err = lit.LoadAppConfig(); err != nil {
		return err
	}
	app.Db = db.NewDB(app.Config.Db.DSN)
	return nil
}

func (app *App) Run(ctx context.Context) (err error) {
	if app.Config == nil {
		return lit.Errorf(lit.EINTERNAL, "app.Config is nil")
	}
	if app.Db == nil {
		return lit.Errorf(lit.EINTERNAL, "app.Db is nil")
	}
	if err = app.Db.Open(); err != nil {
		return err
	}
	log.Printf("App started")
	return nil
}

func (app *App) Close() (err error) {
	// if app.HTTPServer != nil {
	// 	if err := app.HTTPServer.Close(); err != nil {
	// 		return err
	// 	}
	// }
	if app.Db != nil {
		if err := app.Db.Close(); err != nil {
			return err
		}
	}
	return nil
}

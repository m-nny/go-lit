package lit

import (
	"context"
	"log"
)

type App struct {
	Config *AppConfig
	// Db * Db
}

func NewApp() *App {
	config, err := LoadAppConfig()
	if err != nil {
		panic(err)
	}
	return &App{Config: config}
}

func (app *App) Run(ctx context.Context) (err error) {
	log.Printf("App started")
	return nil
}

func (app *App) Close() (err error) {
	// if app.HTTPServer != nil {
	// 	if err := app.HTTPServer.Close(); err != nil {
	// 		return err
	// 	}
	// }
	// if app.DB != nil {
	// 	if err := app.DB.Close(); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

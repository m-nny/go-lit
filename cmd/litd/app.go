package main

import (
	"log"

	"github.com/m-nny/go-lit/internal/lit"
	"github.com/m-nny/go-lit/internal/lit/echo"
	"github.com/m-nny/go-lit/internal/lit/gorm"
)

type App struct {
	config      *lit.AppConfig
	gormService *gorm.GormService
	httpServer  *echo.HttpServer

	userService lit.UserService
}

func NewApp() *App {
	return &App{}
}
func (app *App) Run() (err error) {
	if app.config, err = lit.LoadAppConfig(); err != nil {
		return err
	}
	app.gormService = gorm.NewGormService(&app.config.Db)
	if err = app.gormService.Open(); err != nil {
		return err
	}

	app.userService = gorm.NewUserService(app.gormService.GormDb)

	app.httpServer = echo.NewHttp(&app.config.Http, &app.userService)

	if err = app.httpServer.Open(); err != nil {
		return err
	}
	log.Printf("App started")
	return nil
}

func (app *App) Close() (err error) {
	if app.httpServer != nil {
		if err := app.httpServer.Close(); err != nil {
			return err
		}
	}
	if app.gormService != nil {
		if err := app.gormService.Close(); err != nil {
			return err
		}
	}
	return nil
}

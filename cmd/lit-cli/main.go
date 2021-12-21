package main

import (
	"context"
	"fmt"
	"os"

	"github.com/m-nny/go-lit/internal/lit"
	g "github.com/m-nny/go-lit/internal/lit/gorm"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	config, err := lit.LoadAppConfig()
	if err != nil {
		return err
	}

	gormService := g.NewGormService(&config.Db)

	err = gormService.Open()
	if err != nil {
		return err
	}

	var userModels []g.UserModel
	gormService.GormDb.Find(&userModels)
	fmt.Printf("%+v\n", userModels)

	userService := g.NewUserService(gormService.GormDb)

	users, _, err := userService.FindUsers(context.Background(), &lit.UserFilter{})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", users)

	return nil
}

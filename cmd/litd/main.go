package main

import (
	"context"
	"fmt"
	"os"

	"github.com/m-nny/go-lit/internal/lit"
)

// main is the entry point into our application. However, it provides poor
// usability since it does not allow us to return errors like most Go programs.
// Instead, we delegate most of our program to the Run() function.
func main() {
	// Setup signal handlers.
	// ctx, cancel := context.WithCancel(context.Background())
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// go func() { <-c; cancel() }()

	// Instantiate a new type to represent our application.
	// This type lets us shared setup code with our end-to-end tests.
	app := NewApp()

	if err := app.Run(); err != nil {
		app.Close()
		fmt.Fprintln(os.Stderr, err)
		lit.ReportError(context.Background(), err)
		os.Exit(1)
	}

	if err := app.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

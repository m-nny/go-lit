package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"lit"
	"os"
	"os/signal"
)

const (
	DefaultConfigPath = "~/lit.conf"
	DefaultURL        = "https://lit.org"
)

// main is the entry point into our application. However, it provides poor
// usability since it does not allow us to return errors like most Go programs.
// Instead, we delegate most of our program to the Run() function.
func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	// Execute program.
	//
	// If an ErrHelp error is returned then that means the user has used an "-h"
	// flag and the flag package will handle output. We just need exit.
	//
	// If we have an application error (lit.Error) then we can just display the
	// message. If we have any other error, print the raw error message.
	var e *lit.Error
	if err := Run(ctx, os.Args[1:]); err == flag.ErrHelp {
		os.Exit(1)
	} else if errors.As(err, &e) {
		fmt.Fprintf(os.Stderr, e.Message)
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run(ctx context.Context, args []string) error {
	var cmd string
	if len(args) > 0 {
		cmd, args = args[0], args[1:]
	}
	switch cmd {
	case "dial":
		return (&DialCommand{}).Run(ctx, args)
	case "", "-h", "help":
		usage()
		return flag.ErrHelp
	default:
		return fmt.Errorf("lit %s: unknown command", cmd)
	}
}

// usage prints the top-level CLI usage message.
func usage() {
	fmt.Println(`
Command line utility for interacting with the lit Dial service.

Usage:

	lit <command> [arguments]

The commands are:

	dial        manage your dial
`[1:])
}

type Config struct {
	URL string `toml:"url"`
}

package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"pkg.razonyang.com/go-auth0-web-app/internal/core"
)

var (
	cfg string
	app = &cli.App{
		EnableBashCompletion: true,
		Version:              core.Version,
		Name:                 "go-auth0-web-app",
		Usage:                "Go Auth0 Web Application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       ".env",
				Destination: &cfg,
			},
		},
		Before: func(c *cli.Context) error {
			if err := godotenv.Load(cfg); err != nil {
				return err
			}

			return nil
		},
	}
)

// Execute executes commands.
func Execute() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

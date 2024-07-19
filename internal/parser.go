package internal

//Provides the features for parsing command line arguments.

import (
	"github.com/urfave/cli/v2"
)

// Parse the command line arguments.
func Parse(args []string) error {
	// https://cli.urfave.org/
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "awsprofile",
				Aliases: []string{"p"},
				Usage:   "The AWS profile to get credentials.",
				Action: func(*cli.Context, string) error {

					return nil
				},
			},
		},
	}
	if err := app.Run(args); err != nil {
		return err
	}
	return nil

}

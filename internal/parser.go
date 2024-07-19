package internal

//Provides the features for parsing command line arguments.

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Parse(args []string) error {
	// https://cli.urfave.org/
	app := &cli.App{
		Name:                 "dpreport",
		EnableBashCompletion: true,
		ArgsUsage:            "doge",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "awsprofile",
				Aliases: []string{"p"},
				Usage:   "The AWS profile to get credentials.",
			},
		},
		Action: func(ctx *cli.Context) error {
			aws := ctx.String("awsprofile")
			ctx.Args()
			fmt.Print("Hi\n")
			fmt.Printf("%v\n %v\n", ctx, aws)
			fmt.Printf("%v\n", aws)
			return nil
		},
	}
	if err := app.Run(args); err != nil {
		return err
	}
	return nil

}

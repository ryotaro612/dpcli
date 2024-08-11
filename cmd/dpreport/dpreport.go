// Package main
package main

import (
	_ "context"
	"fmt"
	_ "log"
	"os"

	"github.com/urfave/cli/v2"

	// _ "github.com/aws/aws-sdk-go-v2/aws"
	// _ "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	// _ "github.com/aws/aws-sdk-go/aws/awserr"
	// _ "github.com/aws/aws-sdk-go/aws/session"
	"github.com/ryotaro612/dpcli/internal"
	// _ "github.com/slack-go/slack"
	// _ "github.com/urfave/cli/v2" // imports as package "cli"
	// _ "google.golang.org/api/calendar/v3"

	_ "log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	err := report(os.Args)
	if err != nil {
		internal.MakeLogger(false).Error("Failure", "error", err)
	}
}

func report(args []string) error {
	awsProfile := "awsProfile"
	template := "template"
	verbose := "verbose"
	// https://cli.urfave.org/
	app := &cli.App{
		Name:                 "dpreport",
		EnableBashCompletion: true,
		// make the help subcommand disable.
		HideHelpCommand: true,
		//ArgsUsage:            "doge",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    awsProfile,
				Aliases: []string{"p"},
				Usage:   "The AWS profile to get credentials.",
			},
			&cli.StringFlag{
				Name:    template,
				Aliases: []string{"t"},
				Usage:   "A Go template file. Annotations in the template refer to activities like meeting and pull requests.",
			},
			&cli.BoolFlag{
				Name:    verbose,
				Aliases: []string{"v"},
				Usage:   "Prints more information.",
			},
		},
		Action: func(ctx *cli.Context) error {
			aws := ctx.String(awsProfile)
			template := ctx.String(template)
			verbose := ctx.Bool(verbose)
			reporting, err := internal.MakeReporting(ctx.Context, aws, verbose, template)
			if err != nil {
				return err
			}
			report, err := reporting.Report(ctx.Context)
			if err != nil {
				return err
			}
			fmt.Println(report)
			return nil
		},
	}
	if err := app.Run(args); err != nil {
		return err
	}
	return nil

}

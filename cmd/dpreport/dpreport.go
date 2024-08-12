// Package main
package main

import (
	_ "context"
	_ "log"
	"os"
	"time"

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

	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	err := report(os.Args)
	if err != nil {
		internal.NewLogger(false).Error("Failure", "error", err)
	}
}

func report(args []string) error {
	awsProfile := "awsProfile"
	template := "template"
	verbose := "verbose"
	pullRequestsAfter := "prafter"
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
			&cli.StringFlag{
				Name:    pullRequestsAfter,
				Aliases: []string{"r"},
				Usage:   "Read pull requests created after the specified datetime. The format is like 2006-01-02 15:04:05",
			},
		},
		Action: func(ctx *cli.Context) error {
			aws := ctx.String(awsProfile)
			template := ctx.String(template)
			verbose := ctx.Bool(verbose)
			after := ctx.String(pullRequestsAfter)
			reporting, err := internal.NewReporting(ctx.Context, aws, verbose, template)
			if err != nil {
				return err
			}
			if after != "" {
				//time.Parse(DateTime   = "2006-01-02 15:04:05")
				parsed, err := time.Parse(time.DateTime, after)
				if err != nil {
					return err
				}
				return reporting.ReportOffset(ctx.Context, parsed)
			}
			return reporting.Report(ctx.Context)
		},
	}
	if err := app.Run(args); err != nil {
		return err
	}
	return nil

}

package internal

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ryotaro612/dpcli/internal/calendar"
	"github.com/ryotaro612/dpcli/internal/github"
)

type deps struct {
	github.Client
}

type Option struct {
	awsProfile   *string
	templateFile *string
	verbose      bool
}

type writer struct {
}

func (w writer) Writer() {

}

type template struct {
}

type generator struct {
}

func (g generator) generate(
	t template, events []calendar.Event, pullRequests []github.PullRequest) (string, error) {
	return "", nil
}

type Reporting struct {
	github    github.Client
	calendar  calendar.Client
	template  template
	generator generator
	writer    writer
	logger    *slog.Logger
}

func (r Reporting) Report(ctx context.Context) (string, error) {
	pullRequests, err := r.github.ReadPullRequests()
	if err != nil {
		return "", err
	}
	fmt.Println(pullRequests)
	events, err := r.calendar.ReadEvents()
	if err != nil {
		return "", err
	}
	report, err := r.generator.generate(r.template, events, pullRequests)

	return report, err
}

// MakeReporting creates a new Reporting object with the given options.
func MakeReporting(ctx context.Context, awsProfile string, verbose bool, template string) (Reporting, error) {
	var logOptions *slog.HandlerOptions
	if verbose {
		logOptions = &slog.HandlerOptions{Level: slog.LevelDebug}
	}
	logger := MakeLogger(logOptions)
	secretClient, err := NewSecretClient(ctx, logger, awsProfile)
	if err != nil {
		return Reporting{}, err
	}
	secret, err := secretClient.Secret(ctx)
	if err != nil {
		return Reporting{}, err
	}
	//secretClient.GetSecretValue(ctx, input)
	fmt.Printf("%v", secret.GithubToken)
	return Reporting{logger: logger}, nil
}

package internal

import (
	"context"
	"fmt"
	"github.com/ryotaro612/dpcli/internal/calendar"
	"github.com/ryotaro612/dpcli/internal/github"
	"log/slog"
	"time"
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
	github    *github.Client
	calendar  calendar.Client
	template  template
	generator generator
	writer    writer
	logger    *slog.Logger
}

func (r Reporting) Report(ctx context.Context) (string, error) {
	offset, err := calcOffset(time.Now())
	if err != nil {
		return "", err
	}
	pullRequests, err := r.github.ReadPullRequests(ctx, offset)
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

// NewReporting creates a new Reporting object with the given options.
func NewReporting(ctx context.Context, awsProfile string, verbose bool, template string) (Reporting, error) {
	logger := NewLogger(verbose)
	var r Reporting
	s, err := ReadSecret(ctx, logger, awsProfile)
	if err != nil {
		return r, err
	}
	g := github.NewClient(logger, s.GithubToken)
	fmt.Printf("foobar %v dge", s.GithubToken)
	return Reporting{github: &g, logger: logger}, nil
}

func calcOffset(current time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return current, err
	}
	year, month, day := current.In(loc).Date()
	if time.Date(year, month, day, 11, 0, 0, 0, loc).After(current) {
		return time.Date(year, month, day, 11, 0, 0, 0, loc).AddDate(0, 0, -1), nil

	}
	return time.Date(year, month, day, 11, 0, 0, 0, loc), nil

}

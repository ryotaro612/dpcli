package internal

import (
	"context"
	"fmt"
	"github.com/ryotaro612/dpcli/internal/calendar"
	"github.com/ryotaro612/dpcli/internal/github"
)

type deps struct {
	github.Client
}

type Option struct {
	awsProfile   *string
	templateFile *string
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

func MakeReporting(o Option) Reporting {
	return Reporting{}
}

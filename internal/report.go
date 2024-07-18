package internal

import (
	"context"
	"github.com/ryotaro612/dpcli/internal/calendar"
	"github.com/ryotaro612/dpcli/internal/github"
)

type deps struct {
	github.Client
}

type Option struct {
	awsProfile   *string
	templateFile *string
	output       *string
}

type Reporting struct {
	github   github.Client
	calendar calendar.Client
}

func Report(ctx context.Context, deps deps) error {
	pullRequests, err := deps.ReadPullRequests()
	if err != nil {
		return err
	}
	calendar.ReadEvents()
}

func readWork() {

}

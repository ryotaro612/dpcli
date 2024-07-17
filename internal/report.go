package internal

import (
	"context"
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

func PostReport(ctx context.Context, deps deps) error {
	pullRequests, err := deps.ReadPullRequests()
	if err != nil {
		return err
	}

}

func readWork() {

}

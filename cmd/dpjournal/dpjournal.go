// Package doge
package main

import (
	"context"
	"fmt"
	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/awserr"
	_ "github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/ryotaro612/dpcli/internal"
	_ "github.com/slack-go/slack"
	_ "github.com/urfave/cli/v2" // imports as package "cli"
	_ "google.golang.org/api/calendar/v3"
)

func main() {
	ctx := context.Background()
	svc, err := internal.Load(ctx)
	fmt.Printf("svc: %v, err: %v\n", svc, err)
}

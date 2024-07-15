// Package doge
package main

import (
	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/awserr"
	_ "github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "github.com/slack-go/slack"
	_ "github.com/urfave/cli/v2" // imports as package "cli"
	_ "google.golang.org/api/calendar/v3"
)

func main() {

}

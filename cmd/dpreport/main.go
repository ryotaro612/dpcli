// Package main
package main

import (
	"context"
	"fmt"
	"log"
	// _ "github.com/aws/aws-sdk-go-v2/aws"
	// _ "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	// _ "github.com/aws/aws-sdk-go/aws/awserr"
	// _ "github.com/aws/aws-sdk-go/aws/session"
	"github.com/ryotaro612/dpcli/internal"
	// _ "github.com/slack-go/slack"
	// _ "github.com/urfave/cli/v2" // imports as package "cli"
	// _ "google.golang.org/api/calendar/v3"

	"github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "log/slog"
)

func main() {
	ctx := context.Background()

}

func main2() {
	ctx := context.Background()
	svc, err := internal.Load(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	secretName := "secret"

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}
	//fmt.Printf("svc: %v, err: %v\n", svc, err)

	result, err := svc.GetSecretValue(ctx, input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	//var secretString string = *result.SecretString
	// JSONになる
	fmt.Printf("%v\n", *result.SecretString)
	fmt.Printf("OK")
}

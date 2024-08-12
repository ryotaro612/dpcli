package internal

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log/slog"
)

// Secret is a struct that aggregates credentias.
type Secret struct {
	GithubToken string
}

// secretClient
type secretClient struct {
	logger *slog.Logger
	client *secretsmanager.Client
}

func ReadSecret(ctx context.Context, logger *slog.Logger, awsProfile string) (Secret, error) {
	c, err := newSecretClient(ctx, logger, awsProfile)
	if err != nil {
		return Secret{}, err
	}
	s, err := c.secret(ctx)
	return s, err

}

// newSecretClient creates a new SecretClient object.
func newSecretClient(ctx context.Context, logger *slog.Logger, awsProfile string) (secretClient, error) {

	c, err := newSecretManagerClient(ctx, awsProfile)
	if err != nil {
		return secretClient{}, err
	}
	return secretClient{logger, c}, nil
}

func (s secretClient) secret(ctx context.Context) (Secret, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String("dpcli"),
		VersionStage: aws.String("AWSCURRENT"),
	}

	output, err := s.client.GetSecretValue(ctx, input)
	var secret Secret
	if err != nil {
		return secret, err
	}
	secretString := output.SecretString
	if err := json.Unmarshal([]byte(*secretString), &secret); err != nil {
		return secret, err
	}
	return secret, nil

}

// newSecretManagerClient loads a secret mananger client.
// if awsProfile is an empty string, it uses the default profile.
func newSecretManagerClient(ctx context.Context, awsProfile string) (*secretsmanager.Client, error) {
	var cfg aws.Config
	region := config.WithRegion("ap-northeast-1")
	var err error
	if awsProfile == "" {
		cfg, err = config.LoadDefaultConfig(ctx, region)
	} else {
		cfg, err = config.LoadDefaultConfig(ctx, region, config.WithSharedConfigProfile(awsProfile))
	}
	if err != nil {
		return nil, err
	}
	return secretsmanager.NewFromConfig(cfg), nil

}

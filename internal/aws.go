package internal

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log/slog"
)

// SecretClient
type SecretClient struct {
	logger *slog.Logger
	client *secretsmanager.Client
}

// NewSecretClient creates a new SecretClient object.
func NewSecretClient(ctx context.Context, logger *slog.Logger, awsProfile string) (SecretClient, error) {

	c, err := newSecretManagerClient(ctx, awsProfile)
	if err != nil {
		return SecretClient{}, err
	}
	return SecretClient{logger, c}, nil
}

type Secret struct {
	GithubToken string
}

func (s SecretClient) Secret(ctx context.Context) (Secret, error) {

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String("dpcli"),
		VersionStage: aws.String("AWSCURRENT"),
	}

	output, err := s.client.GetSecretValue(ctx, input)
	if err != nil {
		return Secret{}, err
	}
	secretString := output.SecretString
	s.logger.Debug("secret", "secret", secretString)
	var secret Secret
	if err := json.Unmarshal([]byte(*secretString), &secret); err != nil {
		return Secret{}, err
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

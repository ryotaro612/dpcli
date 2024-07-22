package internal

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func Load(ctx context.Context) (*secretsmanager.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"), config.WithSharedConfigProfile("sandbox"))
	var svc *secretsmanager.Client
	if err != nil {
		return svc, err
	}
	svc = secretsmanager.NewFromConfig(cfg)
	return svc, nil
}

// NewSecretManagerClient loads a secret mananger client.
// if awsProfile is an empty string, it uses the default profile.
func NewSecretManagerClient(ctx context.Context, awsProfile string) (*secretsmanager.Client, error) {
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

// SecretClient
type SecretClient struct {
	client *secretsmanager.Client
}

// NewSecretClient creates a new SecretClient object.
func NewSecretClient(c *secretsmanager.Client) SecretClient {
	return SecretClient{c}
}

type Secret struct {
}

func (s SecretClient) Secret(ctx context.Context, secretName string) (Secret, error) {

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
		//VersionStage: aws.String("AWSCURRENT"),
	}

	output, err := s.client.GetSecretValue(ctx, input)

	if err != nil {
		return Secret{}, err
	}
	return Secret{}, nil

}

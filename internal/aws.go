package internal

import (
	"context"

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

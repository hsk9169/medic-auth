package auth

import (
	"context"
	"errors"
	"time"

	"github.com/medic-basic/auth/pkg/domain/model"
	"github.com/medic-basic/auth/pkg/external/cache"
)

type Cache interface {
	Get(ctx context.Context, key any, value any) (any, error)
	Set(ctx context.Context, key any, value any, ttl time.Duration) error
	Delete(ctx context.Context, key any) error
}

type Config struct {
	Cache *cache.Cache
}

var cfg *Config

type Aggregate struct {
	Cache Cache
}

func Init(config Config) error {
	if config.Cache == nil {
		return errors.New("cacheClient is nil")
	}
	cfg = &config
	return nil
}

func New() model.AuthAggregate {
	return &Aggregate{
		Cache: cfg.Cache,
	}
}

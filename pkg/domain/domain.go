package domain

import (
	"github.com/gin-gonic/gin"
	auth "github.com/medic-basic/auth/pkg/domain/auth"
	"github.com/medic-basic/auth/pkg/external/cache"
	_ "github.com/medic-basic/auth/pkg/external/redis"
	"github.com/medic-basic/auth/pkg/handler/common"
	"github.com/pkg/errors"
)

type Config struct {
	//RedisClient *redis.Cache
	CacheClient *cache.Cache
}

func Init(cfg Config) error {
	// cfg = cfg
	//if cfg.RedisClient == nil {
	//	return errors.New("redisClient is nil")
	//}

	if cfg.CacheClient == nil {
		return errors.New("cacheClient is nil")
	}

	if err := auth.Init(auth.Config{
		Cache: cfg.CacheClient,
	}); err != nil {
		return errors.Wrap(err, "failed to init auth")
	}

	return nil
}

func SetupDomain(ctx *gin.Context) {
	//ctx.Set(common.QueryKeyName, query.New())
	ctx.Set(common.AuthKeyName, auth.New())
}

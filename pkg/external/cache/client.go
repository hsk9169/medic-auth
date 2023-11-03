package cache

import (
	"context"
	"fmt"
	"time"

	gocache_store "github.com/eko/gocache/store/go_cache/v4"
	gocache "github.com/patrickmn/go-cache"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/marshaler"
	"github.com/eko/gocache/lib/v4/metrics"
	"github.com/eko/gocache/lib/v4/store"
)

var CacheImpl Cache

type Cache struct {
	marshaler *marshaler.Marshaler
}

func Init() {
	client, err := NewCache(5*time.Minute, 10*time.Minute) // number of keys per Get buffer

	if err != nil {
		errMsg := fmt.Sprintf("Failed to init ristretto cache: %s", err)
		panic(errMsg)
	}

	cacheClient := client
	CacheImpl = *cacheClient
}

func NewCache(defaultExpiration time.Duration, cleanupInterval time.Duration) (*Cache, error) {

	gocacheClient := gocache.New(defaultExpiration, cleanupInterval)
	gocacheStore := gocache_store.NewGoCache(gocacheClient)

	var caches []cache.SetterCacheInterface[any]

	caches = append(caches, cache.New[any](gocacheStore))

	cacheWithMetric := cache.NewMetric[any](
		metrics.NewPrometheus("medic"),
		cache.NewChain[any](caches...),
	)

	return &Cache{marshaler: marshaler.New(cacheWithMetric)}, nil

}

func (c *Cache) Get(ctx context.Context, key any, value any) (any, error) {
	return c.marshaler.Get(ctx, key, value)
}

func (c *Cache) Set(ctx context.Context, key any, value any, expireTime time.Duration) error {
	return c.marshaler.Set(ctx, key, value, store.WithExpiration(expireTime))

}

func (c *Cache) Delete(ctx context.Context, key any) error {
	return c.marshaler.Delete(ctx, key)
}

func (c *Cache) Clear(ctx context.Context) error {
	return c.marshaler.Clear(ctx)
}

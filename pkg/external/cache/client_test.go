package cache

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCache_성공(t *testing.T) {
	cache, err := NewCache(5*time.Minute, 10*time.Minute)
	assert.Nil(t, err)
	assert.NotNil(t, cache)
}

func Test_Cache_성공(t *testing.T) {
	cache, err := NewCache(5*time.Minute, 10*time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	key := "test"
	val := "123123"

	if err = cache.Set(context.TODO(), key, &val, 100*time.Millisecond); err != nil {
		panic(err)
	}

	rz, err := cache.Get(context.TODO(), key, new(string))
	assert.Equal(t, val, *rz.(*string))
	assert.Nil(t, err)

	time.Sleep(100 * time.Millisecond)

	rz, err = cache.Get(context.TODO(), key, new(string))
	assert.Nil(t, rz)
	assert.ErrorContains(t, err, "value not found in store")

	assert.Nil(t, cache.Delete(context.TODO(), key))
	assert.Nil(t, cache.Clear(context.TODO()))
}

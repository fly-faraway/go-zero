package cache

import (
	"github.com/tal-tech/go-zero/core/stores/internal"
	"github.com/tal-tech/go-zero/core/syncx"
)

type Cache internal.Cache

func NewCache(c CacheConf, statName string, errNotFound error, opts ...Option) Cache {
	return internal.NewCache(c, syncx.NewSharedCalls(), internal.NewCacheStat(statName), errNotFound, opts...)
}


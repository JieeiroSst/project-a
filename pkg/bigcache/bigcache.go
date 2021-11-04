package bigcache

import (
	"github.com/allegro/bigcache"
	"time"
)

type cache struct {
	cache *bigcache.BigCache
}

type Cache interface {
	Get(cookie string)([]byte,error)
	Set(key string,bytes []byte)error
}

func NewBigCache() Cache {
	memory,err:=bigcache.NewBigCache(bigcache.DefaultConfig(30*time.Hour))
	if err!=nil{
		panic(err)
	}
	return &cache{
		cache: memory,
	}
}

func (c *cache) Get(cookie string)([]byte,error){
	bytes,err:=c.cache.Get(cookie)
	return bytes,err
}

func (c *cache) Set(key string,bytes []byte)error{
	return c.cache.Set(key,bytes)
}
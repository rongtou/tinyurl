package bf

import (
	"tinyurl/cache"
	"tinyurl/pkg/bloomfilter"
)

var bf *bloomfilter.BloomFilter

func Init() {
	storage := bloomfilter.NewRedisStorage(cache.Instance().GetClient(), "tinyurl_bf")
	bf = bloomfilter.New(10000000, 3, storage)
}

func Instance() *bloomfilter.BloomFilter {
	return bf
}

package bloomfilter

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestBloomFilter(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
	storage := NewRedisStorage(client, "bf_test")

	bf := New(100000, 3, storage)
	bf.PutString("t1")
	bf.PutString("t2")

	if !bf.ExistString("t1") {
		t.Fatal("t1 should be found")
	}

	if !bf.ExistString("t2") {
		t.Fatal("t2 should be found")
	}

	if bf.ExistString("t3") {
		t.Fatal("t3 should not be found ")
	}

	if bf.ExistString("abc") {
		t.Fatal("abc should not be found ")
	}
}

package bloomfilter

import (
	"github.com/spaolacci/murmur3"
	"sync"
)

type Storage interface {
	Store([]uint64)
	IsSet(uint64) bool
}

type BloomFilter struct {
	lock *sync.RWMutex

	m uint64 // m bits to storage
	n uint64 // the number of inserted elements
	k int    // the number of has function

	storage Storage
}

func New(size uint64, k int, storage Storage) *BloomFilter {
	return &BloomFilter{
		lock:    &sync.RWMutex{},
		m:       size,
		k:       k,
		storage: storage,
	}
}

func (f *BloomFilter) Put(data []byte) {
	f.lock.Lock()
	defer f.lock.Unlock()

	locs := f.locations(data)
	f.storage.Store(locs)
	f.n++
}

func (f *BloomFilter) PutString(data string) {
	f.Put([]byte(data))
}

func (f *BloomFilter) Exist(data []byte) bool {
	f.lock.RLock()
	defer f.lock.RUnlock()

	locs := f.locations(data)
	for _, loc := range locs {
		if !f.storage.IsSet(loc) {
			return false
		}
	}
	return true
}

func (f *BloomFilter) ExistString(data string) bool {
	return f.Exist([]byte(data))
}

func (f *BloomFilter) locations(data []byte) []uint64 {
	locs := make([]uint64, f.k)
	for i := 0; i < f.k; i++ {
		loc := murmur3.Sum64WithSeed(data, uint32(i))
		locs[i] = loc % f.m
	}
	return locs
}

package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct{
	createdAt time.Time
	val []byte
}

type PokeCache struct{
	Entries map[string]CacheEntry
	mu *sync.Mutex
}
type Cache interface{
	Add(string, []byte)
	Get(string) ([]byte, bool)
	reapLoop()
}

func (pc PokeCache) Add(s string, v []byte){
	pc.Entries[s] = CacheEntry{
		createdAt: time.Now(),
		val: v,
	}
}
func (pc PokeCache) Get(s string) (val []byte, exists bool )

// func NewCache(interval time.Duration) Cache {

// }


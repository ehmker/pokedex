package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

type PokeCache struct {
	Entries map[string]CacheEntry
	mu *sync.Mutex
}
type Cache interface {
	Add(string, []byte)
	Get(string) ([]byte, bool)
	reapLoop()
}

func NewCache(interval time.Duration) PokeCache {
	pc := PokeCache{
		Entries: make(map[string]CacheEntry),
		mu: &sync.Mutex{},
	}
	go pc.reapLoop(interval)
	return pc
}

func (pc PokeCache) Add(s string, v []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.Entries[s] = CacheEntry{
		createdAt: time.Now(),
		val: v,
	}
}
func (pc PokeCache) Get(s string) ([]byte, bool ) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	entry, ok := pc.Entries[s]
	return entry.val, ok
}

func (pc PokeCache) reapLoop(i time.Duration) {
	for {
		pc.mu.Lock()
		for key, entry := range pc.Entries {
			if time.Since(entry.createdAt) < i {
				delete(pc.Entries, key)
			}
		}
		pc.mu.Unlock()
		time.Sleep(i)
	}
}





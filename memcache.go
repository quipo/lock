package lock

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Memcache struct {
	Prefix string
	Cache  *memcache.Client
}

// Acquire a lock for the given key. Returns true on success, false on failure.
// Can attempt up to maxTries times, waiting for waitTime seconds between attempts
func (m Memcache) Acquire(key string, expiryTime int32, waitTime time.Duration, maxTries int) bool {
	item := &memcache.Item{Key: m.getLockKey(key), Value: []byte{'d'}, Expiration: expiryTime}
	for tries := 0; tries < maxTries; tries++ {
		if tries > 0 {
			time.Sleep(waitTime)
		}
		err := m.Cache.Add(item)
		if memcache.ErrNotStored != err {
			return true // yay! lock acquired
		}
	}
	// we ended up using up all the retries
	return false
}

// Release the lock
func (m Memcache) Release(key string) error {
	err := m.Cache.Delete(m.getLockKey(key))
	if err == nil || err == memcache.ErrCacheMiss {
		return nil
	}
	return err
}

func (m Memcache) getLockKey(key string) string {
	return m.Prefix + key
}

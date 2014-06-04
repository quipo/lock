package lock

import (
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

const testServer = "localhost:11211"

func TestGetLockKey(t *testing.T) {
	m := &Memcache{Prefix: "test:"}
	if m.getLockKey("k") != "test:k" {
		t.Errorf("Invalid lock key")
	}
}

func TestAcquire(t *testing.T) {
	key := "k"
	mc := memcache.New(testServer)
	m := &Memcache{Prefix: "test:", Cache: mc}
	m.Release(key)
	defer m.Release(key)

	if !m.Acquire(key, 2, time.Microsecond, 1) {
		t.Error("Cannot acquire first lock")
	}
	if m.Acquire(key, 2, time.Microsecond, 1) {
		t.Error("Can acquire lock when already acquired")
	}
	if nil != m.Release(key) {
		t.Error("Cannot release log")
	}
}

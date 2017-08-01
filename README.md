Memcache-based lock utility in Golang
=====================================

[![Build Status](https://travis-ci.org/quipo/lock.png?branch=master)](https://travis-ci.org/quipo/lock) 
[![GoDoc](https://godoc.org/github.com/quipo/lock?status.png)](http://godoc.org/github.com/quipo/lock)

## Introduction

Simple lock acquire/release functions based on memcache.
The library can automatically retry (and wait between attempts) in case of failure to acquire the lock.

## Installation

    go get github.com/quipo/lock

## Sample usage

```go
package main

import (
	"time"
	
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/quipo/lock"
)

func main() {
	var expiryTime int32 = 10           // in seconds
	waitUtime := 250 * time.Millisecond // wait 250ms between two attempts to acquire the lock
	retries := 4                        // retry for 250ms * 4 times = up to 1 seconds 

	// memcache settings
	lockM := lock.Memcache{Prefix: "lock:", Cache: memcache.New("localhost:11211")}


	// attempt acquiring the lock
	acquiredLock := lockM.Acquire("somekey", expiryTime, waitUtime, retries)
	if acquiredLock {
		defer resolver.Lock.Release(cacheKey)
		// do something
	} else {
		// couldn't acquire lock
	}
}
```


## Author

Lorenzo Alberton

* Web: [http://alberton.info](http://alberton.info)
* Twitter: [@lorenzoalberton](https://twitter.com/lorenzoalberton)
* Linkedin: [/in/lorenzoalberton](https://www.linkedin.com/in/lorenzoalberton)


## Copyright

See [LICENSE](LICENSE) document

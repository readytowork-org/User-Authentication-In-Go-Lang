package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Mutex in go lang
// A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is
// running the critical section of code at any point in time to prevent race conditions
// from happening.

// Mutex is available in the sync package.
// There are two methods defined on Mutex namely Lock and Unlock.
// Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine,
// thus avoiding race condition.

// If one Goroutine already holds the lock and if a new Goroutine
// is trying to acquire a lock, the new Goroutine will be blocked until
// the mutex is unlocked.
func main() {

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

// Expected output
// readOps: 83285
// writeOps: 8320
// state: map[1:97 4:53 0:33 2:15 3:2]

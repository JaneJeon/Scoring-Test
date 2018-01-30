package actor

import (
	"math/rand"
	"strconv"
	"sync"
)

const (
	scoreLimit = 100
)

/**
 * Populates a DB from numerous goroutines, safely and concurrently.
 * Showcases a highly contentious workload, where lock contention will be high.
 * n: number of entries to populate
 * f: percentage of entries that should be nonzero
 */
func Populate(server *server, n int, f float32) {
	// to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(n)

	// shoot off a goroutine for each addition operation
	for i := 0; i < n; i++ {
		go func(i int) {
			if rand.Float32() < f {
				server.set("Post "+strconv.Itoa(i), rand.Intn(scoreLimit-1)+1)
			} else {
				server.set("Post "+strconv.Itoa(i), 0)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Interact(server *server) {
	//
}

// TODO: add spinner

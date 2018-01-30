package actor

import (
	"fmt"
	"sync"
)

/**
 * A "scores server", mimics Redis.
 * It supports a few simple commands for CRUD operations.
 * Designed to be safely concurrently accessed by numerous goroutines.
 * No sharding (yet)!
 */
type server struct {
	scores map[string]int
	mu     sync.RWMutex
}

func NewServer() server {
	return server{scores: make(map[string]int)}
}

func (s *server) String() string {
	s.mu.RLock()
	var base = "Scores Server:\n"
	for k, v := range s.scores {
		base += fmt.Sprintf("%s: %d\n", k, v)
	}
	s.mu.RUnlock()
	return base[:len(base)-1]
}

// concurrently safe way of accessing value of a key in map
func (s *server) get(key string) int {
	s.mu.RLock()
	var x = s.scores[key]
	s.mu.RUnlock()
	return x
}

// concurrently safe way of modifying value of a key in map
// can also be used to make a new key in map
func (s *server) set(key string, val int) {
	s.mu.Lock()
	s.scores[key] = val
	s.mu.Unlock()
}

// TODO
func (s *server) del(key string) {
	//
}

/*---------- Commands ----------*/
func (s *server) incr(key string) {
	s.set(key, s.get(key)+1)
}

func (s *server) decr(key string) {
	s.set(key, s.get(key)-1)
}

func (s *server) multby(key string, f float32) {
	s.set(key, int(f*float32(s.get(key))))
}

package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"./actor"
)

func main() {
	// benchmark setup
	var cores = flag.Int("cores", 1, "Number of cores to use")
	var count = flag.Int("count", 1<<21, "Number of goroutines to spawn")
	flag.Parse()

	runtime.GOMAXPROCS(*cores)
	var start = time.Now()

	// spin up a test server
	var server = actor.NewServer()

	// fill with dummy data
	actor.Populate(&server, *count, 0.5)
	fmt.Printf("Inserting %d items into the database with %d core(s) took %s\n",
		*count, *cores, time.Since(start))
}

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"./actor"
)

func main() {
	// benchmark setup
	var benchmark = flag.Bool("benchmark", false, "Run this in benchmark mode")
	var cores = flag.Int("cores", 1, "Number of cores to use")
	var count = flag.Int("count", 1<<21, "Number of goroutines to spawn")
	flag.Parse()

	if *benchmark {
		rand.Seed(42)
		fmt.Println("Running in benchmark mode")
	}
	runtime.GOMAXPROCS(*cores)
	var start = time.Now()

	// spin up a test server
	var server = actor.NewServer()

	// fill with dummy data
	actor.Populate(&server, *count, 0.5)
	fmt.Printf("Inserting %d items into the database with %d core(s) took %s\n",
		*count, *cores, time.Since(start))
}

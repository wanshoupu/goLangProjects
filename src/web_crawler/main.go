package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

/**
The web crawler will utilize multiple threads.
It will be able to crawl all the particular web pages of a website.
It will be able to report back any 2XX and 4XX links.
It will take in the domain name from the command line.
It will avoid the cyclic traversal of links.

*/
func main() {
	seed, concurrency := os.Args[1], os.Args[2]
	runJob(concurrency, seed)
}

func runJob(concurrency string, seed string) {
	visited := make(map[string]int)
	concur, err := strconv.Atoi(concurrency)
	if err != nil {
		log.Fatalf("error parsing cmd line argument: %v", err)
	}
	queue := NewQueue(concur)
	queue.Push(seed)

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go worker(visited, queue, &wg, 1000)
	}
	wg.Wait()
	fmt.Println("done")
	for k, v := range visited {
		fmt.Println(k, v)
	}
}

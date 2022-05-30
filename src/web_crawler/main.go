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
	seed, concurrency, total := os.Args[1], os.Args[2], os.Args[3]
	runJob(concurrency, seed, total)
}

func runJob(concurrencyStr string, seed string, totalStr string) {
	concur := atoi(concurrencyStr)
	total := atoi(totalStr)

	visited := make(map[string]int, concur)
	queue := NewQueue(concur)
	queue.Push(seed)

	var wg sync.WaitGroup
	wg.Add(concur)
	for i := 0; i < concur; i++ {
		go worker(visited, queue, &wg, total)
	}
	wg.Wait()
	fmt.Println("done")
	for k, v := range visited {
		fmt.Println(k, v)
	}
}

func atoi(concurrency string) int {
	concur, err := strconv.Atoi(concurrency)
	if err != nil {
		log.Fatalf("error parsing cmd line argument: %v", err)
	}
	return concur
}

package main

import (
	"flag"
	"fmt"
	"time"

	"dc-simulator/internal/pool"
)

const asciiBanner = `
 ____   ____       _____ _                 _       _             
|  _ \ / ___|     / ____(_)               | |     | |            
| | | | |   _____| (___  _ _ __ ___  _   _| | __ _| |_ ___  _ __ 
| | | | |  |______\___ \| | '_ ' _ \| | | | |/ _' | __/ _ \| '__|
| |_| | |___      ____) | | | | | | | |_| | | (_| | || (_) | |   
|____/ \____|    |_____/|_|_| |_| |_|\__,_|_|\__,_|\__\___/|_|   
                                                                 
          Distributed Computing Simulator in Go
`

func main() {
	// Command-line parameters configuration
	numWorkers := flag.Int("workers", 4, "Number of workers (nodes)")
	dataSize := flag.Int("size", 100, "Total number of integers to generate")
	chunkSize := flag.Int("chunk", 10, "Size of data chunks (jobs)")
	flag.Parse()

	// Display startup banner
	fmt.Println(asciiBanner)
	fmt.Printf("▶ Parameters: Data = %d | Workers = %d | Chunk size = %d\n", *dataSize, *numWorkers, *chunkSize)
	fmt.Println("------------------------------------------------------------------")

	// Test data generation
	numbers := make([]int, *dataSize)
	for i := 0; i < *dataSize; i++ {
		numbers[i] = i + 1
	}

	start := time.Now()

	// Pool execution
	totalSum := pool.Run(*numWorkers, numbers, *chunkSize)

	// Final results
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("✔ Total sum of squares : %d\n", totalSum)
	fmt.Printf("⏱ Execution time       : %v\n", time.Since(start))
}
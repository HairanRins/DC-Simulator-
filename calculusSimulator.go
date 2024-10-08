// Description:
// * Divide a large list of numbers into several subsets.
// * Each subset will be processed by a "worker" (a goroutine) that will calculate the sum of squares of the numbers in that subset.
// * The results will be collected in a channel.
// * Once all goroutines are finished, we will calculate the total sum of squares from the partial results.

package main

import (
	"fmt"
	"sync"
)

func sumOfSquares(numbers []int, results chan<- int, wg *sync.WaitGroup)  {
	defer wg.Done()
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	results <- sum 
}

func main()  {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	
	numWorkers := 4

	chunkSize := len(numbers) / numWorkers

	results := make(chan int, numWorkers)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(numbers)
		}
		wg.Add(1)
		go sumOfSquares(numbers[start:end], results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := 0
	for result := range results {
		totalSum += result
	}

	fmt.Println("Somme totale des carrés: ", totalSum)
}
package pool

import (
	"fmt"
	"sync"

	"dc-simulator/internal/model"
	"dc-simulator/internal/worker"
)

func Run(numWorkers int, numbers []int, chunkSize int) int {
	numJobs := (len(numbers) + chunkSize - 1) / chunkSize
	jobs := make(chan model.Job, numJobs)
	results := make(chan model.Result, numJobs)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(wID int) {
			defer wg.Done()
			worker.Process(wID, jobs, results)
		}(i)
	}

	jobID := 1
	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		jobs <- model.Job{ID: jobID, Numbers: numbers[i:end]}
		jobID++
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := 0
	for res := range results {
		fmt.Printf("[Réseau] Worker %d a terminé le Job %d -> Somme locale: %d\n", res.WorkerID, res.JobID, res.Sum)
		totalSum += res.Sum
	}

	return totalSum
}
package worker

import (
	"time"

	"dc-simulator/internal/model"
)

func Process(workerID int, jobs <-chan model.Job, results chan<- model.Result) {
	for j := range jobs {
		sum := 0
		for _, num := range j.Numbers {
			sum += num * num
		}

		time.Sleep(50 * time.Millisecond)

		results <- model.Result{
			JobID:    j.ID,
			WorkerID: workerID,
			Sum:      sum,
		}
	}
}
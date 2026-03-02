package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type JobResult struct {
	JobID int
	Value int
}

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- JobResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
			fmt.Printf("Worker %d: exiting due to context cancel\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}

			select {
			case <-time.After(5 * time.Second):
				fmt.Printf("Worker %d processing job %d\n", id, job)
				results <- JobResult{JobID: job, Value: job * 2}
			case <-ctx.Done():
				fmt.Printf("Worker %d: cancelled during work\n", id)
				return
			}
		}
	}
}

func WorkerPool() {
	const numWorkers = 3
	const numJobs = 15

	jobs := make(chan int, numJobs)
	results := make(chan JobResult, numJobs)
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	resMap := make(map[int]int)
	for r := range results {
		resMap[r.JobID] = r.Value
	}

	for i := 1; i <= numJobs; i++ {
		fmt.Printf("Job %d Result %d\n", i, resMap[i])
	}
}

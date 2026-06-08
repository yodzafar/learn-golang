package practice

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ctxWorker(ctx context.Context, id int, jobs <-chan int, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-time.After(2 * time.Second):

		case <-ctx.Done():
			fmt.Printf("Worker %d: exiting due to context cancel\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}

			select {
			case <-time.After(2 * time.Second):
				fmt.Printf("Worker %d processing job %d\n", id, job)
				result <- fmt.Sprintf("Worker %d processing job %d", id, job)
			case <-ctx.Done():
				fmt.Printf("Worker %d: cancelled during work\n", id)
				return
			}
		}
	}
}

func CancellableWorkerPool() {
	const numWorkers = 3
	const numJobs = 20

	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go ctxWorker(ctx, w, jobs, results, wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("Result: ", r)
	}

}

package practice

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(500 * time.Millisecond)
		result <- fmt.Sprintf("worker %d has done %d", id, job)
	}

}

func ConcurrentWorkerPool() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}

}

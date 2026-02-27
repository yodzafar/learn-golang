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
			if !ok { // channel yopilgan
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Second * 5) // ishni simulyatsiya qilish
			results <- JobResult{JobID: job, Value: job * 2}
		}
	}
}

func WorkerPool() {
	const numWorkers = 3
	const numJobs = 15

	jobs := make(chan int, numJobs)
	results := make(chan JobResult, numJobs)
	var wg sync.WaitGroup

	// Context timeout yoki cancel qo‘yish mumkin
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Workerlarni ishga tushirish
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	// Joblarni yuborish
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // channel yopildi → workerlar loop tugaydi

	// Natijalar uchun go routine: WaitGroup tugashini kutadi
	go func() {
		wg.Wait()
		close(results)
	}()

	// Results map bilan job order saqlash
	resMap := make(map[int]int)
	for r := range results {
		resMap[r.JobID] = r.Value
	}

	// Job order bilan natijalarni chiqarish
	for i := 1; i <= numJobs; i++ {
		fmt.Printf("Job %d Result %d\n", i, resMap[i])
	}

	fmt.Println("All jobs processed successfully!")
}

package practice

import (
	"fmt"
	"sort"
	"sync"
)

func squares(x int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- x * x
}

func ConcurrentSquares() {
	const numJobs = 10
	result := make(chan int, numJobs)
	var wg sync.WaitGroup

	for n := 1; n <= numJobs; n++ {
		wg.Add(1)
		go squares(n, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	var results = make([]int, 0, numJobs)
	for r := range result {
		results = append(results, r)
	}

	sort.Ints(results)

	for _, v := range results {
		fmt.Printf("pow result: %d\n", v)
	}
}

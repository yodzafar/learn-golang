package concurrency

import (
	"fmt"
	"time"
)

func promise(task func() int) chan int {
	result := make(chan int)

	go func() {
		result <- task()
		close(result)
	}()

	return result
}

func Promise() {
	longRunningTask := func() int {
		time.Sleep(5 * time.Second)
		return 10
	}

	future := promise(longRunningTask)

	fmt.Println("Task has run, We can do something ...")

	result := <-future

	fmt.Println("Result: ", result)
}

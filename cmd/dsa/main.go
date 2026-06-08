package main

import (
	"fmt"

	"learn-golang/internal/dsa"
)

// Scratchpad for data-structures & algorithms practice.
func main() {
	fmt.Println("=== DSA ===")

	nums := []int{1, 3, 5, 7, 9, 11}
	fmt.Println("BinarySearch(7):", dsa.BinarySearch(nums, 7))
	fmt.Println("BinarySearch(8):", dsa.BinarySearch(nums, 8))

	var s dsa.Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for s.Len() > 0 {
		v, _ := s.Pop()
		fmt.Println("pop:", v)
	}
}

package main

import (
	"fmt"

	"learn-golang/internal/codewars"
	"learn-golang/internal/leetCode/easy"
)

// Scratchpad for running LeetCode and Codewars solutions from one place.
func main() {
	fmt.Println("=== LeetCode: Easy ===")
	fmt.Println("1.   Two Sum:          ", easy.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("345. Reverse Vowels:   ", easy.ReverseVowels("hello"))
	fmt.Println("151. Reverse Words:    ", easy.ReverseWords("the sky is blue"))
	fmt.Println("1768 Merge Alternately:", easy.MergeAlternately("abc", "pqr"))

	fmt.Println("\n=== Codewars ===")
	fmt.Println(`StringToNumber("1234"):`, codewars.StringToNumber("1234"))
	fmt.Println(`StringToNumber("-99"): `, codewars.StringToNumber("-99"))
}

package main

import (
	"fmt"
	"golang/leetCode/easy"
)

func main() {
	gcd := easy.GcdOfStrings("LeetCode", "Code")

	fmt.Println(fmt.Sprintf("Great common divisor: %s", gcd))
}

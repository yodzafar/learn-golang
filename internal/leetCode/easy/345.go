package easy

import "strings"

// hello

func ReverseVowels(s string) string {

	isVowel := func(r rune) bool {
		vowels := "aeiouAEIOU"
		return strings.ContainsRune(vowels, r)
	}

	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		for left < right && !isVowel(runes[left]) {
			left++
		}
		for left < right && !isVowel(runes[right]) {
			right--
		}
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}
	return string(runes)

}

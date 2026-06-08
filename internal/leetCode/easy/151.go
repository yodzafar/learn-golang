package easy

func ReverseWords(s string) string {
	words := []byte(s)
	n := len(words)
	start := 0

	// Reverse the entire string
	for i := 0; i < n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}

	// Reverse each word in the reversed string
	for i := 0; i <= n; i++ {
		if i == n || words[i] == ' ' {
			end := i - 1
			for j := start; j < (start+end+1)/2; j++ {
				words[j], words[end-(j-start)] = words[end-(j-start)], words[j]
			}
			start = i + 1
		}
	}

	return string(words)
}
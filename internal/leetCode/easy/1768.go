package easy

func MergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	minLength := len(word1)
	if len(word2) < minLength {
		minLength = len(word2)
	}

	for i := 0; i < minLength; i++ {
		result = append(result, word1[i], word2[i])
	}

	if len(word1) > minLength {
		result = append(result, word1[minLength:]...)
	} else if len(word2) > minLength {
		result = append(result, word2[minLength:]...)
	}

	return string(result)
}

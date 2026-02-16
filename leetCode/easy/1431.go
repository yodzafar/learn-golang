package easy

// KidsWithCandies Kids With the Greatest Number of Candies
func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	max := candies[0]

	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	for i, c := range candies {
		result[i] = c+extraCandies >= max
	}

	return result
}

package easy

func GcdOfStrings(str1 string, str2 string) string {

	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}

		return a
	}

	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

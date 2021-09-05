package strings

// IsPalindrome checks if the word is a palindrome, meaning that it is the same reading backwards.
func IsPalindrome(str string) bool {
	size := len(str) - 1
	for i, char := range str {
		if rune(str[size-i]) != char {
			return false
		}
	}

	return true
}

// PalindromeIndex determines the index of a character that can be removed to make the string a
// palindrome assuming only lowercased letters [a-z] as input.
func PalindromeIndex(s string) int {
	last := len(s) - 1

	for i := 0; i < len(s)/2; i++ { // O(n)
		tail := last - i
		if rune(s[tail]) != rune(s[i]) {
			if IsPalindrome(s[i:tail]) { // O(n)
				return tail
			} else if IsPalindrome(s[i+1 : tail+1]) { // O(n)
				return i
			} else {
				return -1
			}
		}
	}

	return -1
}

func PalindromeIndexLinear(s string) int {
	l := len(s)
	i := 0
	j := l - 1

	for i < l {
		if s[i] != s[j] {
			break
		}

		i += 1
		j -= 1
	}

	if i > j {
		return -1
	}

	a := i + 1
	b := j

	if a == b {
		return a
	}

	for a < j && b > (i+1) {
		if s[a] != s[b] {
			return j
		}

		a += 1
		b -= 1
	}

	return i
}

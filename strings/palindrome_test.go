package strings_test

import (
	"algorithms/strings"
	"bufio"
	"log"
	"os"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		Input    string
		Expected bool
	}{
		{"arara", true},
		{"aia", true},
		{"aibofobia", true},
		{"ala", true},
		{"ama", true},
		{"anilina", true},
		{"ata", true},
		{"arara", true},
		{"asa", true},
		{"ele", true},
		{"esse", true},
		{"mamam", true},
		{"matam", true},
		{"metem", true},
		{"mirim", true},
		{"oco", true},
		{"caramba", false},
	}

	for _, test := range tt {
		if got := strings.IsPalindrome(test.Input); test.Expected != got {
			t.Errorf("expected is palindrome for '%s' to be %v got %v", test.Input, test.Expected, got)
		}
	}
}

var (
	inputs = openInput("../_assets/strings/palindrome-test-14-input")
)

func TestIsPalindromeIndex(t *testing.T) {
	tt := []struct {
		Input    string
		Expected int
	}{
		{"arabra", 3},
		{inputs[0], 16722},
		{inputs[1], 76661},
		{inputs[2], 10035},
		{inputs[3], 57089},
		{inputs[4], 46674},
	}

	for _, test := range tt {
		if got := strings.PalindromeIndex(test.Input); test.Expected != got {
			t.Errorf("expected %d got %d for '%s'", test.Expected, got, test.Input)
		}
	}
}

func TestIsPalindromeIndexLinear(t *testing.T) {
	tt := []struct {
		Input    string
		Expected int
	}{
		{"arabra", 3},
		{"oob", 2},
		{inputs[0], 16722},
		{inputs[1], 76661},
		{inputs[2], 10035},
		{inputs[3], 57089},
		{inputs[4], 46674},
	}

	for _, test := range tt {
		if got := strings.PalindromeIndexLinear(test.Input); test.Expected != got {
			t.Errorf("expected %d got %d for '%s'", test.Expected, got, test.Input)
		}
	}
}

func BenchmarkPalindromeIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.PalindromeIndex(inputs[3])
	}
}

func BenchmarkPalindromeIndexLinear(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.PalindromeIndexLinear(inputs[3])
	}
}

func openInput(fileName string) []string {
	file, _ := os.Open(fileName)
	defer file.Close()

	sc := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)

	var inputs []string
	for sc.Scan() {
		inputs = append(inputs, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return inputs
}

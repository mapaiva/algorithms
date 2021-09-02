package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestRunningMedian_1(t *testing.T) {
	items := []int{5, 2, 3, 1, 6, 4}
	expected := []float64{5, 3.5, 3, 2.5, 3, 3.5}

	medians := runningMedian(items)

	if medians == nil {
		t.Fatalf("expected medians to not be nil")
	}
	for i, median := range medians {
		if median != expected[i] {
			t.Errorf("expected %dth median to be %f got %f", i, expected[i], median)
		}
	}
}

func TestRunningMedian_2(t *testing.T) {
	items := []int{
		94455, 20555, 20535, 53125, 73634, 148, 63772, 17738, 62995, 13401, 95912, 13449, 92211,
		17073, 69230, 22016, 22120, 78563, 16571,
	}
	expected := []float64{
		94455.0, 57505.0, 20555.0, 36840.0, 53125.0, 36840.0, 53125.0, 36840.0, 53125.0, 36840.0,
		53125.0, 36840.0, 53125.0, 36840.0, 53125.0, 37570.5, 22120.0, 37622.5, 22120.0,
	}

	medians := runningMedian(items)

	if medians == nil {
		t.Fatalf("expected medians to not be nil")
	}
	for i, median := range medians {
		if median != expected[i] {
			t.Errorf("expected %dth median to be %f got %f", i, expected[i], median)
		}
	}
}

func TestRunningMedian_3(t *testing.T) {
	items := openInput("../../_assets/heap/test-1-input")
	expected := openOutput("../../_assets/heap/test-1-output")

	medians := runningMedian(items)

	if medians == nil {
		t.Fatalf("expected medians to not be nil")
	}
	for i, median := range medians {
		if median != expected[i] {
			t.Errorf("expected %dth median to be %f got %f", i, expected[i], median)
		}
	}
}

func openInput(fileName string) []int {
	file, _ := os.Open(fileName)
	defer file.Close()

	var items []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		items = append(items, i)
	}
	return items
}

func openOutput(fileName string) []float64 {
	file, _ := os.Open(fileName)
	defer file.Close()

	var items []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.ParseFloat(scanner.Text(), 0)
		items = append(items, i)
	}
	return items
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Parse command line arguments
	meanFlag := flag.Bool("mean", false, "Print mean")
	medianFlag := flag.Bool("median", false, "Print median")
	modeFlag := flag.Bool("mode", false, "Print mode")
	sdFlag := flag.Bool("sd", false, "Print standard deviation")
	flag.Parse()

	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0)
	stopKeyword := "stop"

	for scanner.Scan() {
		text := scanner.Text()

		if text == stopKeyword {
			break
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Invalid input:", err)
			os.Exit(1)
		}
		numbers = append(numbers, num)
	}

	// Sort the numbers
	sort.Ints(numbers)

	// Calculate mean
	mean := calculateMean(numbers)

	// Calculate median
	median := calculateMedian(numbers)

	// Calculate mode
	mode := calculateMode(numbers)

	// Calculate standard deviation
	sd := calculateStandardDeviation(numbers, mean)

	// Print the requested metrics
	if *meanFlag {
		fmt.Printf("Mean: %.2f\n", mean)
	}
	if *medianFlag {
		fmt.Printf("Median: %.2f\n", median)
	}
	if *modeFlag {
		fmt.Printf("Mode: %d\n", mode)
	}
	if *sdFlag {
		fmt.Printf("SD: %.2f\n", sd)
	}
}

func calculateMean(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func calculateMedian(numbers []int) float64 {
	size := len(numbers)
	middle := size / 2

	if size%2 == 0 {
		return float64(numbers[middle-1]+numbers[middle]) / 2.0
	} else {
		return float64(numbers[middle])
	}
}

func calculateMode(numbers []int) int {
	counts := make(map[int]int)
	for _, num := range numbers {
		counts[num]++
	}

	var mode int
	maxCount := 0
	for num, count := range counts {
		if count > maxCount || (count == maxCount && num < mode) {
			mode = num
			maxCount = count
		}
	}

	return mode
}

func calculateStandardDeviation(numbers []int, mean float64) float64 {
	var sumSquares float64
	for _, num := range numbers {
		diff := float64(num) - mean
		sumSquares += diff * diff
	}
	variance := sumSquares / float64(len(numbers))
	return math.Sqrt(variance)
}

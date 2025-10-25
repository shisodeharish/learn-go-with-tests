package main

// for loop - Sum calculates the total from a slice of numbers.
/*
func Sum1(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
} */

// range - Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// slice for loop - Sum calculates the total from a slice of numbers.
/*
func SumAll(nubersToSum ...[]int) []int {
	lengthOfNumbers := len(nubersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range nubersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
} */

// slice append - Sum calculates the total from a slice of numbers.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails calculates the respective sums of every slice passed in.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}

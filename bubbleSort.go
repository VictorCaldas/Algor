package main

import "fmt"

func main() {
	var numbers = []int{5, 80, 2, 3, 1, 0, 9, 7, 6, 4}
	fmt.Println("Unsorted:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}

func bubbleSort(numbers []int) {
	var N = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		sweep(numbers, i)
	}
}

func sweep(numbers []int, prevPasses int) {
	var N = len(numbers)
	var firstIndex = 0
	var secondIndex = 1

	for secondIndex < (N - prevPasses) {
		var firstNumber = numbers[firstIndex]
		var secondNumber = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}

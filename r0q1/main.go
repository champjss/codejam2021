// Reversort
// https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d0a5c

package main

import (
	"fmt"
)

type TestCase struct {
	Size     int
	Sequence []int
}

func findMinIndex(values []int, fromIndex int) int {
	size := len(values)
	minIndex := fromIndex

	for i := fromIndex + 1; i < size; i++ {
		if values[i] < values[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}

func reverseSlice(source []int) []int {
	size := len(source)
	reversed := make([]int, size)

	for i := 0; i < size; i++ {
		reversed[i] = source[size-i-1]
	}

	return reversed
}

func solveTestCase(testCase *TestCase) int {
	sequence := testCase.Sequence
	size := testCase.Size
	cost := 0

	for i := 0; i < testCase.Size-1; i++ {
		minIndex := findMinIndex(sequence, i)
		startIndexToReverse := i
		endIndexToReverse := minIndex
		cost += endIndexToReverse - startIndexToReverse + 1

		newSequence := sequence[:startIndexToReverse]
		newSequence = append(
			newSequence,
			reverseSlice(sequence[startIndexToReverse:endIndexToReverse+1])...,
		)
		if endIndexToReverse < size-1 {
			newSequence = append(
				newSequence,
				sequence[endIndexToReverse+1:]...,
			)
		}
		sequence = newSequence
	}

	return cost
}

func readTestCase() *TestCase {
	testCase := &TestCase{}
	fmt.Scanf("%d", &testCase.Size)

	testCase.Sequence = make([]int, testCase.Size)
	for i := 0; i < testCase.Size; i++ {
		fmt.Scanf("%d", &testCase.Sequence[i])
	}

	return testCase
}

func readTestCases() []*TestCase {
	var totalCases int
	fmt.Scanf("%d", &totalCases)

	testCases := make([]*TestCase, totalCases)
	for i := 0; i < totalCases; i++ {
		testCases[i] = readTestCase()
	}

	return testCases
}

func main() {
	testCases := readTestCases()

	for i, testCase := range testCases {
		answer := solveTestCase(testCase)
		fmt.Printf("Case #%d: %d\n", i+1, answer)
	}
}

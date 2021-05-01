// Reversort Engineering
// https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d12d7

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TestCase struct {
	Size int
	Cost int
}

type Answer struct {
	Possible bool
	Sequence []int
}

func (a *Answer) makeAnswerString() string {
	if !a.Possible {
		return "IMPOSSIBLE"
	}

	sequenceString := make([]string, len(a.Sequence))
	for i, v := range a.Sequence {
		sequenceString[i] = strconv.Itoa(v)
	}
	return strings.Join(sequenceString, " ")
}

func cloneAndAppendSlice(source []int, itemToAppend int) []int {
	cloned := make([]int, len(source))
	copy(cloned, source)
	cloned = append(cloned, itemToAppend)
	return cloned
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeCostsByRound(size int, cost int) []int {
	costsByRound := make([]int, size-1)
	remainingExtraCost := cost - (size - 1)

	for i := 0; i < size-1; i++ {
		maxExtraCost := i + 1
		currentRoundExtraCost := min(maxExtraCost, remainingExtraCost)
		remainingExtraCost -= currentRoundExtraCost

		costsByRound[i] = currentRoundExtraCost + 1
	}

	return costsByRound
}

func reverseSlice(source []int) []int {
	size := len(source)
	reversed := make([]int, size)

	for i := 0; i < size; i++ {
		reversed[i] = source[size-i-1]
	}

	return reversed
}

func getNextRoundToReverse(roundToReverse []int, index int) int {
	if index >= len(roundToReverse) {
		return -1
	}

	return roundToReverse[index]
}

func makeNumberRange(size int) []int {
	sequence := make([]int, size)

	for i := 0; i < size; i++ {
		sequence[i] = i + 1
	}

	return sequence
}

func makeNumberSequenceByCosts(size int, costsByRounds []int) []int {
	sequence := makeNumberRange(size)

	for i, cost := range costsByRounds {
		startIndexToReverse := size - i - 2
		endIndexToReverse := startIndexToReverse + cost - 1
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

	return sequence
}

func getMinCostForSize(size int) int {
	return size - 1
}

func getMaxCostForSize(size int) int {
	return (2 + size) * (size - 1) / 2
}

func solveTestCase(testCase *TestCase) *Answer {
	if testCase.Cost < getMinCostForSize(testCase.Size) {
		return &Answer{
			Possible: false,
		}
	}

	if testCase.Cost > getMaxCostForSize(testCase.Size) {
		return &Answer{
			Possible: false,
		}
	}

	costsByRound := makeCostsByRound(testCase.Size, testCase.Cost)

	return &Answer{
		Possible: true,
		Sequence: makeNumberSequenceByCosts(testCase.Size, costsByRound),
	}
}

func readTestCase() *TestCase {
	testCase := &TestCase{}

	fmt.Scanf(
		"%d %d",
		&testCase.Size,
		&testCase.Cost,
	)

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
		fmt.Printf("Case #%d: %s\n", i+1, answer.makeAnswerString())
	}
}

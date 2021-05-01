// Median Sort
// https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d12d7

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	TotalTestCases int
	Size           int
	MaxQuestions   int
}

type Answer struct {
	Sequence []int
}

func (a *Answer) makeAnswerString() string {
	sequenceString := make([]string, len(a.Sequence))
	for i, v := range a.Sequence {
		sequenceString[i] = strconv.Itoa(v)
	}
	return strings.Join(sequenceString, " ")
}

func askForMedianIndex(numbers []int) int {
	var medianIndex int

	fmt.Printf("%d %d %d\n", numbers[0], numbers[1], numbers[2])
	fmt.Scanf("%d", &medianIndex)

	return medianIndex
}

func cloneSlice(source []int) []int {
	cloned := make([]int, len(source))
	copy(cloned, source)
	return cloned
}

func tryFillNewIndexInSequence(sequence []int, newIndex int) []int {
	size := len(sequence)

	indicesToAsk := make([]int, 3)
	for i := 1; i < size; i++ {
		indicesToAsk[0] = sequence[i-1]
		indicesToAsk[1] = sequence[i]
		indicesToAsk[2] = newIndex
		medianIndex := askForMedianIndex(indicesToAsk)
		if medianIndex == newIndex {
			newSequence := cloneSlice(sequence[:i])
			newSequence = append(
				newSequence,
				newIndex,
			)
			newSequence = append(
				newSequence,
				sequence[i:]...,
			)
			return newSequence
		} else if i == 1 && medianIndex == indicesToAsk[0] {
			newSequence := append(
				[]int{newIndex},
				sequence...,
			)
			return newSequence
		}
	}

	return append(sequence, newIndex)
}

func isSliceHasItem(slice []int, item int) bool {
	size := len(slice)
	for i := 0; i < size; i++ {
		if slice[i] == item {
			return true
		}
	}
	return false
}

func tryFillIndexToSequence(sequence []int, size int) []int {
	for i := 4; i <= size; i++ {
		if !isSliceHasItem(sequence, i) {
			sequence = tryFillNewIndexInSequence(sequence, i)
		}
	}

	return sequence
}

func solveProblem(problem *Problem) *Answer {
	// Start with first 3 items
	var sequence = make([]int, 3)
	medianIndex := askForMedianIndex([]int{1, 2, 3})
	if medianIndex == 1 {
		sequence = []int{2, 1, 3}
	} else if medianIndex == 2 {
		sequence = []int{1, 2, 3}
	} else if medianIndex == 3 {
		sequence = []int{1, 3, 2}
	}

	// Try adding one more item, until all of them filled
	sequence = tryFillIndexToSequence(sequence, problem.Size)
	for len(sequence) < problem.Size {
		sequence = tryFillIndexToSequence(sequence, problem.Size)
	}

	return &Answer{
		Sequence: sequence,
	}
}

func readProblem() *Problem {
	problem := &Problem{}

	fmt.Scanf(
		"%d %d %d",
		&problem.TotalTestCases,
		&problem.Size,
		&problem.MaxQuestions,
	)

	return problem
}

func readJudgeResponse() int {
	var response int
	fmt.Scanf("%d", &response)
	return response
}

func main() {
	problem := readProblem()

	for i := 0; i < problem.TotalTestCases; i++ {
		answer := solveProblem(problem)
		fmt.Printf("%s\n", answer.makeAnswerString())
		response := readJudgeResponse()
		if response != 1 {
			os.Exit(3)
		}
	}
}

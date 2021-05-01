// Moons and Umbrellas
// https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d1145

package main

import (
	"fmt"
	"strings"
)

type TestCase struct {
	CJCost  int
	JCCost  int
	Symbols string
}

type WildcardRange struct {
	StartIndex int
	EndIndex   int
}

func replaceSymbolsBetweenIndexes(symbols string, replacement byte, startIndex int, endIndex int) string {
	return symbols[:startIndex] + strings.Repeat(string(replacement), endIndex-startIndex+1) + symbols[endIndex+1:]
}

func fillBestWildcardsBetweenIndexes(symbols string, startIndex int, endIndex int) string {
	size := len(symbols)

	symbolBeforeWildcard := byte(0)
	symbolAfterWildcard := byte(0)
	if startIndex > 0 {
		symbolBeforeWildcard = symbols[startIndex-1]
	}
	if endIndex < size-1 {
		symbolBeforeWildcard = symbols[endIndex+1]
	}

	if symbolBeforeWildcard == byte(0) && symbolAfterWildcard == byte(0) {
		return replaceSymbolsBetweenIndexes(
			symbols,
			'C',
			startIndex,
			endIndex,
		)
	} else if symbolBeforeWildcard == byte(0) {
		return replaceSymbolsBetweenIndexes(
			symbols,
			symbolAfterWildcard,
			startIndex,
			endIndex,
		)
	} else {
		return replaceSymbolsBetweenIndexes(
			symbols,
			symbolBeforeWildcard,
			startIndex,
			endIndex,
		)
	}
}

func findAllWildcardRanges(symbols string) []*WildcardRange {
	size := len(symbols)
	wildcardRanges := []*WildcardRange{}

	var latestWildcardRange *WildcardRange = nil

	for i := 0; i < size; i++ {
		symbol := symbols[i]
		if symbol == '?' {
			if latestWildcardRange == nil {
				// Found new wildcard range
				latestWildcardRange = &WildcardRange{
					StartIndex: i,
					EndIndex:   i,
				}
				wildcardRanges = append(wildcardRanges, latestWildcardRange)
			} else {
				// Expanding the latest wildcard range
				latestWildcardRange.EndIndex = i
			}
		} else {
			latestWildcardRange = nil
		}
	}

	return wildcardRanges
}

func fillAllWildcards(symbols string) string {
	wildcardRanges := findAllWildcardRanges(symbols)

	filledSymbols := symbols
	for _, wildcardRange := range wildcardRanges {
		filledSymbols = fillBestWildcardsBetweenIndexes(
			filledSymbols,
			wildcardRange.StartIndex,
			wildcardRange.EndIndex,
		)
	}

	return filledSymbols
}

func calculateSymbolsCost(cjCost int, jcCost int, symbols string) int {
	size := len(symbols)
	cost := 0

	for i := 1; i < size; i++ {
		symbolPair := symbols[i-1 : i+1]
		if symbolPair == "CJ" {
			cost += cjCost
		} else if symbolPair == "JC" {
			cost += jcCost
		}
	}

	return cost
}

func solveTestCase(testCase *TestCase) int {
	return calculateSymbolsCost(
		testCase.CJCost,
		testCase.JCCost,
		fillAllWildcards(testCase.Symbols),
	)
}

func readTestCase() *TestCase {
	testCase := &TestCase{}

	fmt.Scanf(
		"%d %d %s",
		&testCase.CJCost,
		&testCase.JCCost,
		&testCase.Symbols,
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
		fmt.Printf("Case #%d: %d\n", i+1, answer)
	}
}

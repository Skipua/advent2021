package day10

import (
	"container/list"
	"sort"
	"strings"
)

var pairs = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
	"<": ">",
}

var scores = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func CheckSyntax(input []string) int {

	lineScores := make([]int, 0)

	for _, line := range input {
		stack := list.New()
		corrupted := false
		for _, elem := range strings.Split(line, "") {
			if _, ok := pairs[elem]; ok {
				stack.PushBack(elem)
			} else {
				lastValue := stack.Back().Value.(string)
				if elem == pairs[lastValue] {
					stack.Remove(stack.Back())
				} else {
					corrupted = true
					break
				}
			}
		}

		if !corrupted {
			lineScores = append(lineScores, calculateScore(stack))
		}
	}

	sort.Ints(lineScores)
	return lineScores[len(lineScores)/2]
}

func calculateScore(stack *list.List) int {
	lineScore := 0
	for stack.Len() != 0 {
		lastElem := stack.Back()
		value := lastElem.Value.(string)
		lineScore *= 5
		lineScore += scores[pairs[value]]
		stack.Remove(lastElem)
	}
	return lineScore
}

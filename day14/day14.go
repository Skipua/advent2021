package day14

import (
	"math"
	"strconv"
	"strings"
)

func DoPoly(template string, polyInstructions []string, steps int) int {
	rules := parseRules(polyInstructions)

	result := make(map[string]int, 0)
	split := strings.Split(template, "")
	memo := make(map[string]map[string]int, 0)
	for i := 0; i < len(split)-1; i++ {
		next := poly(split[i], split[i+1], rules, steps, memo)
		if i == 0 {
			result = merge(result, next)
		} else {
			next[split[i]] -= 1
			result = merge(result, next)
		}
	}

	return countMaxMinusMin(result)
}

func poly(p1, p2 string, rules map[string]string, steps int, memo map[string]map[string]int) map[string]int {
	key := p1 + p2 + strconv.Itoa(steps)
	if v, ok := memo[key]; ok {
		return v
	}

	nextMiddle := rules[p1+p2]
	if steps == 1 {
		m := map[string]int{p1: 1}
		addOrIncrement(m, p2)
		addOrIncrement(m, nextMiddle)
		return m
	} else {
		next1 := poly(p1, nextMiddle, rules, steps-1, memo)
		next2 := poly(nextMiddle, p2, rules, steps-1, memo)
		next2[nextMiddle] -= 1
		memo[key] = merge(next1, next2)
		return memo[key]
	}
}

func addOrIncrement(m map[string]int, token string) {
	if _, ok := m[token]; ok {
		m[token] += 1
	} else {
		m[token] = 1
	}
}

func merge(m1 map[string]int, m2 map[string]int) map[string]int {
	newM := make(map[string]int, 0)

	for k := range m1 {
		newM[k] = m1[k]
	}

	for k := range m2 {
		if _, ok := newM[k]; ok {
			newM[k] += m2[k]
		} else {
			newM[k] = m2[k]
		}
	}

	return newM
}

func countMaxMinusMin(result map[string]int) int {
	max := 0
	min := math.MaxInt

	for k := range result {
		count := result[k]
		if max < count {
			max = count
		}
		if min > count {
			min = count
		}
	}
	return max - min
}

func parseRules(instructions []string) map[string]string {
	m := make(map[string]string, 0)
	for _, v := range instructions {
		split := strings.Split(v, " -> ")
		m[split[0]] = split[1]
	}

	return m
}

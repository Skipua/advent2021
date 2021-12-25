package day14

import (
	"math"
	"strconv"
	"strings"
)

func DoPoly(template string, polyInstructions []string, steps int) int {
	rules := parseRules(polyInstructions)

	finalTemplate := ""
	split := strings.Split(template, "")
	memo := make(map[string]string, 0)
	for i := 0; i < len(split)-1; i++ {
		next := poly(split[i], split[i+1], rules, steps, memo)
		if i == 0 {
			finalTemplate += next
		} else {
			finalTemplate += strings.TrimPrefix(next, split[i])
		}
	}

	return countMaxMinusMin(finalTemplate)
}

func poly(p1, p2 string, rules map[string]string, steps int, memo map[string]string) string {
	key := p1 + p2 + strconv.Itoa(steps)
	newToken := rules[p1+p2]

	if v, ok := memo[key]; ok {
		return v
	}

	if steps == 1 {
		return p1 + newToken + p2
	} else {
		next := poly(p1, newToken, rules, steps-1, memo) + strings.TrimPrefix(poly(newToken, p2, rules, steps-1, memo), newToken)
		memo[key] = next
		return next
	}
}

func countMaxMinusMin(template string) int {
	max := 0
	min := math.MaxInt

	m := make(map[string]int, 0)
	for _, v := range strings.Split(template, "") {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	for k := range m {
		count := m[k]
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

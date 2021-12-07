package day6

var cache = make(map[Key]int, 0)

type Key struct {
	days  int
	state int
}

func SimulateLantern(in []int, days int) int {
	totalPopulation := 0
	for _, p := range in {
		totalPopulation += calcPopulation(days, p, 1)
	}
	return totalPopulation
}

func calcPopulation(days, state, population int) int {
	if p, ok := cache[Key{days, state}]; ok {
		return p
	}

	if days == 0 {
		return population
	}
	if state == 0 {
		p := calcPopulation(days-1, 6, population)
		key := Key{days - 1, 6}
		cache[key] = p
		p2 := calcPopulation(days-1, 8, 1)
		cache[Key{days - 1, 8}] = p2

		return p + p2
	} else {
		return calcPopulation(days-1, state-1, population)
	}
}

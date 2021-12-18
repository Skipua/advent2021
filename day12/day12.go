package day12

import (
	"strings"
)

type Caves struct {
	caves map[string][]string
	paths []string
}

func NewCaves() *Caves {
	return &Caves{make(map[string][]string, 0), make([]string, 0)}
}

func (c *Caves) AddPath(from, to string) {
	if _, ok := c.caves[from]; !ok {
		leadsTo := make([]string, 0)
		leadsTo = append(leadsTo, to)
		c.caves[from] = leadsTo
	} else {
		c.caves[from] = append(c.caves[from], to)
	}
}

func (c Caves) Paths() []string {
	c.findPaths("start")
	return c.paths
}

func (c *Caves) findPaths(path string) {
	caves := strings.Split(path, ",")
	currCave := caves[len(caves)-1]

	if currCave == "end" {
		c.paths = append(c.paths, path)
	}
	for _, nextCave := range c.caves[currCave] {
		nextPath := path + "," + nextCave
		if validPath(nextPath) {
			c.findPaths(nextPath)
		}
	}
}

func validPath(path string) bool {
	countPerSmallCave := make(map[string]int, 0)
	split := strings.Split(path, ",")
	for _, c := range split {
		if strings.ToUpper(c) == c {
			continue
		}

		if _, ok := countPerSmallCave[c]; ok {
			countPerSmallCave[c] += 1
		} else {
			countPerSmallCave[c] = 1
		}
	}

	if countPerSmallCave["start"] == 2 || countPerSmallCave["end"] == 2 {
		return false
	}

	countOfCavesVisitedTwice := 0
	for k := range countPerSmallCave {
		if countPerSmallCave[k] > 2 {
			return false
		}

		if countPerSmallCave[k] > 1 {
			countOfCavesVisitedTwice++
		}
	}

	return countOfCavesVisitedTwice <= 1
}

func CountPaths(paths []string) int {
	caves := NewCaves()
	for _, path := range paths {
		pathParts := strings.Split(path, "-")
		from := pathParts[0]
		to := pathParts[1]

		caves.AddPath(from, to)
		caves.AddPath(to, from)
	}

	return len(caves.Paths())
}

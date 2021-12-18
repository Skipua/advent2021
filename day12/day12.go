package day12

import (
	"strings"
)

type Caves struct {
	caves map[string][]string
}

var set = struct{}{}

func NewCaves() *Caves {
	return &Caves{make(map[string][]string, 0)}
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
	//paths := make([]string, 0)
	markedPaths := make(map[string]struct{}, 0)
	markedCaves := make(map[string]struct{}, 0)
	res := make([]string, 0)

	c.findPaths(&res, markedPaths, markedCaves, "start")
	return res
}

func (c *Caves) findPaths(res *[]string, markedPath map[string]struct{}, markedCaves map[string]struct{}, path string) {
	markedPath[path] = set

	//fmt.Printf("Visiting: %v\n", path)

	caves := strings.Split(path, ",")
	currCave := caves[len(caves)-1]
	markedCaves[currCave] = set

	if currCave == "end" {
		//fmt.Printf("END: %v\n", path)
		*res = append(*res, path)
		//markedCaves = make(map[string]struct{}, 0)
	} else if c.caves[currCave] == nil {
		//fmt.Printf("DEAD: %v\n", path)
		//markedCaves = make(map[string]struct{}, 0)
	}
	for _, nextCave := range c.caves[currCave] {
		nextPath := path + "," + nextCave
		if !visitedTwice(nextPath, nextCave) {
			_, pathVisited := markedPath[nextPath]
			_, caveVisited := markedCaves[nextCave]
			if !caveVisited || !pathVisited {
				c.findPaths(res, markedPath, markedCaves, nextPath)
			}
		}
	}
}

func visitedTwice(path string, cave string) bool {
	if strings.ToUpper(cave) == cave {
		return false
	}
	split := strings.Split(path, ",")
	count := 0
	for _, p := range split {
		if p == cave {
			count++
		}

		if count > 1 {
			return true
		}
	}

	return false
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

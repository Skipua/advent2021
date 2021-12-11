package day11

type P struct {
	i       int
	j       int
	flashed bool
}

func (p P) neighbours() []*P {
	ps := make([]*P, 0)
	ps = append(ps, NewP(p.i+1, p.j))
	ps = append(ps, NewP(p.i+1, p.j+1))
	ps = append(ps, NewP(p.i+1, p.j-1))
	ps = append(ps, NewP(p.i-1, p.j-1))
	ps = append(ps, NewP(p.i-1, p.j+1))
	ps = append(ps, NewP(p.i-1, p.j))
	ps = append(ps, NewP(p.i, p.j-1))
	ps = append(ps, NewP(p.i, p.j+1))
	return ps
}

func NewP(i, j int) *P {
	return &P{i, j, false}
}

func CountFlashes(in [][]int, steps int) int {
	totalFlashes := 0
	for step := 0; step < steps; step++ {
		flashes := 0
		flashed := make(map[P]struct{}, 0)
		for i, row := range in {
			for j, _ := range row {
				p := NewP(i, j)
				visited := make(map[*P]struct{}, 0)
				flashes += doStep(in, visited, flashed, p)
			}
		}
		totalFlashes += flashes
	}
	return totalFlashes
}

func doStep(in [][]int, visited map[*P]struct{}, flashed map[P]struct{}, p *P) int {
	if isOut(in, p) {
		return 0
	}
	if _, ok := visited[p]; !ok {
		visited[p] = struct{}{}
		_, flashedThisStep := flashed[*p]
		if in[p.i][p.j] == 0 && flashedThisStep {
			return 0
		}
		in[p.i][p.j] += 1

		if in[p.i][p.j] == 10 {
			flashed[*p] = struct{}{}
			in[p.i][p.j] = 0
			flashes := 1
			for _, p := range p.neighbours() {
				flashes += doStep(in, visited, flashed, p)
			}
			return flashes
		}
	}
	return 0
}

func isOut(in [][]int, p *P) bool {
	return p.i < 0 || p.j < 0 || p.i == len(in) || p.j == len(in[0])
}

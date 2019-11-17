package twoDee

func treeCount(forest [][]int) int {
	n := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if forest[i][j] > 1 {
				n++
			}
		}
	}
	return n
}

type point struct {
	x, y int
}

type stack struct {
	data []point
}

func NewStack() *stack {
	return &stack{make([]point, 0)}
}

func (s *stack) Push(p point) {
	s.data = append(s.data, p)
}

func (s *stack) Pop() point {
	v := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return v
}

func (s *stack) Len() int {
	return len(s.data)
}

func neighbors(p point, width, height int) []point {
	x, y := p.x, p.y

	n := make([]point, 0)
	if p.x > 0 {
		n = append(n, point{x - 1, y})
	}
	if p.x < width-1 {
		n = append(n, point{x + 1, y})
	}
	if p.y > 0 {
		n = append(n, point{x, y - 1})
	}
	if p.y < height-1 {
		n = append(n, point{x, y + 1})
	}
	return n
}

func getTreeCount(forest [][]int, points map[point]struct{}) int {
	n := 0
	for p := range points {
		// Now that we've traversed everywhre we want we only want to count the TREES (not grass).
		if forest[p.x][p.y] > 1 {
			n++
		}
	}
	return n
}

func reachableTreeCount(forest [][]int) int {
	seen := make(map[point]struct{})
	s := NewStack()
	s.Push(point{0, 0})
	for s.Len() > 0 {
		v := s.Pop()
		if _, ok := seen[v]; !ok {
			ns := neighbors(v, len(forest), len(forest[0]))
			for _, n := range ns {
				// We want to visit every accessible node (grass and not)
				if forest[n.x][n.y] > 0 {
					s.Push(n)
					seen[v] = struct{}{}
				}
			}
		}
	}
	return getTreeCount(forest, seen)
}

func canReachAllTrees(forest [][]int) bool {
	reachable := reachableTreeCount(forest)
	all := treeCount(forest)
	return reachable == all
}

func getVisits(forest [][]int)

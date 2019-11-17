package treesandgraphs

import "container/list"

type graph struct {
	nodes []*graphNode
}

type graphNode struct {
	value int
	nodes []*graphNode
}

type edge struct {
	start, end int
}

func newDirectedGraph(values []int, edges []*edge) *graph {
	nodes := make([]*graphNode, 0, len(values))
	for _, v := range values {
		nodes = append(nodes, &graphNode{v, make([]*graphNode, 0)})
	}
	for _, e := range edges {
		nodes[e.start].nodes = append(nodes[e.start].nodes, nodes[e.end])
	}
	return &graph{nodes}
}

func (g *graph) isConnected(oneIndex int, twoIndex int) bool {
	visited := make(map[*graphNode]struct{})
	queue := list.New()
	queue.PushBack(g.nodes[oneIndex])

	for queue.Len() > 0 {
		e := queue.Front()
		node := e.Value.(*graphNode)
		if node.value == g.nodes[twoIndex].value {
			return true
		}
		queue.Remove(e)
		visited[node] = struct{}{}
		for _, next := range node.nodes {
			if _, ok := visited[next]; !ok {
				queue.PushBack(next)
			}
		}
	}
	return false
}

func initialNodes(nodeValues []int) map[int]*graphNode {
	nodes := make(map[int]*graphNode)
	for _, nv := range nodeValues {
		nodes[nv] = &graphNode{nv, make([]*graphNode, 0)}
	}
	return nodes
}

func addEdges(nodes map[int]*graphNode, edges []*edge) {
	for _, e := range edges {
		node := nodes[e.start]
		node.nodes = append(node.nodes, nodes[e.end])
	}
}

func canBuild(node *graphNode, seen map[*graphNode]struct{}, dependenciesMet map[*graphNode]struct{}, order *list.List) bool {
	if _, ok := dependenciesMet[node]; ok {
		return true
	}
	if _, ok := seen[node]; ok {
		return false
	}

	seen[node] = struct{}{}

	for _, dependency := range node.nodes {
		canBuildDependency := canBuild(dependency, seen, dependenciesMet, order)
		if !canBuildDependency {
			return false
		}
	}

	dependenciesMet[node] = struct{}{}
	order.PushBack(node)
	return true
}

func getOrder(nodeValues []int, edges []*edge) ([]int, bool) {
	nodes := initialNodes(nodeValues)
	addEdges(nodes, edges)
	seen, dependenciesMet, order := make(map[*graphNode]struct{}), make(map[*graphNode]struct{}), list.New()

	for _, node := range nodes {
		if !canBuild(node, seen, dependenciesMet, order) {
			return nil, false
		}
	}
	result := make([]int, 0, order.Len())
	e := order.Front()
	for e != nil {
		result = append(result, e.Value.(*graphNode).value)
		e = e.Next()
	}
	return result, true
}

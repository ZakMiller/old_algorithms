package treesandgraphs

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDirectedGraphIsRoute(t *testing.T) {
	nodeValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	edges := []*edge{&edge{0, 1}, &edge{1, 2}}
	g := newDirectedGraph(nodeValues, edges)

	tests := []struct {
		start    int
		end      int
		expected bool
	}{
		{0, 1, true},
		{0, 2, true},
		{0, 3, false},
		{5, 6, false},
		{1, 0, false},
	}

	for _, test := range tests {
		got := g.isConnected(test.start, test.end)
		if got != test.expected {
			t.Errorf("graph.isConnected(%v, %v) error got %v expected %v", test.start, test.end, got, test.expected)
		}
	}
}

func sliceString(s []int) string {
	b := strings.Builder{}
	for _, v := range s {
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}

func edgeString(e *edge) string {
	return fmt.Sprintf("%v->%v", e.start, e.end)
}

func edgesString(s []*edge) string {
	b := strings.Builder{}
	for _, v := range s {
		b.WriteString(fmt.Sprintf("%v\n", edgeString(v)))
	}
	return b.String()
}

func TestEdgeString(t *testing.T) {
	e := &edge{1, 2}
	if edgeString(e) != "1->2" {
		t.Errorf("Edge didn't print correctly got %v expected %v", edgeString(e), "1->2")
	}
}

func TestEdgesString(t *testing.T) {
	edges := []*edge{&edge{1, 2}, &edge{4, 5}}
	s := edgesString(edges)
	expected := `1->2
4->5
`
	if s != expected {
		t.Errorf("edgesString failure got %v expected %v", s, expected)
	}
}

//func TestGraphDependencies(t *testing.T) {
//	tests := []struct {
//		nodes []int
//		edges []*edge
//		order []int
//	}{
//		{[]int{1, 2, 3, 4, 5, 6}, []*edge{&edge{1, 4}, &edge{6, 2}, &edge{2, 4}, &edge{6, 1}, &edge{4, 3}}, []int{3, 4, 2, 1, 6, 5}},
//	}
//
//	for _, test := range tests {
//		got, _ := getOrder(test.nodes, test.edges)
//		gotString := sliceString(got)
//		expectedString := sliceString(test.order)
//		edgeString := edgesString(test.edges)
//		if gotString != expectedString {
//			t.Errorf("graph dependency failure got %v expected %v for nodes %v and edges %v", gotString, expectedString, sliceString(test.nodes), edgeString)
//		}
//	}
//}

func verifyCanBuild(t *testing.T, canBuild bool, canBuildExpected bool, order *list.List, expectedOrder []int) {
	if canBuild != canBuildExpected {
		t.Errorf("canBuild failure got %v expected %v", canBuild, canBuildExpected)
	}
	if canBuild {
		if len(expectedOrder) != order.Len() {
			t.Errorf("canBuild failure wrong number of nodes")
		}
		node := order.Front()
		for i, v := range expectedOrder {
			if node.Value.(*graphNode).value != v {
				t.Errorf("canBuild failure for index %v got %v expected %v", i, node.Value.(*graphNode).value, v)
			}
			node = node.Next()
		}
	}
}
func TestCanBuild(t *testing.T) {
	node := &graphNode{0, nil}
	seen := make(map[*graphNode]struct{})
	dependenciesMet := make(map[*graphNode]struct{})
	order := list.New()
	buildable := canBuild(node, seen, dependenciesMet, order)
	verifyCanBuild(t, buildable, true, order, []int{0})
	secondNode := &graphNode{1, []*graphNode{node}}
	node.nodes = []*graphNode{secondNode}
	buildable = canBuild(node, make(map[*graphNode]struct{}), make(map[*graphNode]struct{}), list.New())
	verifyCanBuild(t, buildable, false, nil, nil)

	node1 := &graphNode{0, nil}
	node2 := &graphNode{1, []*graphNode{node1}}
	node3 := &graphNode{2, []*graphNode{node2}}
	l := list.New()
	buildable = canBuild(node3, make(map[*graphNode]struct{}), make(map[*graphNode]struct{}), l)
	verifyCanBuild(t, buildable, true, l, []int{0, 1, 2})
}

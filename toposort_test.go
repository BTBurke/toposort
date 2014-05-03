package toposort

import "testing"

type Node struct {
	name         string
	dependencies []string
}

type DAG []Node

func (d DAG) Len() int                    { return len(d) }
func (d DAG) Label(i int) string          { return d[i].name }
func (d DAG) Dependencies(i int) []string { return d[i].dependencies }
func (d DAG) Swap(i, j int)               { d[i], d[j] = d[j], d[i] }

func TestSortIndex(t *testing.T) {
	deps := []string{"test1", "test2"}
	zdeps := []string{}
	node1 := Node{name: "test", dependencies: deps}
	node2 := Node{name: "test1", dependencies: zdeps}
	node3 := Node{name: "test2", dependencies: zdeps}
	dag := DAG{node1, node2, node3}

	ind, err := SortIndex(dag)
	if err != nil {
		t.Fatalf("Received error: %s", err)
	}
	expected1 := []int{3, 2, 1}
	expected2 := []int{3, 1, 2}
	for i := 0; i < len(ind); i++ {
		if ind[i] != expected1[i] && ind[i] != expected2[i] {
			t.Fatalf("Did not sort correctly.")
		}
	}
}

func TestCyclicError(t *testing.T) {
	dep1 := []string{"test1"}
	dep2 := []string{"test2"}
	node1 := Node{name: "test1", dependencies: dep2}
	node2 := Node{name: "test2", dependencies: dep1}
	graph := DAG{node1, node2}

	_, err := SortIndex(graph)
	if err != nil {
		_, ok := err.(GraphCycleError)
		if !ok {
			t.Fatalf("Expected graph cycle error, got: %s", err)
		}
	}
}

func TestMissingDependency(t *testing.T) {
	dep1 := []string{"doesnt_exist"}
	node1 := Node{name: "test", dependencies: dep1}
	graph := DAG{node1}

	_, err := SortIndex(graph)
	if err != nil {
		_, ok := err.(MissingDependencyError)
		if !ok {
			t.Fatalf("Expected missing dependency error, got: %s", err)
		}
	}
}

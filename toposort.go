package toposort

import "fmt"

type node struct {
	dependencies []string
	tempVisit    bool
	finalVisit   bool
	sortIndex    int
}

type MissingDependencyError struct {
	msg               string
	MissingDependency string
	code              int
}

func (e MissingDependencyError) Error() string { return e.msg }

type GraphCycleError struct {
	msg       string
	CycleNode string
	code      int
}

func (e GraphCycleError) Error() string { return e.msg }

type unsortedDagMap map[string]*node

type IndexInterface interface {
	Len() int
	Label(i int) string
	Dependencies(i int) []string
}

func makeMap(n IndexInterface) unsortedDagMap {
	var out = unsortedDagMap{}
	for i := 0; i < n.Len(); i++ {
		out[n.Label(i)] = &node{dependencies: n.Dependencies(i), tempVisit: false, finalVisit: false}
	}
	return out
}

func unmarkedNodes(n unsortedDagMap) bool {
	for _, node := range n {
		if !node.finalVisit {
			return true
		}
	}
	return false
}

func addToIndex(l string, n unsortedDagMap) {
	max := 0
	for _, node := range n {
		if node.sortIndex > max {
			max = node.sortIndex
		}
	}
	n[l].sortIndex = max + 1
}

func visit(l string, graph unsortedDagMap) error {
	n := graph[l]
	if n == nil {
		err := MissingDependencyError{fmt.Sprintf("Missing dependency: %s", l), l, 404}
		return err
	}

	if n.finalVisit {
		return nil
	}
	if n.tempVisit {
		err := GraphCycleError{fmt.Sprintf("Cycle encountered at: %s", l), l, 500}
		return err
	}
	if len(n.dependencies) == 0 {
		n.finalVisit = true
		addToIndex(l, graph)
		return nil
	} else {
		n.tempVisit = true
		for i := range n.dependencies {
			err := visit(n.dependencies[i], graph)
			if err != nil {
				return err
			}
		}
		n.finalVisit = true
		addToIndex(l, graph)
		return nil
	}
}

func SortIndex(n IndexInterface) ([]int, error) {
	nodeMap := makeMap(n)

	for unmarkedNodes(nodeMap) {
		for label, node := range nodeMap {
			if !node.finalVisit {
				err := visit(label, nodeMap)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	out := make([]int, len(nodeMap))
	i := 0
	for _, node := range nodeMap {
		out[i] = node.sortIndex
		i++
	}
	return out, nil

}

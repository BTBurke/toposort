A simple Go library for topological sorting.  This was a test project to learn the fundamentals of Go.  If you have any recommendations on improving the code, please submit a pull request.

The interface to use the library is simple.  It assumes you have a collection of nodes that represent a directed acyclic graph.  Each node has a unique name and a list of dependencies.  You must define three methods on your data structure to implement the interface:

```go
type Interface interface {
	// Return the number of nodes in the graph
	Len()               int

	// Return the unique label/name of the node at index i
    Label(i int)        string

    // Return a list of labels on which node[i] depends
    Dependencies(i int) []string 
}
```

The function `SortIndex(n Interface)` will return a `[]int` which represents the execution order that satisfies all dependencies.  You can in turn use this index with Go's sort method to sort the graph in place, if desired.

Two types of possible errors are returned.  You can use type inference to take additional actions as necessary.  See the tests for ideas.

Error | Returned when
----- | -------------
MissingDependencyError | Returned if a dependency doesn't exist in the graph
CyclicGraphError | Returned if the graph contains a cycle.  

Note:
- The returned order is not guaranteed to be unique.
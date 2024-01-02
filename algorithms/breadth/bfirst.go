package breadth

import "fmt"

// Breadth First Search - big O is O(n2)
func Breadthfirst(target string, graph map[int][]string) (distance int, destination int) {
	fmt.Println("--------------------------------")
	fmt.Println("Breadth First Search")
	fmt.Println("Find Shortest path")
	fmt.Println("--------------------------------")
	short_stop := 0
	var route int
	for ix, _ := range graph {
		stops := 0
		for i := 0; i < len(graph[ix]); i++ {
			stops++
			if graph[ix][i] == target {
				if short_stop == 0 {
					short_stop = stops
					route = ix
				} else if short_stop > stops {
					short_stop = stops
					route = ix
				}
			}
		}
	}
	return short_stop, route
}

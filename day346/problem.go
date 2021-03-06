package day346

// Flights represents all the flights and their prices.
type Flights map[string]map[string]int

// CheapestItinerary returns the cheapest itinerary that respects the
// maxConnections limit.
// Effectively, DFS for all possible answers that don't exceed the maximum
// number of connections.
// Runs in DFS performance. O(N+M)
func CheapestItinerary(flights Flights, origin, destination string, maxConnections int) ([]string, int) {
	visited := make(map[string]struct{})
	return cheapestItinerary(flights, origin, destination, maxConnections, visited)
}

func cheapestItinerary(flights Flights, origin, destination string, connections int,
	visited map[string]struct{}) ([]string, int) {
	if connections < 0 {
		return nil, 0
	}
	if origin == destination {
		return []string{destination}, 0
	}
	cheapest := int(^uint(0) >> 1)
	var itinerary []string
	for next := range flights[origin] {
		if _, seen := visited[next]; seen {
			continue
		}
		visited[next] = struct{}{}
		segments, price := cheapestItinerary(flights, next, destination, connections-1, visited)
		if len(segments) != 0 {
			if total := flights[origin][next] + price; total < cheapest {
				cheapest = total
				itinerary = append([]string{origin}, segments...)
			}
		}
		delete(visited, next)
	}
	return itinerary, cheapest
}

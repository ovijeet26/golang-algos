package graph

import "math"

// Leet code link -> https://leetcode.com/problems/network-delay-time/
// You are given a network of n nodes, labeled from 1 to n. You are also given times,
// a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node,
// vi is the target node, and wi is the time it takes for a signal to travel from source to target.
func networkDelayTimeBf(times [][]int, n int, k int) int {

	// Bellman-ford uses dynamic programming concepts under the hood.
	// It is used to find the minimum distance between nodes.
	// It can also work with negative weights.
	// It does not work ONLY when there is a negative weighted cycle.

	distances := make([]int, 0)

	infinity := int(math.Inf(1))

	// Apend the final distances array with all negative weights.
	for i := 0; i < n; i++ {

		distances = append(distances, infinity)
	}

	// Mark the source node distance as 0, as it's the starting index.
	distances[k-1] = 0

	// Bellman-ford required us to iterate over the graph n-1 times.
	// We can skip looping if we detect no change in the loop.

	for i := 0; i < n-1; i++ {
		changed := false
		// Iteraing over all the edges
		for j := 0; j < len(times); j++ {

			source := times[j][0] - 1
			target := times[j][1] - 1
			weight := times[j][2]

			if distances[source]+weight < distances[target] {

				distances[target] = distances[source] + weight
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	max := 0

	for _, distance := range distances {
		if distance > max {
			max = distance
		}
	}

	if max == infinity {

		return -1
	}
	return max
}

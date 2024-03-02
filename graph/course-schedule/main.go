package main

func canFinish(numCourses int, prerequisites [][]int) bool {

	// Key Idea: We form the graph of each course using adjacency list and then
	// perform a DFS to detect cycles. For this, we keep track of two maps
	// `visited` and `visting` to detect the vistited vertex and visiting vertex in each iteration

	graph := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		course, prereqs := prerequisite[0], prerequisite[1]
		graph[course] = append(graph[course], prereqs)
	}

	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	// Define a DFS function to detect cycles
	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {

		// Already detected in current recursion => cycle detected
		if visiting[course] {
			return true
		}

		// Already checked and detected no cycles
		if visited[course] {
			return false
		}

		// Mark it the in the current visiting map
		visiting[course] = true
		for _, prereq := range graph[course] {
			if hasCycle(prereq) {
				return true
			}
		}

		// Mark visiting back to false after finished traversing the graph for that course
		visiting[course] = false

		// Make the course as visited
		visited[course] = true

		return false
	}

	// Check starts from here
	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return false
		}
	}

	return true
}

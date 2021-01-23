package maze

import "math/rand"

func (maze *DummyMaze) Loopify() {
	// Loop through all cells in the maze
	for v := 0; v < (maze.width * maze.height); v++ {
		// If there is only one path from the current cell
		if maze.grid.Degree(v) == 1 {
			var posibleExpansions = make([]int, 0, 4)
			var cellToExtend int

			// If the cell is *not* on the top edge
			if v > maze.width {
				// And there isn't an edge already
				if !maze.grid.Edge(v, v - maze.width) {
					// Add the cell as a posibility for expansion
					posibleExpansions = append(posibleExpansions, v - maze.width)
				}
			}
		
			// If the cell is *not* on the right hand edge
			if (v % maze.width) != (maze.width - 1) {
				// And there isn't an edge already
				if !maze.grid.Edge(v, v + 1) {
					// Add the cell as a posibility for expansion
					posibleExpansions = append(posibleExpansions, v + 1)
				}
			}
		
			// If the cell is *not* on the bottom edge
			if (v < (maze.width * (maze.height - 1)))  {
				// And there isn't an edge already
				if !maze.grid.Edge(v, v + maze.width) {
					// Add the cell as a posibility for expansion
					posibleExpansions = append(posibleExpansions, v + maze.width)
				}
			}
		
			// If the cell is *not* on the left edge
			if (v % maze.width) != 0 {
				// And there isn't an edge already
				if !maze.grid.Edge(v, v - 1) {
					// Add the cell as a posibility for expansion
					posibleExpansions = append(posibleExpansions, v - 1)
				}
			}

			if len(posibleExpansions) > 1 {
				cellToExtend = posibleExpansions[rand.Intn(len(posibleExpansions))]
			} else {
				cellToExtend = posibleExpansions[0]
			}

			maze.grid.AddBoth(v, cellToExtend)
		}
	}
}
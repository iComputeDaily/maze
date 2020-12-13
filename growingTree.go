package maze

import "math/rand"
import "github.com/yourbasic/graph"

// Defines one maze
type GTreeMaze struct {
	// Generic atributes/functions
	DummyMaze
	
	// Algorithum specific atributes/functions
	visited []bool // Tells weather or not a cell is visited
	currentList []int // Holds a list of the currently in use cell ids(ordered by age)
}

// Initalizes data structures
func (maze *GTreeMaze) init(width int, height int) {
	// Sets the maze size
	maze.width = width
	maze.height = height
	
	// Creates a new grid with the correct number of vertecies
	maze.grid = graph.New(maze.width * maze.height)
	
	// Alocate the correct ammount of memory to the lists
	maze.visited = make([]bool, maze.width * maze.height)
	maze.currentList = make([]int, 0, maze.width * (maze.height / 2)) // BUG(iComputeDaily): optimize this value later
	
	// Set all cells to unvisited
	for i := 0; i < (maze.width * maze.height); i++ {
		maze.visited[i] = false
	}
	
	return
}

func (maze *GTreeMaze) Generate(width int, height int) {
	// Initalize data structure
	maze.init(width, height)
	
	// Add the first cell to the list
	randCellId := rand.Intn(maze.width * maze.height)
	maze.currentList = append(maze.currentList, randCellId)
	
	// Set the cell to visited
	maze.visited[randCellId] = true
	
	// Generate the maze
	for len(maze.currentList) != 0 {
		// Holds the id of the cell curently being considered
		var workingCell int
		// Holds the ids of cells that the maze could be extended to
		var posibleExpansions = make([]int, 0, 3) // BUG(iComputeDaily): optimize this value later
		// Holds the id of the cell in the currentList(just to speed things up by not having to search for it)
		var currentListCellId int
		// Holds the cell selected to extend the maze to
		var cellToExtend int
		
		// Select the newest cell from the working cell list
		currentListCellId = len(maze.currentList) - 1
		workingCell = maze.currentList[currentListCellId]

		
		// If the cell is *not* on the top edge
		if workingCell > maze.width {
			// If the cell above is unvisited
			if maze.visited[workingCell - maze.width] == false {
				// Add the cell above to the list of posible expansions
				posibleExpansions = append(posibleExpansions, workingCell - maze.width)
			}
		}
		
		// If the cell is *not* on the right hand edge
		if (workingCell % maze.width) != (maze.width - 1) {
			// If the cell to the right is unvisited
			if maze.visited[workingCell + 1] == false {
				// Add the cell to the right to the list of posible expansions
				posibleExpansions = append(posibleExpansions, workingCell + 1)
			}
		}
		
		// If the cell is *not* on the bottom edge
		if (workingCell < (maze.width * (maze.height - 1)))  {
			// If the cell bellow is unvisited
			if maze.visited[workingCell + maze.width] == false {
				// Add the cell bellow to the list of posible expansions
				posibleExpansions = append(posibleExpansions, workingCell + maze.width)
			}
		}
		
		// If the cell is *not* on the left edge
		if (workingCell % maze.width) != 0 {
			// If the cell to the left is unvisited
			if maze.visited[workingCell - 1] == false {
				// Add the cell to the left to the list of posible expansions
				posibleExpansions = append(posibleExpansions, workingCell - 1)
			}
		}
		
		// Remove the working cell if no unvisited neibors were found
		if len(posibleExpansions) == 0 {
			// Shift a[i+1:] left one index.
			copy(maze.currentList[currentListCellId:], maze.currentList[currentListCellId + 1:])
			// Erase last element (write zero value).
			maze.currentList[len(maze.currentList) - 1] = 0
			// Truncate slice.
			maze.currentList = maze.currentList[:len(maze.currentList) - 1]
		
		// Otherwise extend the maze to a random neibor
		} else {
			if len(posibleExpansions) > 1 {
				cellToExtend = posibleExpansions[rand.Intn(len(posibleExpansions))]
			} else {
				cellToExtend = posibleExpansions[0]
			}
			// And extend the maze
			// Add the cell to the list
			maze.currentList = append(maze.currentList, cellToExtend)
			// Mark it as visited
			maze.visited[cellToExtend] = true
			// And carve a path
			maze.grid.Add(workingCell, cellToExtend)
		}
	}
	
	// Clear uneeded memory
	maze.visited = nil
	maze.currentList = nil
	
	return
}

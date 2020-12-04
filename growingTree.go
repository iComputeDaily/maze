package main

import "fmt"
import "math/rand"
import "github.com/yourbasic/graph"

// Defines one maze
type gTreeMaze struct {
	grid *graph.Mutable // Holds a graph to represent borders of all the cells
	visited []bool // Tells weather or not a cell is visited
	currentList []int // Holds a list of the currently in use cell ids(ordered by age)
	
	height int
	width int
}

// Initalizes data structures
func (maze *gTreeMaze) init() error {
	// Sets the maze size
	maze.width = 15
	maze.height = 15
	
	// Creates a new grid with the correct number of vertecies
	maze.grid = graph.New(maze.width * maze.height)
	
	// Alocate the correct ammount of memory to the lists
	maze.visited = make([]bool, maze.width * maze.height)
	maze.currentList = make([]int, 0, maze.width * (maze.height / 2)) // BUG(iComputeDaily): optimize this value later
	
	// Set all cells to unvisited
	for i := 0; i < (maze.width * maze.height); i++ {
		maze.visited[i] = false
	}
	
	return nil
}

func (maze *gTreeMaze) generate()  error {
	// Initalize data structure
	err := maze.init()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
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
	
	return nil
}

// Returns the maze as a string of ascii art
func (maze *gTreeMaze) stringify() (string, error) {
	var stringyMaze string
	
	// Draws the edges on the top of the maze
	for v := 0; v < (maze.width * 2) + 1; v++ {
		stringyMaze = stringyMaze + "_"
	}
	// Go to next line for next row
	stringyMaze = stringyMaze + "\n"
	
	// Loop through all rows exept for last
	for y := 0; y < (maze.height - 1); y++ {
		// Print the left border
		stringyMaze = stringyMaze + "|"
		
		// Loop through all cells in current row
		for v := y * maze.width; v < ((y + 1) * maze.width); v++ {
			var vertical bool
			var horizontal bool
			
			// If the current cell and the cell below connect
			if maze.grid.Edge(v, v + maze.width) || maze.grid.Edge(v + maze.width, v) == true {
				vertical = true
			} else {
				vertical = false
			}
			
			// If the current cell connects with the cell to the right
			if maze.grid.Edge(v, v + 1) || maze.grid.Edge(v + 1, v) == true {
				horizontal = true
			} else {
				horizontal = false
			}
			
			if horizontal && vertical == true {
				// If the cell to the right and the cell below are connected
				if maze.grid.Edge(v + 1, (v + 1) + maze.width) || maze.grid.Edge((v + 1) + maze.width, v + 1) {
					stringyMaze = stringyMaze + "  "
				} else {
					stringyMaze = stringyMaze + " _"
				}
			} else if horizontal == true {
				stringyMaze = stringyMaze + "__"
			} else if vertical == true {
				stringyMaze = stringyMaze + " |"
			} else {
				stringyMaze = stringyMaze + "_|"
			}
		}
		// Go to next line for next row
		stringyMaze = stringyMaze + "\n"
	}
	
	// For the last row
	// Print the left border
	stringyMaze = stringyMaze + "|"
	
	// Loop through all cells in current row
	for v := (maze.height - 1) * maze.width; v < (maze.height * maze.width); v++ {
		// If the current cell connects with the cell to the right
		if maze.grid.Edge(v, v + 1) || maze.grid.Edge(v + 1, v) {
			stringyMaze = stringyMaze + "__"
		} else {
			stringyMaze = stringyMaze + "_|"
		}
	}
	return stringyMaze, nil
}

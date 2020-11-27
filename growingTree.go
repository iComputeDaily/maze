package main

import "fmt"
import "math/rand"
import "github.com/yourbasic/graph"

// Defines one maze
type gTreeMaze struct {
	data struct { // Holds the cell ids and edge list
		grid *graph.Mutable // Holds a graph to represent borders of all the cells
		unvisited []int // Contains all cells initaly and cells are removed when they are visited
		currentList []int // Holds a list of the currently in use cells(ordered by age)
	}
	metadata struct {
		size struct {
			x int
			y int
		}
	}
}

// Initalizes data structures
func (maze *gTreeMaze) init() error {
	// Sets the maze size
	maze.metadata.size.x = 15
	maze.metadata.size.y = 15
	
	// Creates a new grid with the correct number of vertecies
	maze.data.grid = graph.New(maze.metadata.size.x * maze.metadata.size.y)
	
	// Alocate the correct ammount of memory to the list
	maze.data.unvisited = make([]int, maze.metadata.size.x * maze.metadata.size.y)
	
	// Add all cells to the list of unvisited cells
	for i := 0; i < (maze.metadata.size.x * maze.metadata.size.y); i++ {
		maze.data.unvisited[i] = i
	}
	
	// Debugging
	fmt.Println("edges:", maze.data.unvisited)
	
	return nil
}

func (maze *gTreeMaze) generate()  error {
	// Initalize data structure
	err := maze.init()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	// Add the first cell to the list
	maze.data.currentList = append(maze.data.currentList, rand.Intn((maze.metadata.size.x * maze.metadata.size.y) - 1))
	
	// Generate the maze
	for len(maze.data.currentList) != 0 {
		// Select a random cell from the working cell list
		workingCell := maze.data.currentList[rand.Intn(len(maze.data.currentList) - 1)]
		
		// If the cell is on the right hand edge
		if workingCell == (maze.metadata.x - 1) {
			
		}
	}
	
	return nil
}

// Returns the maze as a string of ascii art
func (maze *gTreeMaze) stringify() (string, error) {
	var stringyMaze string
	
	// Draws the edges on the top of the maze
	for v := 0; v < (maze.metadata.size.x * 2) + 1; v++ {
		stringyMaze = stringyMaze + "_"
	}
	// Go to next line for next row
	stringyMaze = stringyMaze + "\n"
	
	// Loop through all rows exept for last
	for y := 0; y < (maze.metadata.size.y - 1); y++ {
		// Print the left border
		stringyMaze = stringyMaze + "|"
		
		// Loop through all cells in current row
		for v := y * maze.metadata.size.x; v < ((y + 1) * maze.metadata.size.x); v++ {
			// If the current cell and the cell below connect
			vertical := maze.data.grid.Edge(v, v + maze.metadata.size.x)
			
			// If the current cell connects with the cell to the right
			horizontal := maze.data.grid.Edge(v, v + 1)
			
			if horizontal && vertical == true {
				// If the cell to the right and the cell below are connected
				if maze.data.grid.Edge(v + 1, (v + 1) + maze.metadata.size.x) {
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
	for v := (maze.metadata.size.y - 1) * maze.metadata.size.x; v < (maze.metadata.size.y * maze.metadata.size.x); v++ {
		// If the current cell connects with the cell to the right
		if maze.data.grid.Edge(v, v + 1) {
			stringyMaze = stringyMaze + "__"
		} else {
			stringyMaze = stringyMaze + "_|"
		}
	}
	return stringyMaze, nil
}

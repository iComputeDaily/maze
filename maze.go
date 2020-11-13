package main

import "fmt"
import "github.com/yourbasic/graph"

// Defines one maze
type maze struct {
	data struct { // Holds the cell ids and edge list
		grid *graph.Mutable // Holds a graph to represent borders of all the cells
		sets []uint32 // Represents cell ids
		edgesList []edge // A list of all the edges
	}
	metadata struct {
		size struct {
			x int
			y int
		}
	}
}

// Defines an edge as the border between 2 cells
type edge struct {
	cell1 uint32
	cell2 uint32
}

// Initalizes a grid with all the vertecies and edges
// BUG(iComputeDaily): later remove edge initalization
func (maze *maze) init() error {
	// Sets the maze size
	maze.metadata.size.x = 5
	maze.metadata.size.y = 5
	
	// Creates a new grid with the correct number of vertecies
	maze.data.grid = graph.New(maze.metadata.size.x * maze.metadata.size.y)
	
	// Creates all horizontal edges
	for y := 0; y < maze.metadata.size.y; y++ { // Loops over y values; top to bottom
		for x := 0; x < maze.metadata.size.x - 1; x++ { // Loops over x values; left to right
			// Connects the current cell to the cell to it's right
			maze.data.grid.Add(x + (y * maze.metadata.size.y), x + (y * maze.metadata.size.x) + 1)
		}
	}
	
	// Creates all vertical edges
	for x := 0; x < maze.metadata.size.x; x++ { // Loops over x values; left to right
		for y := 0; y < maze.metadata.size.y - 1; y++ { // Loops over y values; top to bottom
			// Connects the current cell to the cell bellow it
			maze.data.grid.Add(x + (y * maze.metadata.size.x), x + ((y + 1) * maze.metadata.size.y))
		}
	}
	
	// debuging perposes
	fmt.Println("grid is: ", maze.data.grid.String())
	
	return nil
}

// Randomly genrates a maze
func (maze *maze) generate() error {
	return nil
}

// Returns the maze as a string of ascii art
func (maze *maze) stringify() (string, error) {
	var stringyMaze string
	
	// Draws the edges on the top of the maze
	for v := 0; v < (maze.metadata.size.x); v++ {
		if maze.data.grid.Edge(v, v + 1) == true {
			stringyMaze = stringyMaze + "_"
		} else {
			stringyMaze = stringyMaze + " "
		}
	}
	
	return stringyMaze, nil
}

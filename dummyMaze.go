package main

import "strings"
import "github.com/yourbasic/graph"

// Represents a maze with no generation algoritum
type dummyMaze struct {
	grid *graph.Mutable // Represents the grid of squares
	
	// You can probobly guess
	height int
	width int
}

// Returns the maze as a string of ascii art
func (maze *dummyMaze) stringify() (string, error) {
	// Create a string builder for the maze and prealcoate memory to avoid memory alocations
	var stringyMaze strings.Builder
	stringyMaze.Grow((maze.height + 1) * ((maze.width * 2) + 1))
	
	// Draws the edges on the top of the maze
	for v := 0; v < (maze.width * 2) + 1; v++ {
		_, _ = stringyMaze.WriteString("_")
	}
	// Go to next line for next row
	_, _ = stringyMaze.WriteString("\n")
	
	// Loop through all rows exept for last
	for y := 0; y < (maze.height - 1); y++ {
		// Print the left border
		_, _ = stringyMaze.WriteString("|")
		
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
					_, _ = stringyMaze.WriteString("  ")
				} else {
					_, _ = stringyMaze.WriteString(" _")
				}
			} else if horizontal == true {
				_, _ = stringyMaze.WriteString("__")
			} else if vertical == true {
				_, _ = stringyMaze.WriteString(" |")
			} else {
				_, _ = stringyMaze.WriteString("_|")
			}
		}
		// Go to next line for next row
		_, _ = stringyMaze.WriteString("\n")
	}
	
	// For the last row
	// Print the left border
	_, _ = stringyMaze.WriteString("|")
	
	// Loop through all cells in current row
	for v := (maze.height - 1) * maze.width; v < (maze.height * maze.width); v++ {
		// If the current cell connects with the cell to the right
		if maze.grid.Edge(v, v + 1) || maze.grid.Edge(v + 1, v) {
			_, _ = stringyMaze.WriteString("__")
		} else {
			_, _ = stringyMaze.WriteString("_|")
		}
	}
	return stringyMaze.String(), nil
}

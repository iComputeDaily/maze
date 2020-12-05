package main

import "strings"
import "fmt"
import "math/rand"
import "github.com/yourbasic/graph"

// Defines one maze
type kruskalMaze struct {
	grid *graph.Mutable // Holds a graph to represent borders of all the cells
	sets []int // Represents cell ids
	edgesList []edge // A list of all the edges
	
	height int
	width int
}

// Defines an edge as the border between 2 cells
type edge struct {
	cell1 int
	cell2 int
}

// Initalizes a grid with all the vertecies and edges
func (maze *kruskalMaze) init() error {
	// Sets the maze size
	maze.width = 40
	maze.height = 5000
	
	// Creates a new grid with the correct number of vertecies
	maze.grid = graph.New(maze.width * maze.height)
	
	// Alocate memory to the slices
	maze.edgesList = make([]edge, 0, ((maze.width - 1) * maze.height) +
	(maze.width * (maze.height - 1)))
	maze.sets = make([]int, maze.width * maze.height)
	
	// Give all cells a uniqe cell id
	for i := 0; i < (maze.width * maze.height); i++ {
		maze.sets[i] = i
	}
	
	// Add all edges to the list
	// Adds all horizontal edges
	for y := 0; y < maze.height; y++ { // Loops over y values; top to bottom
		for x := 0; x < maze.width - 1; x++ { // Loops over x values; left to right
			// Adds the edge between the current cell and the cell to it's right
			maze.edgesList = append(maze.edgesList, edge{
				cell1: x + (y * maze.width),
				cell2: x + ((y * maze.width) + 1)})
		}
	}
	
	// Adds all vertical edges
	for x := 0; x < maze.width; x++ { // Loops over x values; left to right
		for y := 0; y < maze.height - 1; y++ { // Loops over y values; top to bottom
			// Adds the edge between the current cell and the cell bellow it
			maze.edgesList = append(maze.edgesList, edge{
				cell1: x + (y * maze.width),
				cell2: x + ((y + 1) * maze.width)})
		}
	}
	
	// Randomizes the order of the slice of edges
	rand.Shuffle(len(maze.edgesList), func(i, j int) {
		maze.edgesList[i], maze.edgesList[j] =
		maze.edgesList[j], maze.edgesList[i] })
	
	return nil
}

// Randomly genrates a maze
func (maze *kruskalMaze) generate() error {
	// Initalize the grid and other data structures
	err := maze.init()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	// Generate the maze
	for e := 0; e < len(maze.edgesList); e++ { // For each edge(from our randomized list)
		// Check if the cells on either side of this edge are not of the same set
		if maze.sets[maze.edgesList[e].cell1] !=
			maze.sets[maze.edgesList[e].cell2] {
				
				// If not carve a path
				maze.grid.Add(maze.edgesList[e].cell1, maze.edgesList[e].cell2)
				
				// Store the cell ids in variables to avoid canging the cell id that is matched for during the loop
				cell1, cell2 := maze.sets[maze.edgesList[e].cell1], maze.sets[maze.edgesList[e].cell2]
				
				// And join the sets
				for c := 0; c < len(maze.sets); c++ { // For every cell in the sets list
					// If its index is the same as cell one change it to the index of cell 2
					if maze.sets[c] == cell1 {
						maze.sets[c] = cell2
					}
				}
			}
	}
	
	return nil
}

// Returns the maze as a string of ascii art
func (maze *kruskalMaze) stringify() (string, error) {
	var stringyMaze strings.Builder
	
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
			// If the current cell and the cell below connect
			vertical := maze.grid.Edge(v, v + maze.width)
			
			// If the current cell connects with the cell to the right
			horizontal := maze.grid.Edge(v, v + 1)
			
			if horizontal && vertical == true {
				// If the cell to the right and the cell below are connected
				if maze.grid.Edge(v + 1, (v + 1) + maze.width) {
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
		if maze.grid.Edge(v, v + 1) {
			stringyMaze = stringyMaze + "__"
		} else {
			stringyMaze = stringyMaze + "_|"
		}
	}
	return stringyMaze, nil
}

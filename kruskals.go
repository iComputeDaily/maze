package main

import "fmt"
import "math/rand"
import "github.com/yourbasic/graph"

// Defines one maze
type kruskalMaze struct {
	data struct { // Holds the cell ids and edge list
		grid *graph.Mutable // Holds a graph to represent borders of all the cells
		sets []int // Represents cell ids
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
	cell1 int
	cell2 int
}

// Initalizes a grid with all the vertecies and edges
func (maze *kruskalMaze) init() error {
	// Sets the maze size
	maze.metadata.size.x = 10
	maze.metadata.size.y = 10
	
	// Creates a new grid with the correct number of vertecies
	maze.data.grid = graph.New(maze.metadata.size.x * maze.metadata.size.y)
	
	// debuging perposes
	fmt.Println("grid is: ", maze.data.grid.String())
	
	// Alocate memory to the slices
	maze.data.edgesList = make([]edge, 0, ((maze.metadata.size.x - 1) * maze.metadata.size.y) +
	(maze.metadata.size.x * (maze.metadata.size.y - 1)))
	maze.data.sets = make([]int, maze.metadata.size.x * maze.metadata.size.y)
	
	// Give all cells a uniqe cell id
	for i := 0; i < (maze.metadata.size.x * maze.metadata.size.y); i++ {
		maze.data.sets[i] = i
	}
	
	// Add all edges to the list
	// Adds all horizontal edges
	for y := 0; y < maze.metadata.size.y; y++ { // Loops over y values; top to bottom
		for x := 0; x < maze.metadata.size.x - 1; x++ { // Loops over x values; left to right
			// Adds the edge between the current cell and the cell to it's right
			maze.data.edgesList = append(maze.data.edgesList, edge{
				cell1: x + (y * maze.metadata.size.x),
				cell2: x + ((y * maze.metadata.size.x) + 1)})
		}
	}
	
	// Adds all vertical edges
	for x := 0; x < maze.metadata.size.x; x++ { // Loops over x values; left to right
		for y := 0; y < maze.metadata.size.y - 1; y++ { // Loops over y values; top to bottom
			// Adds the edge between the current cell and the cell bellow it
			maze.data.edgesList = append(maze.data.edgesList, edge{
				cell1: x + (y * maze.metadata.size.x),
				cell2: x + ((y + 1) * maze.metadata.size.x)})
		}
	}
	
	// Randomizes the order of the slice of edges
	rand.Shuffle(len(maze.data.edgesList), func(i, j int) {
		maze.data.edgesList[i], maze.data.edgesList[j] =
		maze.data.edgesList[j], maze.data.edgesList[i] })
	
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
	for e := 0; e < len(maze.data.edgesList); e++ { // For each edge(from our randomized list)
		// Check if the cells on either side of this edge are not of the same set
		if maze.data.sets[maze.data.edgesList[e].cell1] !=
			maze.data.sets[maze.data.edgesList[e].cell2] {
				
				// If not carve a path
				maze.data.grid.Add(maze.data.edgesList[e].cell1, maze.data.edgesList[e].cell2)
				
				// Store the cell ids in variables to avoid canging the cell id that is matched for during the loop
				cell1ID, cell2ID := maze.data.sets[maze.data.edgesList[e].cell1], maze.data.sets[maze.data.edgesList[e].cell2]
				
				// And join the sets
				for c := 0; c < len(maze.data.sets); c++ { // For every cell in the sets list
					// If its index is the same as cell one change it to the index of cell 2
					if maze.data.sets[c] == cell1ID {
						maze.data.sets[c] = cell2ID
					}
				}
			}
	}
	
	return nil
}

// Returns the maze as a string of ascii art
func (maze *kruskalMaze) stringify() (string, error) {
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

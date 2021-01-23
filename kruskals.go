package maze

import "github.com/theodesp/unionfind"
import "math/rand"
import "github.com/yourbasic/graph"

// Defines one maze
type KruskalMaze struct {
	// Generic atributes/functions
	DummyMaze
	
	// Algorithum specific atributes/functions
	sets *unionfind.UnionFind // Represents witch cells are connected to eachother
	edgesList []edge // A list of all the edges
}

// Defines an edge as the border between 2 cells
type edge struct {
	cell1 int
	cell2 int
}

// Initalizes a grid with all the vertecies and edges
func (maze *KruskalMaze) init(width int, height int) {
	// Sets the maze size
	maze.width = width
	maze.height = height
	
	// Creates a new grid with the correct number of vertecies
	maze.grid = graph.New(maze.width * maze.height)
	
	// Alocate memory to the edges
	maze.edgesList = make([]edge, 0, ((maze.width - 1) * maze.height) +
	(maze.width * (maze.height - 1)))
	
	// Alocate the sets data structure
	maze.sets = unionfind.New(maze.width * maze.height)
	
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
	
	return
}

// Randomly genrates a maze
func (maze *KruskalMaze) Generate(width int, height int) {
	// Initalize the grid and other data structures
	maze.init(width, height)
	
	// Generate the maze
	for e := 0; e < len(maze.edgesList); e++ { // For each edge(from our randomized list)
		// Check if the cells on either side of this edge are not of the same set
		if maze.sets.Connected(maze.edgesList[e].cell1, maze.edgesList[e].cell2) == false {
				
				// If not carve a path
				maze.grid.AddBoth(maze.edgesList[e].cell1, maze.edgesList[e].cell2)
				
				// And join the sets
				maze.sets.Union(maze.edgesList[e].cell1, maze.edgesList[e].cell2)
			}
	}
	
	// Free uneeded memory
	maze.sets = nil
	maze.edgesList = nil
	
	return
}

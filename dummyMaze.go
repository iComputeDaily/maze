package maze

import "strings"
import "github.com/yourbasic/graph"

// Repersents any maze
type Maze interface {
	Generate(width int, height int)
	Stringify() string
	
	MoveLeft() bool
	MoveUp() bool
	MoveDown() bool
	MoveRight() bool
}

// Represents a maze with no generation algoritum
type DummyMaze struct {
	grid *graph.Mutable // Represents the grid of squares
	
	// You can probobly guess
	height int
	width int
	
	playerX int
	playerY int
}
// Move the player left, returns weather or not the move was sucsessfull
func (maze *DummyMaze) MoveLeft() bool {
	playerCell := maze.playerX + (maze.playerY * maze.width)
	
	// If there is and edge between where the player is and where they want to move then move them, and return true
	if maze.grid.Edge(playerCell, playerCell - 1) || maze.grid.Edge(playerCell - 1, playerCell) {
		maze.playerX -= 1
		return true
	}
	// Otherwise return false
	return false
}
// Move the player up, returns weather or not the move was sucsessfull
func (maze *DummyMaze) MoveUp() bool {
	playerCell := maze.playerX + (maze.playerY * maze.width)
	
	// If there is and edge between where the player is and where they want to move then move them, and return true
	if maze.grid.Edge(playerCell, playerCell - maze.width) || maze.grid.Edge(playerCell - maze.width, playerCell) {
		maze.playerY -= 1
		return true
	}
	// Otherwise return false
	return false
}
// Move the player down, returns weather or not the move was sucsessfull
func (maze *DummyMaze) MoveDown() bool {
	playerCell := maze.playerX + (maze.playerY * maze.width)
	
	// If there is and edge between where the player is and where they want to move then move them, and return true
	if maze.grid.Edge(playerCell, playerCell + maze.width) || maze.grid.Edge(playerCell + maze.width, playerCell) {
		maze.playerY += 1
		return true
	}
	// Otherwise return false
	return false
}
// Move the player right, returns weather or not the move was sucsessfull
func (maze *DummyMaze) MoveRight() bool {
	playerCell := maze.playerX + (maze.playerY * maze.width)
	
	// If there is and edge between where the player is and where they want to move then move them, and return true
	if maze.grid.Edge(playerCell, playerCell + 1) || maze.grid.Edge(playerCell + 1, playerCell) {
		maze.playerX += 1
		return true
	}
	// Otherwise return false
	return false
}

// Prints the current cell
func (maze *DummyMaze) printCurrentCell(stringyMaze *strings.Builder, v int, y int) {
	var horizontal bool
	var vertical bool
	
	// If the current cell and the cell below connect
	if maze.grid.Edge(v, v + maze.width) || maze.grid.Edge(v + maze.width, v) == true {
		vertical = true } else { vertical = false }
	
	// If the current cell connects with the cell to the right
	if maze.grid.Edge(v, v + 1) || maze.grid.Edge(v + 1, v) == true {
		horizontal = true } else { horizontal = false }
	
	// If the player is not on this square
	if (maze.playerY != y) || (maze.playerX != v - (y * maze.width)) {
		switch {
			case horizontal && vertical == true:
				// If the cell to the right and the cell below the cell to the right are connected
				if maze.grid.Edge(v + 1, (v + 1) + maze.width) || maze.grid.Edge((v + 1) + maze.width, v + 1) {
					_, _ = stringyMaze.WriteString("  ")
				} else {
					_, _ = stringyMaze.WriteString(" _")
				}
			case horizontal == true:
				_, _ = stringyMaze.WriteString("__")
			case vertical == true:
				_, _ = stringyMaze.WriteString(" |")
			default:
				_, _ = stringyMaze.WriteString("_|")
		}
	} else { // If the player is on this square
		switch {
			case horizontal && vertical == true:
				// If the cell to the right and the cell below the cell to the right are connected
				if maze.grid.Edge(v + 1, (v + 1) + maze.width) || maze.grid.Edge((v + 1) + maze.width, v + 1) {
					_, _ = stringyMaze.WriteString("O ")
				} else {
					_, _ = stringyMaze.WriteString("O_")
				}
			case horizontal == true:
				_, _ = stringyMaze.WriteString("O_")
			case vertical == true:
				_, _ = stringyMaze.WriteString("O|")
			default:
				_, _ = stringyMaze.WriteString("O|")
		}
	}
}

// Returns the maze as a string of ascii art
func (maze *DummyMaze) Stringify() string {
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
			maze.printCurrentCell(&stringyMaze, v, y)
		}
		// Go to next line for next row
		_, _ = stringyMaze.WriteString("\n")
	}
	
	// For the last row
	// Print the left border
	_, _ = stringyMaze.WriteString("|")
	
	// Loop through all cells in last row
	for v := (maze.height - 1) * maze.width; v < (maze.height * maze.width); v++ {
		// If the player is not on the current square
		if (maze.playerY != (maze.height - 1)) || (maze.playerX != v - ((maze.height - 1) * maze.width)) {
			// If the current cell connects with the cell to the right
			if maze.grid.Edge(v, v + 1) || maze.grid.Edge(v + 1, v) {
				_, _ = stringyMaze.WriteString("__")
			} else {
				_, _ = stringyMaze.WriteString("_|")
			}
		} else { // If the player is on the current square
			// If the current cell connects with the cell to the right
			if maze.grid.Edge(v, v + 1) || maze.grid.Edge(v + 1, v) {
				_, _ = stringyMaze.WriteString("O_")
			} else {
				_, _ = stringyMaze.WriteString("O|")
			}
		}
	}
	return stringyMaze.String()
}

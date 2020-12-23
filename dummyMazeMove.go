package maze

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

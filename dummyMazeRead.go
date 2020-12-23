package maze

import "strings"
import "fmt"

// Reads one row of the maze from string and returns weather or not reading was sucsessfull
func (maze *DummyMaze) readRow(mazeReader *strings.Reader) bool {
	// Read first pipe
	currentChar, err := mazeReader.ReadByte()
	// If we reached EOF return false
	if (err != nil) | (currentChar != '|') { return false }
	
	if mazeReader.Len()
	
	for x := 0; x < maze.width; x++ {
		// Read charachter
		currentChar, err := mazeReader.ReadByte()
		// Make shure we haven't reached the end of the line
		if (err != nil) | (currentChar == '\n') {
			break
		} else {
			// Add one to maze height
			maze.height += 1
		}
	}
}

// Reads the maze from a string including finding the start
func (maze *DummyMaze) ReadMaze(inMaze string) {
	var realMaze string // Holds the actual maze
	
	// Split the string up around code block tags
	splitStrings := strings.Split(inMaze, "```")
	
	// Iterate over all strings
	for _, realMaze = range splitStrings {
		// Find whitch one contains the actual maze, and terminate the loop when found
		if strings.HasPrefix(realMaze, "maze\n") {
			realMaze = strings.TrimPrefix(realMaze, "maze\n")
			fmt.Println(realMaze)
			break
		}
	}
	mazeReader := strings.NewReader(realMaze)
	
	// Count the number of underscores to determine the width
	// Create variable to store the number of underscores
	var numUnderscore int
	
	for {
		// Read a maze charachter
		currentChar, err := mazeReader.ReadByte()
		// Detect EOF
		if err != nil { break }
		// If there are still underscores
		if currentChar == '_' {
			// Add one to the underscore count
			numUnderscore += 1
		} else {
			// Otherwise we have reached the end, so calculate the width
			maze.width = int(numUnderscore / 2)
			// and exit
			break
		}
	}
	
	maze.readRow(mazeReader)
}

package maze

import "strings"
import "fmt"

// Calculate the width of the maze and return weather or not it was sucsessfull
func (maze *DummyMaze) calcWidth(mazeReader *strings.Reader) bool {
	// Count the number of underscores to determine the width
	// Create variable to store the number of underscores
	var numUnderscore int
	
	for {
		// Read a maze charachter
		currentChar, err := mazeReader.ReadByte()

		switch {
			// Make shure we havent reached the end
			case err != nil:
				return false

			// If there are still underscores
			case currentChar == '_':
				// Add one to the underscore count
				numUnderscore += 1

			// Otherwise we have reached the end, so calculate the width and exit
			default:
				maze.width = int(numUnderscore / 2)
				fmt.Println("width:", maze.width)
				return true
		}
	}
}

// Reads one row of the maze from string and returns weather or not reading was sucsessfull
func (maze *DummyMaze) calcHeight(mazeReader *strings.Reader) bool {
	if mazeReader.Len() <= 6 {
		return false
	}
	
	// Calculate the height of the maze
	for {
		// Read charachter
		currentChar, err := mazeReader.ReadByte()

		// Make shure we haven't reached the end of the line
		if currentChar == '\n' {
			// Add one to maze height
			maze.height += 1
		} else if err != nil {
			maze.height += 1
			return true
		}
	}
	return false
}

// Reads the maze from a string including finding the start
func (maze *DummyMaze) ReadMaze(inMaze string) {
	// Holds the actual maze
	var realMaze string
	
	// Get only the maze and nothing else
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

	// Calculate the width of the maze and return if failed
	if !maze.calcWidth(mazeReader) {
		return
	}

	maze.calcHeight(mazeReader)
	fmt.Println("height:", maze.height)
}

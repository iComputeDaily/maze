package maze

import "time"
import "fmt"
import "math/rand"

func main() {
	var maze Maze = &GTreeMaze{}
	
	// Initalize the random number generator
	rand.Seed(time.Now().UTC().UnixNano())
	
	err := maze.Generate(3, 3)
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	stringyMaze, err := maze.Stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	fmt.Println(stringyMaze)
}

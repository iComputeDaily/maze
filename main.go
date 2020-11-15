package main

import "time"
import "fmt"
import "math/rand"

func main() {
	var maze maze
	
	// Initalize the random number generator
	rand.Seed(time.Now().UTC().UnixNano())
	
	err := maze.generate()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	stringyMaze, err := maze.stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	fmt.Println(stringyMaze)
}

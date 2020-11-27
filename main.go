package main

import "time"
import "fmt"
import "math/rand"

type maze interface{
	generate()  error
	stringify() (string, error)
}

func main() {
	var maze maze = &gTreeMaze{}
	
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

package main

import "fmt"

func main() {
	var maze maze
	
	err := maze.init()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	err = maze.generate()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	stringyMaze, err := maze.stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	fmt.Println(stringyMaze)
}

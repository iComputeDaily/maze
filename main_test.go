package main

import "github.com/pkg/profile"
import "time"
import "os"
import "fmt"
import "math/rand"
import "testing" 

type maze interface{
	generate()  error
	stringify() (string, error)
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	
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
	
	
	
	maze = &kruskalMaze{}
	
	err = maze.generate()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	stringyMaze, err = maze.stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	fmt.Println(stringyMaze)
}

func TestMain(m *testing.M) {
	// Initalize the random number generator
	rand.Seed(time.Now().UTC().UnixNano())
	
	// prof := profile.Start(profile.CPUProfile)
	
	code := m.Run()
	
	// prof.Stop()
	
	os.Exit(code)
}

func BenchmarkGTree(b *testing.B) {
	var maze maze = &gTreeMaze{}
	
	err := maze.generate()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	_, err = maze.stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	// fmt.Println(stringyMaze)
}

func BenchmarkKruskals(b *testing.B) {
	var maze maze = &kruskalMaze{}
	
	err := maze.generate()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	_, err = maze.stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	// fmt.Println(stringyMaze)
}

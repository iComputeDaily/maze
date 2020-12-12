package maze

import "time"
import "os"
import "fmt"
import "math/rand"
import "testing" 

func TestMain(m *testing.M) {
	// Initalize the random number generator
	rand.Seed(time.Now().UTC().UnixNano())
	
	code := m.Run()
	
	os.Exit(code)
}

func BenchmarkGTree(b *testing.B) {
	var maze Maze = &GTreeMaze{}
	
	err := maze.Generate(40, 5000)
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	stringyMaze, err := maze.Stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	fmt.Println(stringyMaze)
}

func BenchmarkKruskals(b *testing.B) {
	var maze Maze = &KruskalMaze{}
	
	err := maze.Generate(40, 5000)
	if err != nil {
		fmt.Println("How did we get here?")
	}
	
	stringyMaze, err := maze.Stringify()
	if err != nil {
		fmt.Println("How did we get here?")
	}
	fmt.Println(stringyMaze)
}

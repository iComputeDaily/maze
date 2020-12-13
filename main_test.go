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
	
	maze.Generate(40, 5000)
	
	stringyMaze := maze.Stringify()
	
	fmt.Println(stringyMaze)
}

func BenchmarkKruskals(b *testing.B) {
	var maze Maze = &KruskalMaze{}
	
	maze.Generate(40, 5000)
	
	stringyMaze := maze.Stringify()
	
	fmt.Println(stringyMaze)
}

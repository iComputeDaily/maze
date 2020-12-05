package main

import "time"
import "os"
import "fmt"
import "math/rand"
import "testing" 

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

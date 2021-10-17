package main

import (
	"fmt"
	"maze_solver/config"
	"maze_solver/maze"
	"os"
)

func main() {
	// Getting the path to the Map file specified in the Config file
	mapFile, err := config.GetMapFile()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}
	// Calling the MazeSolver to solve the Maze represented in the Map
	maze.MazeSolver(mapFile)
}

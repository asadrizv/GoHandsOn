package maze

import (
	"fmt"
)

/*
	A function to display the solution to the Maze.

    Parameters:
    argument1 (string): Path to the Map File representing the Maze

*/
func MazeSolver(mapPath string) {
	path, err := findExit(mapPath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(path)
	}
}

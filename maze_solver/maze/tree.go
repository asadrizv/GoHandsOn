package maze

import (
	"errors"
	"maze_solver/json"
	"maze_solver/models"
)

type mazeTree struct {
	name      string
	nextNodes []*mazeTree
	parent    *mazeTree
}

/*
	A function to find the Shortest Path to Exit the Maze, the core logic is based on Breadth-First Search
	and aims to return as soon as the shortest exit is found.

    Parameters:
    argument1 (string): Path to the Map File

    Returns:
    string[]: A list of moves the Player must take to reach the exit, the fastest.
	error: Returns an error in case the Exit was never found or when an internal function failed.

*/
func findExit(mapPath string) ([]string, error) {
	tree := mazeTree{name: models.KMazeStart, nextNodes: nil, parent: nil}
	json, err := json.ParseJson(mapPath)
	path := []string{}

	if err != nil {
		return path, err
	}

	loadMazeTree(json, &tree)

	var nodeQue []*mazeTree

	nodeQue = append(nodeQue, &tree)
	for len(nodeQue) > 0 {
		var currentNode *mazeTree
		currentNode, nodeQue = nodeQue[0], nodeQue[1:]

		if currentNode.name == models.KMazeExit {
			return getPath(currentNode), nil
		}

		if len(currentNode.nextNodes) == 0 {
			continue
		}
		nodeQue = append(nodeQue, currentNode.nextNodes...)

	}
	return path, errors.New(models.KErrorNotFound)
}

/*
	A Recursive function to load the Map containing the parsed JSON data into a Tree, to aid Tree based traversal.

    Parameters:
    argument1 map[string]interface{}: Map containing the JSON data.
	argument2 *mazeTree: A pointer to a tree which will have the data loaded into it on completion of the function.

*/
func loadMazeTree(data map[string]interface{}, parent *mazeTree) {
	for key, val := range data {
		var curr mazeTree
		curr.name = key
		curr.parent = parent

		switch val.(type) {
		case string:
			nextNode := mazeTree{val.(string), nil, &curr}
			curr.nextNodes = append(curr.nextNodes, &nextNode)

		case map[string]interface{}:
			loadMazeTree(val.(map[string]interface{}), &curr)
		}

		parent.nextNodes = append(parent.nextNodes, &curr)
	}
}

/*
	An Iterative Function which would start from a Node in the tree and back-track to the root of the tree to print out all parent Nodes.
	This will give us the list of moves to reach a specific point (passed as a parameter to this function)

    Parameters:
    argument1 *mazeTree: Node from which the Algorithm should back-track

	Returns:
	[]string: A list of moves required to reach the Destination node passed in as the argument.

*/
func getPath(destination *mazeTree) []string {
	path := []string{}
	for destination.name != models.KMazeStart {
		path = append([]string{destination.name}, path...)
		destination = destination.parent
	}
	return path
}

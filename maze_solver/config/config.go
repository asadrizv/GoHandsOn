package config

import (
	"maze_solver/json"
	"maze_solver/models"
)

/*
	A utility function to get the Map file from the Config.

	Returns:
	string: The path to the Map file mentioned in the Config file.
	error: The error encountered in case there is a failure internally.

*/
func GetMapFile() (string, error) {
	data, err := json.ParseJson(models.KConfigPath)
	if err != nil {
		return "", err
	}
	return data["map_file"].(string), nil
}

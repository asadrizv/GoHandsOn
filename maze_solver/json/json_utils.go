package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*
	A utility function to parse JSON files into a Map format which could easily be traversed at runtime.

    Parameters:
    argument1 (string): Path to the Map File

    Returns:
    map[string]interface{}: Parsed Map of Key-Value pairs representing respective Key-Value pairs in the JSON file

*/
func ParseJson(path string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(byteValue, &result)
	return result, nil

}

package json

import (
	"testing"
)

func TestMazeSolver(t *testing.T) {
	path := "../data/maps/single_exit.json"
	data, err := ParseJson(path)

	if err != nil {
		t.Errorf(err.Error())
	}

	if data["forward"] != "tiger" {
		t.Errorf("incorrect data parsed, got: %s, want: %s.", data["forward"], "tiger")
	}

}

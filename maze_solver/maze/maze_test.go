package maze

import (
	"maze_solver/models"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	MazeSolver(models.TestMazeSingleExitMap)
}

package solver

import (
	"fmt"
	"errors"
	"math/rand"
	"github.com/Baabah/maze/model"
	"time"
)

func Solve(maze model.Maze) {
	startPos, errorResponse := findStart(maze)
	if errorResponse != nil {
		fmt.Println(errorResponse)
		return
	}

	fmt.Println("Starting at", startPos)

	rand.Seed(time.Now().UTC().UnixNano())
	var steps int
	for {
		steps++
		possibleNewPositions := findPathOptions(maze, startPos)
		fmt.Println(possibleNewPositions)
		newPos := possibleNewPositions[rand.Intn(len(possibleNewPositions))]
		fmt.Println("Moving to", newPos)
		if maze.IsEnd(newPos) == true {
			fmt.Println("Found the end at", newPos)
			fmt.Println("Steps taken", steps)
			return
		}
		startPos = newPos
	}
	return
}

func findPathOptions(maze model.Maze, startPos model.Position) []model.Position {
	var possibleNewPositions []model.Position
	possiblePositions := []model.Position{
		{startPos.X, startPos.Y - 1}, // 1 up
		{startPos.X, startPos.Y + 1}, // 1 down
		{startPos.X - 1, startPos.Y}, // 1 left
		{startPos.X + 1, startPos.Y}, // 1 right
	}

	for _, possiblePosition := range possiblePositions {
		if maze.Contains(possiblePosition) {
			if maze.GetValueByPosition(possiblePosition) == ' ' {
				possibleNewPositions = append(possibleNewPositions, possiblePosition)
			}
		}
	}
	return possibleNewPositions
}

func findStart(maze model.Maze) (model.Position, error) {
	for y, row := range maze.Points {
		for x, element := range row {
			if element == '*' {
				position := model.Position{x, y}
				return position, nil;
			}
		}
	}
	return model.Position{}, errors.New("No start found")
}
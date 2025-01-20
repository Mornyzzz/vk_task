package alghoritm

import (
	"vk_task/internal/io"
	"vk_task/internal/models"
)

func MazeSolver() error {
	var (
		graph      = make(map[models.Point][]models.Point)
		weights    = make(map[models.Point]int)
		start, end models.Point
	)

	err := io.ReadInput(graph, weights, &start, &end)
	if err != nil {
		return err
	}

	path, err := FindPath(graph, weights, start, end)
	if err != nil {
		return err
	}
	
	io.PrintPath(path, start, end)
	return nil
}

package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vk_task/internal/models"
)

func ReadInput(graph map[models.Point][]models.Point, weights map[models.Point]int, start, end *models.Point) error {
	const op = "readInput"
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	line0 := strings.Split(str[:len(str)-1], " ")
	if len(line0) != 2 {
		return fmt.Errorf("%s: %w", op, errors.New("invalid rows cols input"))
	}

	mazeRows, err := strconv.Atoi(line0[0])
	if err != nil || mazeRows < 1 {
		return fmt.Errorf("%s: %w", op, errors.New("invalid rows"))
	}
	mazeCols, err := strconv.Atoi(line0[1])
	if err != nil || mazeCols < 1 {
		return fmt.Errorf("%s: %w", op, errors.New("invalid cols"))
	}

	for row := range mazeRows {
		str, err = reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		line := strings.Split(str[:len(str)-1], " ")

		if len(line) != mazeCols {
			return fmt.Errorf("%s: %w", op, errors.New("invalid row"))
		}
		for col, weightStr := range line {
			weight, err := strconv.Atoi(weightStr)
			if weight < 0 || weight > 9 {
				return fmt.Errorf("%s: %w", op, errors.New("invalid weight"))
			}
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}

			if weight != 0 {
				current := models.Point{Row: row, Col: col}
				leftNeighbour := models.Point{Row: row, Col: col - 1}
				upNeighbour := models.Point{Row: row - 1, Col: col}

				weights[current] = weight
				if _, exists := weights[leftNeighbour]; col-1 >= 0 && exists {
					graph[current] = append(graph[current], leftNeighbour)
					graph[leftNeighbour] = append(graph[leftNeighbour], current)
				}
				if _, exists := weights[upNeighbour]; row-1 >= 0 && exists {
					graph[current] = append(graph[current], upNeighbour)
					graph[upNeighbour] = append(graph[upNeighbour], current)
				}
			}
		}
	}

	str, err = reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	cordsLine := strings.Split(str[:len(str)-1], " ")
	if len(cordsLine) != 4 {
		return fmt.Errorf("%s: %w", op, errors.New("invalid cords"))
	}
	var cords [4]int
	for i := range 4 {
		cords[i], err = strconv.Atoi(cordsLine[i])
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	*start = models.Point{Row: cords[0], Col: cords[1]}
	if start.Row < 0 || start.Row >= mazeRows || start.Col < 0 || start.Col >= mazeCols {
		return fmt.Errorf("%s: %w", op, errors.New("invalid start cords"))
	}

	*end = models.Point{Row: cords[2], Col: cords[3]}
	if end.Row < 0 || end.Row >= mazeRows || end.Col < 0 || end.Col >= mazeCols {
		return fmt.Errorf("%s: %w", op, errors.New("invalid end cords"))
	}

	return nil
}

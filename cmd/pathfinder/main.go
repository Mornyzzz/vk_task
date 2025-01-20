package main

import (
	"fmt"
	"os"
	"vk_task/internal/alghoritm"
)

// Решение задачи с помощью алгоритма Дейкстры

func main() {
	err := alghoritm.MazeSolver()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

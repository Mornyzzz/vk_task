package io

import (
	"bufio"
	"os"
	"strconv"
	"vk_task/internal/models"
)

func PrintPath(path map[models.Point]models.Point, start, end models.Point) {
	var (
		out     = bufio.NewWriter(os.Stdout)
		current = start
	)

	for current != end {
		out.WriteString(strconv.Itoa(current.Row))
		out.WriteString(" ")
		out.WriteString(strconv.Itoa(current.Col))
		out.WriteString("\n")
		current = path[current]
	}
	out.WriteString(strconv.Itoa(current.Row))
	out.WriteString(" ")
	out.WriteString(strconv.Itoa(current.Col))
	out.WriteString("\n")
	out.WriteString(".")
	out.Flush()
}

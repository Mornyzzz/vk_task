package alghoritm

import (
	"container/heap"
	"errors"
	"vk_task/internal/models"
)

const MaxInt = int(^uint(0) >> 1)

func FindPath(graph map[models.Point][]models.Point, weights map[models.Point]int, start, end models.Point) (map[models.Point]models.Point, error) {
	var (
		path      = make(map[models.Point]models.Point)
		distances = make(map[models.Point]int)
		pq        = make(PriorityQueue, 0)
	)

	// установление максимальных значений дистанции
	for w := range weights {
		distances[w] = MaxInt
	}

	// куча для получения в первую очередь клеток с минимальной дистанцией от старта до них

	// в path записываем предыдущие выбранные клетки с минимальной дистанцией
	// тем самым сохраняем путь от старта до финиша
	heap.Init(&pq)
	heap.Push(&pq, Item{end, weights[end]})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(Item)

		// если дошли до старта - возвращаем путь, он минимальный
		if current.Node == start {
			return path, nil
		}

		//алгоритм дейкстры
		if current.Distance <= distances[current.Node] {
			for _, neighbor := range graph[current.Node] {
				newDistance := current.Distance + weights[neighbor]
				if newDistance < distances[neighbor] {
					distances[neighbor] = newDistance
					path[neighbor] = current.Node
					heap.Push(&pq, Item{neighbor, newDistance})
				}
			}
		}
	}

	return nil, errors.New("no path")
}

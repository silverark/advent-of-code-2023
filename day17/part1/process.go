package part1

import (
	"log"
	"strconv"
)

func process(input []string) int {

	//Breadth-First Search (BFS)
	heatLoss := bfs(input, 1, 3)
	var result = heatLoss
	return result
}

type Vertex struct {
	x, y int
}

type Edge struct {
	point         Vertex
	direction     Vertex
	straightCount int
}

func insideGrid(grid []string, pos Vertex) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}

func bfs(grid []string, minStraight, maxStraight int) int {
	start := Vertex{0, 0}
	end := Vertex{len(grid[0]) - 1, len(grid) - 1}
	// Add the first two edges to check.
	pointsToCheck := []Edge{{start, Vertex{1, 0}, 0}, {start, Vertex{0, 1}, 0}}
	visited := map[Edge]int{{start, Vertex{0, 0}, 0}: 0}
	minHeatLoss := 1 << 30

	for len(pointsToCheck) > 0 {

		//TODO: Try to implement a priority queue
		current := pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]

		if current.point == end && current.straightCount >= minStraight {
			minHeatLoss = min(minHeatLoss, visited[current])
			continue
		}

		straightEdge := Edge{Vertex{current.point.x + current.direction.x, current.point.y + current.direction.y}, current.direction, current.straightCount + 1}
		if insideGrid(grid, straightEdge.point) && current.straightCount < maxStraight {
			totalHeatLoss := visited[current] + int(grid[straightEdge.point.y][straightEdge.point.x]-'0')
			if val, found := visited[straightEdge]; !found || val > totalHeatLoss {
				visited[straightEdge] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, straightEdge)
			}
		}

		leftDirection := Vertex{current.direction.y, -current.direction.x}
		leftEdge := Edge{Vertex{current.point.x + leftDirection.x, current.point.y + leftDirection.y}, leftDirection, 1}
		if insideGrid(grid, leftEdge.point) && current.straightCount >= minStraight {
			totalHeatLoss := visited[current] + int(grid[leftEdge.point.y][leftEdge.point.x]-'0')
			if val, found := visited[leftEdge]; !found || val > totalHeatLoss {
				visited[leftEdge] = totalHeatLoss
				pointsToCheck = append(pointsToCheck, leftEdge)
			}
		}

		rightDirection := Vertex{-current.direction.y, current.direction.x}
		rightState := Edge{Vertex{current.point.x + rightDirection.x, current.point.y + rightDirection.y}, rightDirection, 1}
		if insideGrid(grid, rightState.point) && current.straightCount >= minStraight {
			heatLossWeight, _ := strconv.Atoi(string(grid[rightState.point.y][rightState.point.x]))
			predictedHeatloss := visited[current] + heatLossWeight
			if val, found := visited[rightState]; !found || val > predictedHeatloss {
				visited[rightState] = predictedHeatloss
				pointsToCheck = append(pointsToCheck, rightState)
			}
		}
	}

	for edge, value := range visited {
		log.Println("value:", value, "edge:", edge.point, edge.direction, edge.straightCount)
	}

	return minHeatLoss
}

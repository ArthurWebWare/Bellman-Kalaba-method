package main

import (
	"fmt"
)

const maxWeight = 1000

var (
	graph                                              [100][100]int
	currentPath, bestPath                              [100]int
	numVertices, minDistance, z, numSteps, vertexIndex int
	pathIndex                                          = 1
	currentElem                                        = maxWeight + 1
)

func printGraph(graph [100][100]int, numVertices, numCols int) {
	width := numCols * 5
	fmt.Print(" ")
	for i := 1; i <= width; i++ {
		if i == width {
			fmt.Print(" ")
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println()
	fmt.Print("|    |")
	for i := 1; i < numCols; i++ {
		fmt.Printf("x%2d |", i)
	}
	fmt.Println()
	fmt.Print("|")
	for i := 1; i <= width; i++ {
		if i%5 == 0 {
			fmt.Print("|")
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println()

	for i := 1; i < numVertices; i++ {
		if i >= numCols {
			fmt.Printf("|V%2d |", pathIndex)
			pathIndex++
		} else {
			fmt.Printf("|x%2d |", i)
		}
		for j := 1; j < numCols; j++ {
			if graph[i][j] == maxWeight {
				fmt.Print(" +  |")
			} else {
				fmt.Printf("%4d|", graph[i][j])
			}
		}
		fmt.Println()
		fmt.Print("|")
		for k := 1; k <= width; k++ {
			if k%5 == 0 {
				fmt.Print("|")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func findMinPath(currentVertex int) {
	if currentElem < maxWeight {
		for i := 0; i < numSteps; i++ {
			if bestPath[i] == currentElem {
				vertexIndex = i
			}
		}
		for i := 0; i < vertexIndex; i++ {
			currentPath[z] = bestPath[i]
			z++
		}
		currentElem = maxWeight + 1
	}

	currentPath[z] = currentVertex
	z++

	if currentVertex == numVertices-1 {
		fmt.Printf("\n#The minimum path[ %d ] : ", pathIndex)
		pathIndex++
		for i := 0; i < z; i++ {
			fmt.Printf("%3d", currentPath[i])
		}
		fmt.Println()
		numSteps = z

		for i := 0; i < z; i++ {
			bestPath[i] = currentPath[i]
		}
		z = 0
	}

	for nextVertex := currentVertex + 1; nextVertex < numVertices; nextVertex++ {
		if graph[numVertices-1][currentVertex] != maxWeight && graph[currentVertex][nextVertex] != maxWeight {
			if (graph[numVertices-1][currentVertex] - graph[currentVertex][nextVertex]) == graph[numVertices-1][nextVertex] {
				findMinPath(nextVertex)
				currentElem = nextVertex
			}
		}
	}
}

func main() {
	fmt.Print("Enter the number of vertices: ")
	fmt.Scan(&numVertices)
	numVertices++

	fmt.Println("Enter your Weighted Graph: ")
	for i := 1; i < numVertices; i++ {
		for j := 1; j < numVertices; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	fmt.Println("You entered the following weighted graph:")
	printGraph(graph, numVertices, numVertices)

	for i := 1; i < numVertices; i++ {
		graph[numVertices][i] = graph[i][numVertices-1]
	}
	numVertices++

	for {
		minDistance = maxWeight
		for col := 1; col < numVertices; col++ {
			for row := col + 1; row < numVertices; row++ {
				if graph[numVertices-1][row] != maxWeight && graph[col][row] != maxWeight {
					if (graph[numVertices-1][row] + graph[col][row]) < minDistance {
						minDistance = graph[numVertices-1][row] + graph[col][row]
					}
				}
			}
			if col == numVertices-1 {
				minDistance = 0
			}
			graph[numVertices][col] = minDistance
		}

		isMinPath := true
		for i := 1; i < numVertices; i++ {
			if graph[numVertices-1][i] != graph[numVertices][i] {
				isMinPath = false
				break
			}
		}
		if isMinPath {
			numVertices++
			break
		} else {
			numVertices++
		}
	}

	fmt.Println("\n\nBellman-Kalaba algorithm:")

	printGraph(graph, numVertices, numVertices)

	fmt.Printf("\n\nMinimum path length = %d \n", graph[numVertices-1][1])

	fmt.Println("\nMinimal paths:")

	findMinPath(1)
}

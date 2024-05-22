package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
)

const (
	gridSize = 20
	n        = 10
	runs     = 100 // liczba powtórzeń dla każdej liczby
)

type Result struct {
	Number  int
	Average float64
}

func clearGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}
}

func moveWalker(x, y *int, gridSize int) {
	direction := rand.Intn(4)
	switch direction {
	case 0:
		*x = (*x + 1) % gridSize // Move right
	case 1:
		*x = (*x - 1 + gridSize) % gridSize // Move left
	case 2:
		*y = (*y + 1) % gridSize // Move down
	case 3:
		*y = (*y - 1 + gridSize) % gridSize // Move up
	}
}

func calculateDistance(wg *sync.WaitGroup, distanceChan chan<- float64) {
	defer wg.Done()

	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	clearGrid(grid)
	x, y := n, n
	grid[x][y] = 1

	for i := 0; i < n; i++ {
		moveWalker(&x, &y, gridSize)
		clearGrid(grid)
		grid[x][y] = 1
	}

	distance := math.Sqrt(float64((x-n)*(x-n) + (y-n)*(y-n)))
	distanceChan <- distance
}

func calculateAverageDistanceForNumber(number int, resultsChan chan<- Result) {
	var wg sync.WaitGroup
	distanceChan := make(chan float64, runs)
	totalDistance := 0.0

	for i := 0; i < runs; i++ {
		wg.Add(1)
		go calculateDistance(&wg, distanceChan)
	}

	go func() {
		wg.Wait()
		close(distanceChan)
	}()

	count := 0
	for distance := range distanceChan {
		totalDistance += distance
		count++
	}

	averageDistance := totalDistance / float64(count)
	resultsChan <- Result{Number: number, Average: averageDistance}
}

func main() {
	var wg sync.WaitGroup
	resultsChan := make(chan Result, 1000)
	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	for number := 5; number <= 5000; number += 5 {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			calculateAverageDistanceForNumber(number, resultsChan)
		}(number)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		_, err := fmt.Fprintf(file, "%d %.2f\n", result.Number, result.Average)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
	}
}

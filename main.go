package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

const (
	Up         = 0
	Down       = 1
	Left       = 2
	Right      = 3
	UpLeft     = 4
	UpRight    = 5
	DownLeft   = 6
	DownRight  = 7
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var n int
	fmt.Print("Podaj liczbę iteracji: ")
	fmt.Scan(&n)
	gridSize := 2 * n + 1
	grid := make([][]int, gridSize)
	distances := make([]float64, n)


	
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}
	// x, y := n, n
	// grid[x][y] = 1
	// printGrid(grid)

	for j:= 0; j < n; j++ {
		clearGrid(grid)
		x, y := n, n
		grid[x][y] = 1

		for i := 0; i < n; i++ {
			moveWalker(&x, &y, gridSize)
			clearGrid(grid)
			grid[x][y] = 1
			// printGrid(grid)
		}
		// fmt.Printf("Na koniec wędrowca jest w punkcie %d, %d\n", x, y)
		distance := math.Sqrt(float64((x-n)*(x-n)+(y-n)*(y-n)))
		fmt.Printf("Odległość od punktu startowego: %.2f\n", distance)
		// distances = append(distances, distance)
		distances[j] = distance
	}

	sum := 0.0
    for _, v := range distances {
        sum += v
    }
	average := float64(sum) / float64(len(distances))
    fmt.Printf("Średnia wartość w tablicy to %.2f\n", average)
}

func moveWalker(x, y *int, size int) {
	direction := rand.Intn(8)
	switch direction {
	case Up:
		if *x > 0 {
			*x--
		}
	case Down:
		if *x < size-1 {
			*x++
		}
	case Left:
		if *y > 0 {
			*y--
		}
	case Right:
		if *y < size-1 {
			*y++
		}
	case UpLeft:
		if *x > 0 && *y > 0 {
			*x--
			*y--
		}
	case UpRight:
		if *x > 0 && *y < size-1 {
			*x--
			*y++
		}
	case DownLeft:
		if *x < size-1 && *y > 0 {
			*x++
			*y--
		}
	case DownRight:
		if *x < size-1 && *y < size-1 {
			*x++
			*y++
		}
	}
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func clearGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}
}

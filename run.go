package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"
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
	runs       = 100
	maxWorkers = 10  
)

type Result struct {
	Number  int
	Average float64
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

func calculateDistance(number int, wg *sync.WaitGroup, distanceChan chan<- float64) {
	defer wg.Done()

	x, y := number, number

	for i := 0; i < number; i++ {
		moveWalker(&x, &y, 2*number)
	}

	distance := math.Sqrt(float64((x-number)*(x-number) + (y-number)*(y-number)))
	distanceChan <- distance
}

func worker(id int, jobs <-chan int, resultsChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range jobs {
		var wgRuns sync.WaitGroup
		distanceChan := make(chan float64, runs)
		totalDistance := 0.0

		for i := 0; i < runs; i++ {
			wgRuns.Add(1)
			go calculateDistance(number, &wgRuns, distanceChan)
		}

		go func() {
			wgRuns.Wait()
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
}

func main() {
	jobs := make(chan int, 200)
	resultsChan := make(chan Result, 200)
	var wg sync.WaitGroup

	for w := 0; w < maxWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, resultsChan, &wg)
	}


	go func() {
		for number := 1; number <= 10000; number++ {
			jobs <- number
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Write results to file
	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	for result := range resultsChan {
		_, err := fmt.Fprintf(file, "%d %.2f\n", result.Number, result.Average)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
	}
}

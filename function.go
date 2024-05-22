package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// DataPoint struct to hold the data from the file
type DataPoint struct {
	Number  int
	Average float64
}

// Function to load data from the file
func loadData(filename string) ([]DataPoint, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []DataPoint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		number, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		average, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, err
		}

		data = append(data, DataPoint{Number: number, Average: average})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// Function to calculate the mean squared error
func meanSquaredError(data []DataPoint, f func(float64) float64) float64 {
	var mse float64
	for _, point := range data {
		x := float64(point.Number)
		y := point.Average
		mse += (y - f(x)) * (y - f(x))
	}
	return mse / float64(len(data))
}

// Define candidate functions
var functions = map[string]func(float64) float64{
	"sqrt(x)":  math.Sqrt,
	"ln(x)":    math.Log,
	"x^2":      func(x float64) float64 { return x * x },
	"x^0.5":    func(x float64) float64 { return math.Pow(x, 0.5) },
	"x^1.5":    func(x float64) float64 { return math.Pow(x, 1.5) },
	"x":        func(x float64) float64 { return x },
	"1/x":      func(x float64) float64 { return 1 / x },
	"sin(x)":   math.Sin,
	"cos(x)":   math.Cos,
	"exp(x)":   math.Exp,
	"log2(x)":  math.Log2,
	"log10(x)": math.Log10,
	"x/2": func(x float64) float64 { return x / 2 },
}

func main() {
	filename := "results.txt"
	data, err := loadData(filename)
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	if len(data) == 0 {
		log.Fatalf("No data loaded")
	}

	// Calculate mean squared error for each function
	bestFunc := ""
	bestMSE := math.MaxFloat64

	for name, fn := range functions {
		mse := meanSquaredError(data, fn)
		fmt.Printf("Function %s has MSE: %.4f\n", name, mse)
		if mse < bestMSE {
			bestMSE = mse
			bestFunc = name
		}
	}

	fmt.Printf("The best fitting function is: %s with MSE: %.4f\n", bestFunc, bestMSE)
}

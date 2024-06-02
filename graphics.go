package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	filePath := "results.txt" 
	dataX, dataY := readDataFromFile(filePath)

	points := make(plotter.XYs, len(dataX))
	for i := range dataX {
		points[i].X = dataX[i]
		points[i].Y = dataY[i]
	}

	p := plot.New()

	p.Title.Text = "Simulation results"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	p.Add(s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "image.png"); err != nil {
		panic(err)
	}
	fmt.Println("Image saved to image.png")
}


func readDataFromFile(filePath string) ([]float64, []float64) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var dataX, dataY []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		x, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			continue
		}
		y, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			continue
		}
		dataX = append(dataX, x)
		dataY = append(dataY, y)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return dataX, dataY
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
}

func main() {
	var n int
	fmt.Println("Podaj wartość n:")
	fmt.Scanln(&n)

	boardSize := 2 * n + 1

	board := make([][]string, boardSize)
	for i := range board {
		board[i] = make([]string, boardSize)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	rand.Seed(time.Now().UnixNano())

	startX := n
	startY := n
	currPos := Point{startX, startY}

	
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			dist := math.Sqrt(float64((i-startY)*(i-startY) + (j-startX)*(j-startX)))
			if dist == math.Sqrt(float64(n)) {
				board[i][j] = "O" 
			}
		}
	}

	for i := 0; i < n; i++ {
		move := rand.Intn(7)
		switch move {
		case 0:
			if currPos.x > 0 {
				currPos.x--
			}
		case 1:
			if currPos.x < boardSize-1 {
				currPos.x++
			}
		case 2:
			if currPos.y > 0 {
				currPos.y--
			}
		case 3:
			if currPos.y < boardSize-1 {
				currPos.y++
			}
		case 4:
			if currPos.y < boardSize-1 && currPos.x < boardSize-1 {
				currPos.y++
				currPos.x++
			}
		case 5: 
			if currPos.y < boardSize-1 &&  currPos.x > 0 {
				currPos.y++
				currPos.x--
			}
		
		case 6:
			if currPos.y > 0 && currPos.x < boardSize-1 {
				currPos.y--
				currPos.x++
			}
		case 7:
			if  currPos.y > 0  &&  currPos.x > 0{
				currPos.y--
				currPos.x--
			}
		}
		for i := range board {
			for j := range board[i] {
				if board[i][j] == "X" {
					board[i][j] = "."
				}
			}
		}
		board[currPos.y][currPos.x] = "X"
		printBoard(board)
		time.Sleep(500 * time.Millisecond)
		clearConsole()
	}

	fmt.Println("Ostateczna plansza:")
	printBoard(board)
}


func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, " "))
	}
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

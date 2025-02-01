package main

import (
	"fmt"
)

func drawChessBoard(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	var size int
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&size)
	drawChessBoard(size)
}

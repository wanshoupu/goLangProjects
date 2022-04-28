package main

import "fmt"

/*
BUGS:
1. Boundary checking: len(board) <= x not len(board) < x
2. When counting live cells within 3x3 patch, the center cell should be excluded
3. Whenever dereferencing board for cells, one should mask the bits to take only the lowest bit.
 */
func gameOfLife(board [][]int) {
	for i, row := range board {
		for j := range row {
			board[i][j] |= calc(board, i, j) << 1
		}
	}

	// shift down to finalize
	for i, row := range board {
		for j := range row {
			board[i][j] >>= 1
		}
	}
}

func calc(board [][]int, i, j int) int {
	count := sum(board, i, j)
	// Any live cell with fewer than two live neighbors dies, as if caused by under-population.
	// Any live cell with more than three live neighbors dies, as if by over-population..
	if count < 2 || 3 < count {
		return 0
	}
	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	if count == 3 {
		return 1
	}
	// Any live cell with two or three live neighbors lives on to the next generation.
	return board[i][j] & 1
}

func sum(board [][]int, m, n int) int {
	s := 0
	for i := -1; i < 2; i++ {
		x := m + i
		if x < 0 || len(board) <= x {
			continue
		}
		for j := -1; j < 2; j++ {
			y := n + j
			if y < 0 || len(board[x]) <= y {
				continue
			}
			s += board[x][y] & 1
		}
	}
	return s - board[m][n] & 1
}

func main() {
	tests := [][][]int{
		{
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		},
	}
	for _, test := range tests {
		runGame(test)
	}

}

func runGame(test [][]int) {
	for i := 0; i < 10; i++ {
		fmt.Println(test)
		gameOfLife(test)
	}
}

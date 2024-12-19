package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read and parse the input
	file, err := os.ReadFile("day4.input")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	grid := [][]rune{}
	for _, line := range lines {
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	target := []rune{'M', 'A', 'S'}
	height := len(grid)
	width := len(grid[0])
	ans := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] != 'X' {
				continue
			}

			// Horizontal backward
			if j > 2 && equals(grid[i][j-3:j], []rune{'S', 'A', 'M'}) {
				ans++
			}
			// Horizontal forward
			if j < width-3 && equals(grid[i][j+1:j+4], target) {
				ans++
			}
			// Vertical upward
			if i > 2 && matchesVertical(grid, target, i, j, -1) {
				ans++
			}
			// Vertical downward
			if i < height-3 && matchesVertical(grid, target, i, j, 1) {
				ans++
			}
			// Diagonal upward-left
			if i > 2 && j > 2 && matchesDiagonal(grid, target, i, j, -1, -1) {
				ans++
			}
			// Diagonal downward-left
			if i < height-3 && j > 2 && matchesDiagonal(grid, target, i, j, 1, -1) {
				ans++
			}
			// Diagonal upward-right
			if i > 2 && j < width-3 && matchesDiagonal(grid, target, i, j, -1, 1) {
				ans++
			}
			// Diagonal downward-right
			if i < height-3 && j < width-3 && matchesDiagonal(grid, target, i, j, 1, 1) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

// Helper function to compare two slices
func equals(slice1 []rune, slice2 []rune) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

// Helper function to check vertical patterns
func matchesVertical(grid [][]rune, target []rune, i, j, step int) bool {
	for k := 1; k < 4; k++ {
		if grid[i+step*k][j] != target[k-1] {
			return false
		}
	}
	return true
}

// Helper function to check diagonal patterns
func matchesDiagonal(grid [][]rune, target []rune, i, j, stepI, stepJ int) bool {
	for k := 1; k < 4; k++ {
		if grid[i+stepI*k][j+stepJ*k] != target[k-1] {
			return false
		}
	}
	return true
}

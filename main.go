package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

//http://www.theasciicode.com.ar/extended-ascii-code/box-drawing-character-single-line-upper-left-corner-ascii-code-218.html
func main() {
	height, width := dimensions()

	var buffer bytes.Buffer

	generation := getFirstGeneration()

	// Contains the output
	buffer, generation = drawGeneration(buffer, width, height, generation)

	fmt.Println(buffer.String())
	buffer.Reset()
	fmt.Println("-----")

	foo := nextGeneration(generation)

	buffer, _ = drawGeneration(buffer, width, height, foo)

	fmt.Println(buffer.String())
	// // Should loop trough the generations
	// NextGeneration := true
	// for loop := true; loop; loop = NextGeneration {

	// 	clear()

	// 	fmt.Println(buffer.String())

	// 	time.Sleep(time.Second / time.Duration(5))
	// }
}

/**
Determines the dimensions of the game (terminal width/height)
*/
func dimensions() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(out), " ")

	x, err := strconv.Atoi(parts[0])

	y, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", 1))

	return int(x), int(y) // Minus the sides, Minus the top
}

/**
Draws the next generation
*/
func drawGeneration(b bytes.Buffer, width int, height int, generation map[int]map[int]bool) (bytes.Buffer, map[int]map[int]bool) {
	// Draw in between
	var death = color.HiBlackString("0")
	var life = color.HiGreenString("1")

	for x := int(0); x < height; x++ {
		for y := int(0); y < width; y++ {
			// Retrieve the cell state
			if generation[x][y] {
				b.WriteString(life)
			} else {
				b.WriteString(death)
			}
		}
		b.WriteString("\n")
	}

	return b, generation
}

/**
First generation
*/
func getFirstGeneration() map[int]map[int]bool {
	c := map[int]map[int]bool{}

	c[1] = map[int]bool{}
	c[2] = map[int]bool{}
	c[3] = map[int]bool{}
	c[4] = map[int]bool{}
	c[5] = map[int]bool{}

	c[1][3] = true
	c[1][4] = true
	c[1][5] = true
	c[1][6] = true

	c[2][6] = true
	c[2][7] = true

	c[3][6] = true
	c[4][5] = true

	c[1][13] = true
	c[1][14] = true
	c[1][15] = true
	c[1][16] = true

	c[2][12] = true
	c[2][16] = true

	c[3][16] = true
	c[4][16] = true
	c[4][13] = true

	return c
}

/**
Calculates based on the current generation & draws it out using the drawGeneration function
*/
func nextGeneration(currentGeneration map[int]map[int]bool) map[int]map[int]bool {

	nextGeneration := map[int]map[int]bool{}

	for x := range currentGeneration {
		nextGeneration[x] = map[int]bool{}

		for y, value := range currentGeneration[x] {
			nextGeneration[x][y] = value
			// fmt.Println(x)
			// fmt.Println(y)
			// fmt.Println(value)
			// neightbour fmt.Println(currentGeneration[x-1][y-1])

			neighbours := getNeighbours(currentGeneration, x, y)
			// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
			if value == true && neighbours < 2 {
				nextGeneration[x][y] = false
			}

			// Any live cell with two or three live neighbours lives on to the next generation.
			if value == true && (neighbours == 2 || neighbours == 3) {
				nextGeneration[x][y] = true
			}

			// Any live cell with more than three live neighbours dies, as if by overpopulation.
			if value == true && neighbours > 3 {
				nextGeneration[x][y] = false
			}

			// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			// Dead cell with == 3 live neighbours = LIVE
			if value == false && neighbours == 3 {
				nextGeneration[x][y] = true
			}

		}
	}

	return nextGeneration
}

func getNeighbours(generation map[int]map[int]bool, x int, y int) int {
	// Check neighbours
	neighbours := 0

	if generation[x-1][y] {
		neighbours++
	}

	if generation[x+1][y] {
		neighbours++
	}

	if generation[x][y-1] {
		neighbours++
	}

	if generation[x][y+1] {
		neighbours++
	}

	return neighbours
}

/*
Clear output
*/
func clear() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

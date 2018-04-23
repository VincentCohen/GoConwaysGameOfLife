package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	height, width := dimensions()

	var buffer bytes.Buffer

	generation := getFirstGeneration(width, height)

	// Contains the output
	buffer, generation = drawGeneration(buffer, width, height, generation)

	generation = nextGeneration(generation)

	fmt.Println(buffer.String())
	buffer.Reset()

	// Loop trough the generations for ever!
	for loop := true; loop; loop = true {

		generation = nextGeneration(generation)

		buffer, _ = drawGeneration(buffer, width, height, generation)

		clear()

		fmt.Println(buffer.String())
		time.Sleep(time.Second / time.Duration(5))

		buffer.Reset()
	}
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
func getFirstGeneration(width int, height int) map[int]map[int]bool {
	g := map[int]map[int]bool{}
	for x := int(0); x < height; x++ {
		g[x] = map[int]bool{}
		for y := int(0); y < width; y++ {
			g[x][y] = false
		}
	}

	g[1][3] = true
	g[1][4] = true
	g[1][5] = true
	g[1][6] = true
	g[1][13] = true
	g[1][14] = true
	g[1][15] = true
	g[1][16] = true
	// g[3][6] = true
	// g[3][7] = true
	// g[3][8] = true

	g[2][6] = true
	g[2][7] = true

	g[3][6] = true
	g[4][5] = true
	g[2][12] = true
	g[2][16] = true

	g[3][16] = true
	g[4][16] = true
	g[4][13] = true

	return g
}

/**
Calculates based on the current generation & draws it out using the drawGeneration function
*/
func nextGeneration(currentGeneration map[int]map[int]bool) map[int]map[int]bool {

	nextGeneration := map[int]map[int]bool{}

	for x := range currentGeneration {
		nextGeneration[x] = map[int]bool{}

		for y, isAlive := range currentGeneration[x] {
			nextGeneration[x][y] = isAlive
			neighbours := getNeighbours(currentGeneration, x, y)

			// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
			// Any live cell with two or three live neighbours lives on to the next generation.
			// Any live cell with more than three live neighbours dies, as if by overpopulation.
			// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			willLive := false
			if isAlive {
				if neighbours == 2 || neighbours == 3 {
					willLive = true
				}
			} else {
				if neighbours == 3 {
					willLive = true
				}
			}

			nextGeneration[x][y] = willLive
		}
	}

	return nextGeneration
}

func getNeighbours(generation map[int]map[int]bool, x int, y int) int {

	neighbours := 0

	if generation[x][y-1] {
		neighbours++
	}

	if generation[x][y+1] {
		neighbours++
	}

	if generation[x-1][y] {
		neighbours++
	}

	if generation[x-1][y-1] {
		neighbours++
	}

	if generation[x-1][y+1] {
		neighbours++
	}

	if generation[x+1][y] {
		neighbours++
	}

	if generation[x+1][y-1] {
		neighbours++
	}

	if generation[x+1][y+1] {
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

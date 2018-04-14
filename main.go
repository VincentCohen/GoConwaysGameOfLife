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

//http://www.theasciicode.com.ar/extended-ascii-code/box-drawing-character-single-line-upper-left-corner-ascii-code-218.html
func main() {
	height, width := dimensions()

	var buffer bytes.Buffer

	generation := getFirstGeneration()

	// Contains the output
	buffer = top(buffer, width)
	buffer, generation = drawGeneration(buffer, width, height, generation)
	buffer = bottom(buffer, width)

	// Should loop trough the generations
	NextGeneration := true
	for loop := true; loop; loop = NextGeneration {

		clear()

		fmt.Println(buffer.String())

		time.Sleep(time.Second / time.Duration(5))
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

	return int(x - 2), int(y - 3) // Minus the sides, Minus the top
}

/**
Draws the top of the game
*/
func top(b bytes.Buffer, width int) bytes.Buffer {
	b.WriteString("┌")
	for i := int(0); i < width; i++ {
		b.WriteString("─")
	}
	b.WriteString("┐\n")

	return b
}

/**
Draws the next generation
*/
func drawGeneration(b bytes.Buffer, width int, height int, generation map[int]map[int]bool) (bytes.Buffer, map[int]map[int]bool) {
	// Draw in between
	var death = color.HiBlackString("0")
	var life = color.HiGreenString("1")

	for x := int(0); x < height; x++ {
		b.WriteString("│")

		for y := int(0); y < width; y++ {
			// Retrieve the cell state
			if generation[x][y] {
				b.WriteString(life)
			} else {
				b.WriteString(death)
			}
		}
		b.WriteString("│")
		b.WriteString("\n")
	}

	return b, generation
}

/**
Draws the bottom of the game
*/
func bottom(b bytes.Buffer, width int) bytes.Buffer {
	b.WriteString("└")
	for i := int(0); i < width; i++ {
		b.WriteString("─")
	}
	b.WriteString("┘")

	return b
}

/**
First generation
*/
func getFirstGeneration() map[int]map[int]bool {
	c := map[int]map[int]bool{}
	c[0] = map[int]bool{}

	c[0][0] = true
	c[0][1] = true
	c[0][2] = true
	c[0][3] = true
	c[0][4] = true

	c[1] = map[int]bool{}
	c[1][4] = true
	c[1][5] = true
	c[1][6] = true
	c[1][7] = true
	c[1][8] = true

	return c
}

/**
Calculates based on the current generation & draws it out using the drawGeneration function
*/
func nextGeneration(currentGeneration map[int]map[int]bool, x int, y int) bool {

	if currentGeneration[x][y] {
		return true
	}

	// drawGeneration
	return false
	// Decide the normal live / death states as they are advancing generation
	// Track states

	// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
	// Live cell < 2 live = dead

	// Any live cell with two or three live neighbours lives on to the next generation.
	// Live cell with  >2 OR 3  live neighbours = next generation

	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	// Live cell with >3 live neighbours = dead

	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	// Dead cell with == 3 live neighbours = LIVE

}

/*
Clear output
*/
func clear() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

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
	height, width := grid()
	width = width - 2
	height = height - 3 // minus top, minus bottom, minus enter

	var buffer bytes.Buffer

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

	// for cont := true; cont; cont = true {
	// Draw top, start at 2 for the corners
	buffer = drawGridTop(buffer, width)
	buffer = drawCenter(buffer, width, height, c)
	buffer = drawGridBottom(buffer, width)

	// fmt.Println(buffer.String())
	for i := 0; i < 1000; i++ {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()

		time.Sleep(time.Second / time.Duration(120))

		fmt.Println(buffer.String())
	}
}

func grid() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(out), " ")

	x, err := strconv.Atoi(parts[0])

	y, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", 1))

	return int(x), int(y)
}

func drawGridTop(b bytes.Buffer, width int) bytes.Buffer {
	b.WriteString("┌")
	for i := int(0); i < width; i++ {
		b.WriteString("─")
	}
	b.WriteString("┐\n")

	return b
}

func drawCenter(b bytes.Buffer, width int, height int, generation map[int]map[int]bool) bytes.Buffer {
	// Draw in between
	var death = color.HiBlackString("0")
	var life = color.HiGreenString("1")

	for x := int(0); x < height; x++ {
		b.WriteString("│")

		for y := int(0); y < width; y++ {
			// Generate the playing field for the first go
			if generation[x][y] {
				b.WriteString(life)
			} else {
				// Decide the normal live / death states as they are advancing generation
				b.WriteString(death)
			}

			// Track states

			// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
			// Live cell < 2 live = dead

			// Any live cell with two or three live neighbours lives on to the next generation.
			// Live cell with  >2 OR 3  live neighbours = next generation

			// Any live cell with more than three live neighbours dies, as if by overpopulation.
			// Live cell with >3 live neighbours = dead

			// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			// Dead cell with == 3 live neighbours = LIVE

			// Define
			// b.WriteString(life)

		}
		b.WriteString("│")
		b.WriteString("\n")
	}

	return b
}

func drawGridBottom(b bytes.Buffer, width int) bytes.Buffer {
	b.WriteString("└")
	for i := int(0); i < width; i++ {
		b.WriteString("─")
	}
	b.WriteString("┘")

	return b
}

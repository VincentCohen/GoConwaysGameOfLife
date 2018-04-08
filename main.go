package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ┌http://www.theasciicode.com.ar/extended-ascii-code/box-drawing-character-single-line-upper-left-corner-ascii-code-218.html
func main() {
	height, width := grid()

	fmt.Println(width)
	fmt.Println(height)

	var buffer bytes.Buffer

	// for cont := true; cont; cont = true {
	// Draw top, start at 2 for the corners
	buffer = drawGridTop(buffer, width)
	buffer = drawCenter(buffer, width, height)
	buffer = drawGridBottom(buffer, width)

	fmt.Println(buffer.String())

	// clear := exec.Command("clear")
	// clear.Stdout = os.Stdout
	// clear.Run()
	// }
}

func grid() (uint, uint) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(out), " ")

	fmt.Println(parts)

	x, err := strconv.Atoi(parts[0])

	y, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", 1))

	return uint(x), uint(y)
}

func drawGridTop(b bytes.Buffer, width uint) bytes.Buffer {
	b.WriteString("┌")
	for i := uint(2); i < width; i++ {
		b.WriteString("─")
	}
	b.WriteString("┐\n")

	return b
}

func drawCenter(b bytes.Buffer, width uint, height uint) bytes.Buffer {
	// Draw in between
	// var death = "▒"
	var death = " "
	var life = "▓"
	for i := uint(0); i < height; i++ {
		b.WriteString("│")
		for j := uint(2); j < width; j++ {
			if (j % 2) == 0 {
				b.WriteString(life)
			} else {
				b.WriteString(death)
			}
		}
		b.WriteString("│")
		b.WriteString("\n")
	}

	return b
}

func drawGridBottom(b bytes.Buffer, width uint) bytes.Buffer {
	b.WriteString("└")
	for i := uint(2); i < width; i++ {
		b.WriteString("─")
	}
	b.WriteString("┘")

	return b
}


package main

import (
	"post6.net/gohexdump/internal/screen"
	"os"
	"io"
	"flag"
)

var height, width int

func init() {

	flag.IntVar(&height, "height", 720, "height")
	flag.IntVar(&width, "width", 1280, "width")
}

func getPositions(w, h int) []int{

	coords := screen.NewHexScreen().Coords()

	var positions = make([]int, len(coords))

	lx, ly, hx, hy := 9001., 9001., -1., -1.

	for _,pos := range coords {
		x, y := pos.X, pos.Y

		if x < lx {
			lx = x
		}
		if x > hx {
			hx = x
		}
		if y < ly {
			ly = y
		}
		if y > hy {
			hy = y
		}
	}

	fx := float64(w-1)/(hx-lx)
	fy := float64(h-1)/(hy-ly)

	 for i,pos := range coords {
		x, y := int((pos.X-lx)*fx), int((pos.Y-ly)*fy)
		if x > (w-1) {
			x = w-1
		}
		if y > (h-1) {
			y = h-1
		}
		if i % 16 == 15 {
			positions[i] = -1
		} else {
			positions[i] = x + y*width
		}
	}

	return positions
}

func main() {

	flag.Parse()

	positions := getPositions(width, height)
	inframe := make([]byte, width*height)
	outframe := make([]byte, len(positions))

	for {
		if _, err := io.ReadFull(os.Stdin, inframe); err != nil {
			break
		}

		for i := range outframe {
			if i % 16 != 15 {
				outframe[i] = inframe[positions[i]]
			}
		}

		os.Stdout.Write(outframe)
	}
}


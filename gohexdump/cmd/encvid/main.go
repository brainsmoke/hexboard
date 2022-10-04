
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

	info := screen.GetScreenInfo()

	var positions = make([]int, len(info.Segments))

	lx, ly, hx, hy := 9001., 9001., -1., -1.

	for _,seg := range info.Segments {
		x, y := seg.Position.X, seg.Position.Y

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

	 for i,seg := range info.Segments {
		x, y := int((seg.Position.X-lx)*fx), int((seg.Position.Y-ly)*fy)
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


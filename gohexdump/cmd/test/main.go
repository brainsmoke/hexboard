
package main

import (
	"post6.net/gohexdump/internal/screen"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/drivers"
	"os"
	"fmt"
	"bufio"
	"flag"
	"time"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

const (
	nop = iota
	left
	down
	up
	right
	quit
)

func cmdHandler(file *os.File, events chan<- int) {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		switch scanner.Text() {
		case "h":
			events <- left
		case "j":
			events <- down
		case "k":
			events <- up
		case "l":
			events <- right
		case "quit":
			events <- quit
			break
		}
	}

	close(events)
}

var normal, mid, hi, lowoop, hiwoop screen.Style

func writeScreen(s screen.HexScreen, x, y int) {
	s.Hold()


	for xi:= 0; xi<16; xi++ {

		if y == xi {
			s.SetStyle(hi)
		} else {
			s.SetStyle(mid)
		}
		s.WriteOffset(fmt.Sprintf("      %02X", xi*16), xi)

		if x == xi {
			s.SetStyle(hi)
		} else {
			s.SetStyle(mid)
		}
		s.WriteTitle(fmt.Sprintf("%X", xi), 5+8+1+xi*3+4*(xi/8))

		for yi:= 0; yi<16; yi++ {
			if x == xi && y == yi {
				s.SetStyle(hiwoop)
			} else if x == xi || y == yi {
				s.SetStyle(lowoop)
			} else {
				s.SetStyle(normal)
			}
			c := xi+yi*16
			s.WriteHexField(fmt.Sprintf("%02X",  c), c)
			printC := c
			if c < 0x20 || c > 0x7e {
				printC = int('.')
			}
			s.WriteAsciiField(fmt.Sprintf("%c", printC), c)
		}
	}
	s.Update()
}

func main() {

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	s := screen.NewHexScreen()

	s.SetFont(font.GetFont())

	multi, screenChan := screen.NewMultiScreen()

	filters := []screen.Filter { screen.DefaultGamma(), screen.NewAfterGlowFilter(.9) }

	screenChan <- screen.NewFilterScreen(s, filters)

	q := make(chan bool)
	events := make(chan int)

	go cmdHandler(os.Stdin, events)

	go screen.DisplayRoutine(drivers.GetDriver(s.SegmentCount()), multi, s, q)

	normal = screen.NewBrightness(.1)
	mid = screen.NewBrightness(.2)
	hi = screen.NewBrightness(.6)
	lowoop = screen.NewBounce(.1,.3, time.Second)
	hiwoop = screen.NewBounce(.6,1, time.Second)

	x, y := 3, 4
	writeScreen(s, x, y)

	loop: for {
		select {
			case e := <-events:
				switch e {
					case quit:
						break loop
					case left:
						x = (x+15)%16
					case down:
						y = (y+1)%16
					case up:
						y = (y+15)%16
					case right:
						x = (x+1)%16
				}
				writeScreen(s, x, y)
		}
	}

	close(q)

}


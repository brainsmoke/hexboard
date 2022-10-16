
package main

import (
	"post6.net/gohexdump/internal/screen"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/drivers"
	"os"
//	"fmt"
	"bufio"
	"flag"
//	"time"
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
	s.WriteHexField("S", 7*16+2)
	s.WriteHexField("Y", 7*16+3)
	s.WriteHexField("S", 7*16+4)
	s.WriteHexField("T", 7*16+5)
	s.WriteHexField("E", 7*16+6)
	s.WriteHexField("M", 7*16+7)

	s.WriteHexField("F ", 7*16+8)
	s.WriteHexField("A ", 7*16+9)
	s.WriteHexField("I ", 7*16+10)
	s.WriteHexField("L ", 7*16+11)
	s.WriteHexField("U ", 7*16+12)
	s.WriteHexField("R ", 7*16+13)
	s.WriteHexField("E ", 7*16+14)

	multi, screenChan := screen.NewMultiScreen()

	rippleCursor := screen.NewRippleCursor(.25, nil, s)
	filters := []screen.Filter { rippleCursor, screen.DefaultGamma(), screen.NewAfterGlowFilter(.96) }
	//filters := []screen.Filter { rippleCursor, screen.NewAfterGlowFilter(.8)  }

	screenChan <- screen.NewFilterScreen(s, filters)

	q := make(chan bool)
	events := make(chan int)

	go cmdHandler(os.Stdin, events)

	go screen.DisplayRoutine(drivers.GetDriver(s.SegmentCount()), multi, s, q)

	x, y := 62,9
	rippleCursor.SetCursor(x, y)

	loop: for {
		select {
			case e := <-events:
				switch e {
					case quit:
						break loop
					case left:
						x, y = s.LeftWrap(x, y)
					case down:
						x, y = s.DownWrap(x, y)
					case up:
						x, y = s.UpWrap(x, y)
					case right:
						x, y = s.RightWrap(x, y)
				}
				rippleCursor.SetCursor(x, y)
		}
	}

	close(q)

}


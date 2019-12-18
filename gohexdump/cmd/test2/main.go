
package main

import (
	"post6.net/gohexdump/internal/screen"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/drivers"
	"os"
	"fmt"
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

func writeScreen(s screen.HexScreen) {
	s.Hold()

	normal := screen.NewBrightness(.05)
	mid := screen.NewBrightness(.1)

	for xi:= 0; xi<16; xi++ {

		s.SetStyle(mid)
		s.WriteOffset(fmt.Sprintf("      %02X", xi*16), xi)
		s.WriteTitle(fmt.Sprintf("%X", xi), 5+8+1+xi*3+4*(xi/8))

		for yi:= 0; yi<16; yi++ {
			s.SetStyle(normal)
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

const (
	U = (1<<6) | (1<<7)
	D = (1<<6) | (1<<7)
	R = (1<<1) | (1<<2)
	L = (1<<4) | (1<<5)
	RU = (1<<2) | U
	RD = (1<<1) | D
	LU = (1<<4) | U
	LD = (1<<5) | D
)

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
	writeScreen(s)
	hi := screen.NewBrightness(.4)
	s.SetStyle(hi)
	s.WriteRawHexField([]font.Glyph{LU, U}, 6*16+1)
	s.WriteRawHexField([]font.Glyph{L, 0}, 7*16+1)
	s.WriteRawHexField([]font.Glyph{LD, D}, 8*16+1)
	s.WriteHexField("S ", 7*16+2)
	s.WriteHexField("Y ", 7*16+3)
	s.WriteHexField("S ", 7*16+4)
	s.WriteHexField("T ", 7*16+5)
	s.WriteHexField("E ", 7*16+6)
	s.WriteHexField("M ", 7*16+7)

	s.WriteHexField("F ", 7*16+8)
	s.WriteHexField("A ", 7*16+9)
	s.WriteHexField("I ", 7*16+10)
	s.WriteHexField("L ", 7*16+11)
	s.WriteHexField("U ", 7*16+12)
	s.WriteHexField("R ", 7*16+13)
	s.WriteHexField("E ", 7*16+14)
	s.WriteRawHexField([]font.Glyph{U,RU}, 6*16+15)
	s.WriteRawHexField([]font.Glyph{0,R}, 7*16+15)
	s.WriteRawHexField([]font.Glyph{D,RD}, 8*16+15)
	for i:=2; i<15; i++ {
		s.WriteRawHexField([]font.Glyph{U,U}, 6*16+i)
		s.WriteRawHexField([]font.Glyph{D,D}, 8*16+i)
	}


	multi, screenChan := screen.NewMultiScreen()

	rippleCursor := screen.NewRippleCursor()
	filters := []screen.Filter { rippleCursor, screen.NewAfterGlowFilter(.8) }

	screenChan <- screen.NewFilterScreen(s, filters)

	q := make(chan bool)
	events := make(chan int)

	go cmdHandler(os.Stdin, events)

	go screen.DisplayRoutine(drivers.GetDriver(960*16), multi, q)

	x, y := 8+15*2, 7
	rippleCursor.SetCursor(64 + x + y*56)

	loop: for {
		select {
			case e := <-events:
				switch e {
					case quit:
						break loop
					case left:
						x = (x+55)%56
					case down:
						y = (y+1)%16
					case up:
						y = (y+15)%16
					case right:
						x = (x+1)%56
				}
				rippleCursor.SetCursor(64 + x + y*56)
		}
	}

	close(q)

}


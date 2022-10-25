
package main

import (
	"post6.net/gohexdump/internal/screen"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/drivers"
	"post6.net/gohexdump/internal/util/keys"
	"os"
	"fmt"
	"flag"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func terminal(ch <-chan byte, s screen.TextScreen, cursor screen.Cursor) {

	var err error
	x, y := 62,9
	cursor.SetCursor(x, y)

	loop: for {
		select {
			case key, ok := <-ch:
				if !ok {
					break loop
				}
				switch key {
					case keys.Enter:
						x, y, err = s.Down(0, y)
					case keys.Backspace:

						x, y, err = s.Previous(x, y)

//						if err == nil {
//							s.WriteAt(" ", x, y)
//						}
//						err = nil

					case keys.Up:
						x, y, _ = s.Up(x, y)
					case keys.Down:
						x, y, _ = s.Down(x, y)
					case keys.Left:
						x, y, _ = s.Left(x, y)
					case keys.Right:
						x, y, _ = s.Right(x, y)
					default:
//						if key < 32 || key > 126 { print(" (",key,") ") }

//						c := strings.ToUpper(string(key))
//						x, y, err = s.WriteAt(c, x, y)
						x, y, err = s.Next(x, y)
				}
				if err != nil {
					s.Scroll(0,1)
					x, y, err = 0, y, nil
				}
				cursor.SetCursor(x, y)
		}
	}
}

func writeScreen(s screen.HexScreen) {
	s.Hold()

	normal := screen.NewBrightness(.1)
	mid := screen.NewBrightness(.2)

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

	cursor := screen.NewRippleCursor(1, .5, nil, nil, s)
	filters := []screen.Filter { cursor, screen.DefaultGamma(), screen.NewAfterGlowFilter(.85) }

	screenChan <- screen.NewFilterScreen(s, filters)

	q := make(chan bool)
	ch := make(chan byte)


	go terminal(ch, s, cursor)

	go func(){
	keys.Raw(os.Stdin, ch)
	screenChan <- screen.NewExitScreen(.5)
	}()

	screen.DisplayRoutine(drivers.GetDriver(s.SegmentCount()), multi, s, q)
//	close(q)

}



package main

import (
	"post6.net/gohexdump/internal/screen"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/drivers"
	"post6.net/gohexdump/internal/util/keys"

	"os"
//	"fmt"
//	"bufio"
	"flag"
//	"time"
	"strings"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func transform(pos screen.Vector2) screen.Vector2 {
	return screen.Vector2 { X: pos.X, Y: pos.Y }
}

func terminal(ch <-chan byte, s screen.TextScreen, cursor screen.Cursor) {

	var err error
	x, y := 0,0
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

						if err == nil {
							s.WriteAt(" ", x, y)
						}
						err = nil

					case keys.Up:
						x, y, _ = s.Up(x, y)
					case keys.Down:
						x, y, _ = s.Down(x, y)
					case keys.Left:
						x, y, _ = s.Left(x, y)
					case keys.Right:
						x, y, _ = s.Right(x, y)
					default:
						if key < 32 || key > 126 { print(" (",key,") ") }

						c := strings.ToUpper(string(key))
						x, y, err = s.WriteAt(c, x, y)
				}
				if err != nil {
					s.Scroll(0,1)
					x, y, err = 0, y, nil
				}
				cursor.SetCursor(x, y)
		}
	}
}


func main() {

	flag.Parse()
	args := flag.Args()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	conf := screen.Configuration{
		{ 0, 0, screen.HorizontalPanel },
		{ 0, 1, screen.HorizontalPanel },
		{ 0, 2, screen.HorizontalPanel },
		{ 0, 3, screen.HorizontalPanel },
	}

	q := make(chan bool)
	ch := make(chan byte)

	s := screen.NewTextScreen(conf)

	s.SetFont(font.GetFont())
	s.SetStyle(screen.NewBrightness(1))

	multi, screenChan := screen.NewMultiScreen()

	cursor := screen.NewCursor(1, s)
//	cursor := screen.NewRippleCursor(1, .5, nil, s)
	filters := []screen.Filter { cursor, screen.DefaultGamma(), screen.NewAfterGlowFilter(.85) }

	screenChan <- screen.NewFilterScreen(s, filters)

	go terminal(ch, s, cursor)

	go func(){
	for i,s := range args {
		if i != 0 {
			ch <- ' '
		}
		for _, c := range s {
			ch <- byte(c)
		}
	}

	keys.Raw(os.Stdin, ch)
	screenChan <- screen.NewExitScreen(.5)
	}()

//	close(q)

	screen.DisplayRoutine(drivers.GetDriver(s.SegmentCount()), multi, s, q)
}


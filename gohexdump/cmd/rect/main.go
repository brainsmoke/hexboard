
package main

import (
	"post6.net/gohexdump/internal/screen"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/drivers"
	"os"
//	"fmt"
//	"bufio"
	"flag"
//	"time"
	"strings"
	"runtime/pprof"
	"golang.org/x/term"
)

const (
	ctrlC = 3
	ctrlD = 4
	enter = 13
	escape = 27
	backspace = 127
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func rawKeys(file *os.File, keys chan<- rune) {

	fd := int(file.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}

	for {
		var buf [1]byte
		_, err := file.Read(buf[0:1])
		c := buf[0]
		if err != nil || c == ctrlD || c == ctrlC {
			break
		}
		keys <- rune(c)
	}

	term.Restore(fd, oldState)
	close(keys)
}

func transform(pos screen.Vector2) screen.Vector2 {
	return screen.Vector2 { X: pos.X, Y: pos.Y }
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

	conf := screen.Configuration{
		{ 0, 0, screen.HorizontalPanel },
		{ 0, 1, screen.HorizontalPanel },
	}

	q := make(chan bool)
	keys := make(chan rune)

	s := screen.NewTextScreen(conf)

	s.SetFont(font.GetFont())
	s.SetStyle(screen.NewBrightness(1))

	multi, screenChan := screen.NewMultiScreen()

	rippleCursor := screen.NewRippleCursor(1, transform, s)
	filters := []screen.Filter { rippleCursor, screen.DefaultGamma(), screen.NewAfterGlowFilter(.8) }

	screenChan <- screen.NewFilterScreen(s, filters)

	go screen.DisplayRoutine(drivers.GetDriver(s.SegmentCount()), multi, s, q)

	go rawKeys(os.Stdin, keys)

	var err error
	x, y := 0,0
	rippleCursor.SetCursor(x, y)

	loop: for {
		select {
			case key, ok := <-keys:
				if !ok {
					break loop
				}
				switch key {
					case enter:
						x, y, err = s.Down(0, y)
					case backspace:

						x, y, err = s.Previous(x, y)

						if err == nil {
							s.WriteAt(" ", x, y)
						}
						err = nil

					default:
print(" (",key,") ")
						c := strings.ToUpper(string(key))
						x, y, err = s.WriteAt(c, x, y)
				}
				if err != nil {
					s.Scroll(0,1)
					x, y = 0, y
				}
				rippleCursor.SetCursor(x, y)
		}
	}

	close(q)
}


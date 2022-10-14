
package main

import (
	"post6.net/gohexdump/internal/screen"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/drivers"
	"os"
	"bufio"
	"flag"
)

const (
	nop = iota
	quit
)

func cmdHandler(file *os.File, events chan<- int) {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		switch scanner.Text() {
		case "quit":
			events <- quit
			break
		}
	}

	close(events)
}

func main() {

	flag.Parse()

	s := screen.NewHexScreen()
	info := s.Info()

	s.SetFont(font.GetFont())
/*
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
*/
	multi, screenChan := screen.NewMultiScreen()

	filters := []screen.Filter { screen.NewRaindropFilter(info), screen.DefaultGamma() }

	screenChan <- screen.NewFilterScreen(s, filters)

	q := make(chan bool)
	events := make(chan int)

	go cmdHandler(os.Stdin, events)

	go screen.DisplayRoutine(drivers.GetDriver(info.Size*16), multi, info, q)

	loop: for {
		select {
			case e := <-events:
				switch e {
					case quit:
						break loop
				}
		}
	}

	close(q)

}


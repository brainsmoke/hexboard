
package keys

import (
	"golang.org/x/term"
	"os"
)

const (
	Home  = 1
	Left  = 2
	CtrlC = 3
	CtrlD = 4
	End   = 5
	Right = 6
	Backspace = 8
	Enter  = 13
	Escape = 27
	Down   = 14
	Up     = 16
	otherBackspace = 127
)

func getC(file *os.File) (byte, error) {
	var buf [1]byte
	_, err := file.Read(buf[0:1])
	return buf[0], err
}

func Raw(file *os.File, keys chan<- byte) {

	defer close(keys)

	fd := int(file.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	for {
		c, err := getC(file)
		if err != nil || c == CtrlD || c == CtrlC {
			break
		}

		switch c {
			case Escape:
				cmd, err := getC(file)
				if err != nil {
					break
				}
				if cmd == byte('[') {
					cmd2, err := getC(file)
					if err != nil {
						break
					}
					switch cmd2 {
						case byte('A'):
							keys <- Up
						case byte('B'):
							keys <- Down
						case byte('C'):
							keys <- Right
						case byte('D'):
							keys <- Left
						case byte('E'):
							keys <- Home
						case byte('F'):
							keys <- End
						default:
							keys <- c
							keys <- cmd
							keys <- cmd2
					}
				} else {
					keys <- c
					keys <- cmd
				}
			case otherBackspace:
				keys <- Backspace
			default:
				keys <- c
		}
	}
}


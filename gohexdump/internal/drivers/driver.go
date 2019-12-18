package drivers

import (
	"flag"
	"os"
//	"math"
)

var serialDevice string

type Driver struct {
	file *os.File
	buf []byte
}

func init() {

	flag.StringVar(&serialDevice, "device", "/dev/ttyACM0", "serial output device")
}

func GetDriver(size int) *Driver {

	var err error

	d := new(Driver)
	d.buf = make([]byte, size+1)
	d.buf[size] = 0xfe

	d.file, err = os.OpenFile(serialDevice, os.O_RDWR, 0)
	SetBaudrate(d.file, 12000000)
	SetBinary(d.file)

	if err != nil {
		panic("could not open serial device")
	}

	d.file.Write([]byte{0xff}) // Discard frame

	return d
}

func (d *Driver) Write(data []float64) (int, error) {
	l := len(data)
	if l > len(d.buf)-1 {
		l = len(d.buf)-1
	}

	frame := d.buf[:l]
	for i := range(frame) {
		f := 0xfd*data[i]
		if f < 0 {
			f = 0
		} else if f > 0xfd {
			f = 0xfd
		}
		frame[i] = byte( f ) //math.Min(0xfd, math.Max(0, 0xfd * data[i])) )
	}
	return d.file.Write(d.buf)
}

func (d *Driver) Close() error {
	return d.file.Close()
}


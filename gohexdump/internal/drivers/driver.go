package drivers

import (
	"flag"
	"os"
	"post6.net/gohexdump/internal/util/clip"
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

	size *= 2
	d := new(Driver)
	d.buf = make([]byte, size+4)
	d.buf[size] = 0xff
	d.buf[size+1] = 0xff
	d.buf[size+2] = 0xff
	d.buf[size+3] = 0xf0

	d.file, err = os.OpenFile(serialDevice, os.O_RDWR, 0)
	SetBaudrate(d.file, 480000000)
	SetBinary(d.file)

	if err != nil {
		panic("could not open serial device")
	}

	d.file.Write([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xf0}) // Discard frame

	return d
}

func (d *Driver) Write(data []float64) (int, error) {
	l := len(data)
	if l > (len(d.buf)-4)/2 {
		l = (len(d.buf)-4)/2
	}

	for i:=0; i<l ;i++ {
		v := clip.FloatToUintRange(data[i]*0xff00, 0, 0xff00)
		d.buf[i*2  ] = byte(v & 0xff)
		d.buf[i*2+1] = byte(v >> 8)
	}
	return d.file.Write(d.buf)
}

func (d *Driver) Close() error {
	return d.file.Close()
}


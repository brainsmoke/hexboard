package drivers

/*

#include <sys/ioctl.h>
#include <asm/termbits.h>

int set_baudrate(int fd, unsigned int baudrate)
{
	struct termios2 tio;

	ioctl(fd, TCGETS2, &tio);
	tio.c_cflag &= ~(CBAUD);
	tio.c_cflag |= BOTHER;
	tio.c_ispeed = baudrate;
	tio.c_ospeed = baudrate;

	return ioctl(fd, TCSETS2, &tio);
}


int set_binary(int fd)
{
	struct termios2 tio;

	ioctl(fd, TCGETS2, &tio);
	tio.c_iflag &= ~(ICRNL|BRKINT);
	tio.c_oflag &= ~(OPOST|ONLCR|ECHO);
	tio.c_lflag &= ~(ICANON|ISIG);

	return ioctl(fd, TCSETS2, &tio);
}

*/
import "C"

import (
	"os"
)

func SetBinary(file *os.File) (int, error) {

	var cint_ok C.int
	var err error
	cint_ok, err = C.set_binary(C.int(file.Fd()))
	return int(cint_ok), err
}

func SetBaudrate(file *os.File, baudrate uint) (int, error) {

	var cint_ok C.int
	var err error
	cint_ok, err = C.set_baudrate(C.int(file.Fd()), C.uint(baudrate))
	return int(cint_ok), err
}


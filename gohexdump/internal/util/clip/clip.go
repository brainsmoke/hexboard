package clip

func FloatToByte(f float64) byte {

	if f > 255. {
		return 255
	} else if f < 1 {
		return 0
	} else {
		return byte(f)
	}
}

func FloatToUintRange(f float64, min, max uint) uint {

	u := uint(f)

	if u >= max {
		return max
	} else if u <= min {
		return min
	} else {
		return u
	}
}

func FloatBetween(f, min, max float64) float64 {
	if f < min {
		return min
	} else if f < max {
		return f
	} else {
		return max
	}
}

func IntToByte(i int) byte {

	if i > 255 {
		return 255
	} else if i < 1 {
		return 0
	} else {
		return byte(i)
	}
}

func AddBytes(a, b byte) byte {
	if int(a)+int(b) > 255 {
		return 255
	} else {
		return a + b
	}
}

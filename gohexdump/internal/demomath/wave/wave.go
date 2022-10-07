
package wave

import (
	"math"
)

const Fps = 60
const tableBits = 11
const tableSize = 1<<tableBits
const cycleBits = 39
const Multiplier = (  (1000000000<<cycleBits) - ( (1000000000<<cycleBits) % Fps ) )/Fps
const Shift = cycleBits-tableBits
const Mask = tableSize-1

var Wave []float64

func init() {
	Wave = make([]float64, tableSize)
	for i := range Wave {
		t := float64(i)*2*math.Pi/float64(tableSize)
		Wave[i] = math.Min(1, math.Max(0, .5*(1+math.Sin(t))))
	}
}


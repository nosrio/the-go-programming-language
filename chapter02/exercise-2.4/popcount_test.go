// Wr ite a version of PopCount that counts bits by shif ting its argument through 64
// bit position s, testing the rig htmost bit each time. Compare its per for mance to the table-
// lo oku p version.

package popcount

import (
	"testing"
)

func PopCountTableLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(1))])
	}
	return sum
}

func PopCountShiftValue(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; 1 < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>=1
	}
	return count
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkTable(b *testing.B) {
	bench(b, PopCountTable)
}

func BenchmarkShiftValue(b *testing.B) {
	bench(b, PopCountShiftValue)
}
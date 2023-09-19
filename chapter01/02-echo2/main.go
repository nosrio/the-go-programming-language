// Anot her for m of the for lo op iterates over a ra nge of values fro m a dat a type like a st ring or a
// slice.

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}
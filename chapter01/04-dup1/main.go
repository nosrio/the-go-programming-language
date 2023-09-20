// The firs t version of dup pr ints each line that app ears more than once in the stand ard inp ut,
// preceded by its count. This program introduces the if st atement, the map data typ e, and the
// bufio package.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
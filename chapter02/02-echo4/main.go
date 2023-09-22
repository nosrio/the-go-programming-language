// To illustrate, this var iat ion
// on the earlier echo command takes two opt ion al flags: -n caus es echo to omit the trai ling
// ne wline that wou ld nor mal ly be print ed, and -s sep caus es it to sep arate the out put argu-
// ments by the contents of the str ing sep instead of the defau lt single space.

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
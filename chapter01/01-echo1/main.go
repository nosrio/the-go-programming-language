// Here’s an imp lementation of the Unix echo command, which prints its command-line argu-
// ments on a single line.

package main

import (
	"fmt"
	"os"
)

func main(){
	var s, sep string
	sep =" "
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}
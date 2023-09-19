// Mo dif y the echo prog ram to print the index and value of each of its arguments,
// on e per line.

package main

import (
	"fmt"
	"os"
)

func main() {

	for index, value := range os.Args {
		fmt.Printf("On position %d the argument value is %s\n", index, value)
	}
}
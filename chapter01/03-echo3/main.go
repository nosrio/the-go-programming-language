// If the amount of dat a invo l ved is large , this cou ld be costly. A simpler and more efficient
// solut ion wou ld be to use the Join func tion fro m the strings package

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
}
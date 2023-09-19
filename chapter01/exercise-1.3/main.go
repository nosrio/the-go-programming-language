// Experiment to measure the dif ference in running time bet ween our pot ent ial ly
// inefficient versions and the one that uses strings.Join .

// go run main.go esto es una simple prueba de cuanto puede llegar a tardar estas funciones
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep = "", ""
	
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Using for took %v\n", elapsed)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:]," "))
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Using joins took %v\n", elapsed)

}
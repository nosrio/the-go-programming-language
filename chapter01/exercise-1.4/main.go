// The next version of dup can read fro m the stand ard inp ut or handle a list of file names,
// using os.Open to open each one

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for file, lines := range counts {
		for line, n  := range lines {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, file)
			}
		} 
	}
}

func countLines(f *os.File, counts map[string]map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[f.Name()] == nil {
			counts[f.Name()] = make(map[string]int)
		}
		counts[f.Name()][input.Text()]++
	}
}
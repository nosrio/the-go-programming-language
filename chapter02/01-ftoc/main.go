// The function fToC below enc apsu-
// lates the temperature conversion log ic so that it is define d on ly once but may be used fro m
// mu ltiple places. Here main calls it twice, using the values of two dif ferent local cons tants:

package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64{
	return (f - 32) * 5 / 9
}
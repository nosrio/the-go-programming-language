// To illustrate typ e de clarat ions, let’s tur n the dif ferent temperature scales into dif ferent typ es:

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC 	Celsius = -273.15
	FreezingC 		Celsius = 0
	BolingC			Celsius = 100 
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
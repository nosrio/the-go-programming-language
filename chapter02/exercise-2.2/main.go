// Wr ite a general-pur pos e unit-conv ersion program analogou s to cf that reads
// numb ers fro m its command-line arguments or fro m the stand ard inp ut if there are no argu-
// ments, and converts each number int o units like temperature in Cel siu s and Fahren heit,
// lengt h in feet and meters, weig ht in pound s and kilog rams, and the like.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Measurement interface {
	String() string
}

type Distance struct {
	meters float64
}

func FromFeet(f float64) Distance {
	return Distance{f * 0.3048}
}

func FromMeters(f float64) Distance {
	return Distance{f}
}

func (d Distance) String() string {
	return fmt.Sprintf("%.3gm = %.3gft", d.Meters(), d.Feet())
}

func (d Distance) Meters() float64 {
	return d.meters
}

func (d Distance) Feet() float64 {
	return d.meters / 0.3048
}

type Temperature struct {
	kelvin float64
}

func FromCelsius(f float64) Temperature {
	return Temperature{f - 273.15}
}

func FromFarenheit(f float64) Temperature {
	return Temperature{((f - 32) * 5 / 9)- 273.15}
}

func FromKelvin(f float64) Temperature {
	return Temperature{f}
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.3gk = %.3gc = %.3gf", t.Kelvin(), t.Celsius(), t.Fahrenheit())
}

func (t Temperature) Kelvin() float64 {
	return t.kelvin
}

func (t Temperature) Celsius() float64 {
	return t.kelvin + 273.15
}

func (t Temperature) Fahrenheit() float64 {
	return ( t.Celsius() * 9 / 5) + 32
}

type Weight struct {
	kilograms float64
}

func FromKilograms(f float64) Weight {
	return Weight{f}
}

func FromPounds(f float64) Weight {
	return Weight{f * 2.205	}
}

func (w Weight) String() string {
	return fmt.Sprintf("%.3gkg = %.3g£", w.Kilograms(), w.Pounds())
}

func (w Weight) Kilograms() float64 {
	return w.kilograms
}

func (w Weight) Pounds() float64 {
	return w.kilograms / 2.205
}

func convert(f float64, unit string) (Measurement, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "m":
		return FromMeters(f), nil
	case "ft":
		return FromFeet(f), nil
	case "c":
		return FromCelsius(f), nil
	case "f":
		return FromFarenheit(f), nil
	case "k":
		return FromKelvin(f), nil
	case "kg":
		return FromKilograms(f), nil
	case "£":
		return FromPounds(f), nil
	default:
		return nil, fmt.Errorf("unexpected unit %v", unit)
	}
}

func showConvert(s string) {
	re := regexp.MustCompile((`-?(\d+(?:\.\d+)?)(\w+)`))
	match := re.FindStringSubmatch(s)
	if match == nil {
		log.Fatalf("Expecting <number><unit>, got %q", s)
	}
	f, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		log.Fatalf("%v is not a number", match[1])
	}
	if match[2] == "" {
		log.Fatalf("No unit specified")
	}
	unit := match[2]
	m, err := convert(f, unit)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(m)
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			showConvert(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			showConvert(scan.Text())
		}
	}
}
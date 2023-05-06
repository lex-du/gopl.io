// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

type Kelvin float64

func CToK(c tempconv.Celsius) Kelvin { return Kelvin(c + 273.15) }
func KToC(k Kelvin) tempconv.Celsius { return tempconv.Celsius(k - 273.15) }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// !+
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//!-celsiusFlag

//!+CelsiusFlag

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

//!-CelsiusFlag
//!-

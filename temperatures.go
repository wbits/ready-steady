package main

import "fmt"

type temperature float64
type celsius temperature
type kelvin temperature
type fahrenheit temperature
type rowWriter func(degrees celsius) (string, string)

const ZeroKelvinCelsius = - 273.15

func drawTable(start int, end int, writeRow rowWriter)  {
	table := ""
	for i := start; i <= end; i++ {
		cRow, fRow := writeRow(celsius(i))
		table += fmt.Sprintf("|%-8v|%-8v|\n", cRow, fRow)
	}

	fmt.Print(table)
}

func drawRow(c celsius) (string, string) {
	return fmt.Sprintf("%8.2f", c), fmt.Sprintf("%8.2f", c.fahrenheit())
}

func main() {
	drawTable(-40, 100, drawRow)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c - ZeroKelvinCelsius)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (k kelvin) celsius() celsius {
	return celsius(k + ZeroKelvinCelsius)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

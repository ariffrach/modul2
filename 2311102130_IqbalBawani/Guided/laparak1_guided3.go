package main

import "fmt"

func main() {
	var temperaturCelcius float64

	fmt.Println("Masukkan temperatur dalam Celcius:")
	fmt.Scanln(&temperaturCelcius)

	temperaturFahrenheit := (temperaturCelcius * 9 / 5) + 32
	temperaturKelvin := temperaturCelcius + 273.15
	temperaturReamur := temperaturCelcius * 4 / 5

	fmt.Printf("Temperatur dalam Fahrenheit: %.2f\n", temperaturFahrenheit)
	fmt.Printf("Temperatur dalam Kelvin: %.2f\n", temperaturKelvin)
	fmt.Printf("Temperatur dalam Reamur: %.2f\n", temperaturReamur)
}

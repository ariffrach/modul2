package main

import (
        "fmt"
        "math"
)

func main() {
        var celsius float64

        fmt.Print("Temperatur Celsius: ")
        fmt.Scanln(&celsius)

        fahrenheit := math.Round((celsius * 9 / 5) + 32)

        reamur := math.Round(celsius * 4 / 5)

        kelvin := math.Round(celsius + 273.15)

        fmt.Printf("Derajat Fahrenheit: %d\n", int(fahrenheit))
        fmt.Printf("Derajat Reamur: %d\n", int(reamur))
        fmt.Printf("Derajat Kelvin: %d\n", int(kelvin))
}

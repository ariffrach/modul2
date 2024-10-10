package main

import (
        "fmt"
        "math"
)

func hitungAkarDua(k int) float64 {
        var hasil float64 = 1.0
        for i := 0; i <= k; i++ {
                pembilang := math.Pow(float64(4*i+2), 2)
                penyebut := float64(4*i+1) * float64(4*i+3)
                hasil *= pembilang / penyebut
        }
        return hasil
}

func main() {
        var k int

        fmt.Print("Nilai K= ")
        fmt.Scan(&k)

        akarDua := hitungAkarDua(k)
        fmt.Printf("Nilai akar 2= %.10f\n", akarDua)
}
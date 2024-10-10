package main

import "fmt"

func main() {
        var bilangan int

        fmt.Print("Bilangan: ")
        fmt.Scan(&bilangan)

        fmt.Print("Faktor: ")
        for i := 1; i <= bilangan; i++ {
                if bilangan%i == 0 {
                        fmt.Printf("%d ", i)
                }
        }
        fmt.Println()

        var jumlahFaktor int = 0
        for i := 1; i <= bilangan; i++ {
                if bilangan%i == 0 {
                        jumlahFaktor++
                }
        }

        if jumlahFaktor == 2 {
                fmt.Println("Prima: true")
        } else {
                fmt.Println("Prima: false")
        }
}
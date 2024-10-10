package main

import "fmt"

func main() {
        var n int
        var bunga, pita string

        fmt.Print("Masukkan jumlah bunga: ")
        fmt.Scanln(&n)

        for i := 1; i <= n; i++ {
                fmt.Printf("Bunga %d: ", i)
                fmt.Scanln(&bunga)

                if bunga == "SELESAI" {
                        break
                }

                pita += bunga + "-"
        }

        if len(pita) > 0 {
                pita = pita[:len(pita)-1]
        }

        fmt.Println("Pita:", pita)
}
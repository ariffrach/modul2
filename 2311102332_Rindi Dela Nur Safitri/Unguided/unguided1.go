package main

import (
        "bufio"
        "fmt"
        "os"
        "strings"
)

func main() {
        correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

        reader := bufio.NewReader(os.Stdin)
        success := true

        for i := 1; i <= 5; i++ {
                fmt.Printf("Percobaan %d: ", i)

                input, _ := reader.ReadString('\n')
                input = strings.TrimSpace(input)
                colors := strings.Split(input, " ")

                if len(colors) != len(correctOrder) {
                        fmt.Println("Jumlah warna tidak sesuai")
                        continue
                }

                for j := 0; j < len(colors); j++ {
                        if colors[j] != correctOrder[j] {
                                success = false
                                break
                        }
                }

                if !success {
                        break
                }
        }

        if success {
                fmt.Println("BERHASIL: TRUE")
        } else {
                fmt.Println("BERHASIL: FALSE")
        }
}

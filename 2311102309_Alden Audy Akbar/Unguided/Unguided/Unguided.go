package main

// 2311102309
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func cekPercobaan(warnaPercobaan [][]string) bool {
    warnaIdeal := []string{"merah", "kuning", "hijau", "ungu"}
    for _, percobaan := range warnaPercobaan {
        if !isSame(percobaan, warnaIdeal) {
            return false
        }
    }
    return true
}
func isSame(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func main() {
    var warnaPercobaan [][]string
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Masukkan 5 percobaan (masukkan 4 warna per percobaan):")
    for i := 0; i < 5; i++ {
        fmt.Printf("Percobaan %d: ", i+1)
        scanner.Scan()
        line := scanner.Text()

        percobaan := strings.Fields(line)

        if len(percobaan) != 4 {
            fmt.Println("Input tidak valid, masukkan tepat 4 warna.")
            i--
            continue
        }
        warnaPercobaan = append(warnaPercobaan, percobaan)
    }

    if cekPercobaan(warnaPercobaan) {
        fmt.Println("True")
    } else {
        fmt.Println("False.")
    }
}



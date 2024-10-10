package main

import "fmt"

func main() {
	var tahun int
	fmt.Println("program menentukan false atau true nilai tahun kabisat")
	fmt.Println("input tahun: ")
	fmt.Scanln(&tahun)

	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

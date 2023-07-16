package main

import (
	"fmt"

	"github.com/manuelfx117/DesignPatterns/Creational"
)

func main() {
	for i := 0; i < 3; i++ {
		go Creational.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

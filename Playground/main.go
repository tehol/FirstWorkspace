// Playground project main.go
package main

import (
	"Playground/FirstSubPkg"
	"fmt"
)

func main() {

	total := 0

	for i := 0; i <= 10; i++ {
		total = total + i

	}
	fmt.Printf("1-10 total:%v\n", total)
	fmt.Println(Test())
}

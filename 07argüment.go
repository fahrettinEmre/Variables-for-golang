package main

import (
	"fmt"
	"os"
)

//argüment go run yaparken dışardan data vermemize yarar
func main() {
	args := os.Args[1:]

	fmt.Println("selam", args[0])
	fmt.Println("selam", args[1])

	a, _ := os.ReadFile("./01helloworld") // bir file okur byte ve array döner
	fmt.Println(string(a))                //byte arrayini stringe dönüştürürüz
}

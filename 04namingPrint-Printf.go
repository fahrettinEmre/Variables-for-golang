package main

import "fmt"

func main() {

	// (Print - Println) - Printf
	fmt.Println("merhaba") // LN ifadesi yeni bir satır başlangıcı yapar ve alta geçer yazdırır.
	fmt.Print("merhaba\n")
	fmt.Print("merhaba2")
	fmt.Printf("merhaba3")

	name := "fahrettin emre "

	fmt.Print(name)
	fmt.Println(name)
	fmt.Printf("%v", name) // Printf formatlanmış değişkenler için kullanılır.
	// %v = value yani değer
	// %T = value type
	fmt.Printf("%v %T", name, name)

	x := 100
	fmt.Println("")
	fmt.Printf("%b", x) // %b 2 lik sayı tabanı
	// %d 10 luk sayı tabanı
	// %o 8 lik sayı tabanı
	fmt.Println("")
	name2 := "fahrettin"
	age := 29

	fmt.Printf("Benim adım %v ve ben %v yaşındayım.", name2, age)
}

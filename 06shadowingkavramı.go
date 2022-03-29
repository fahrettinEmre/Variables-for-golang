package main

import "fmt"

// var x = 8 func main burdayı gölgeler . sonuc 8 cıkmaz.

func main() {

	x := 5

	if true {
		x := 10 // shodowin kavramı burada yapılıuyor. Yeniden declare ediliyor.
		// eğer yeniden değer atarsam 2 değeride aynı print eder.
		fmt.Println(x)
	}
	fmt.Println(x)
	// var matrix [4][7]int
	// fmt.Println((matrix))
	// var t =[][]int {{1, 2, 3}, {2, 4, 6}}
	sonuc := data([][]int{{1, 2, 3}, {2, 4, 6}})
	fmt.Println(sonuc)
}

func data(d [][]int) int {
	var toplam int
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			toplam += d[i][j]
		}
	}
	return toplam
}

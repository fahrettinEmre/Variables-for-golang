/* uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32 */

package main

import "fmt"

func main() {

	x := 10 // y=float64(x)
	y := 10.5

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Println(x + int(y)) // burada x ve y aynı veri tipinden olmaz ise toplama işlemi yapılamaz.
	// burada yeni bir değer oluşturuyoruz.Ama asıl veri tipi değişmez aynı kalır.

	// type conversion yaparken type'lar da aynı olmak zorundadır. Örnek string int e dönüşmez.
	ascı()

}

func ascı() {
	num1 := 97
	str1 := string(num1) // ASCI karaktere göre int veri tipini stringe çevirdik ve desimal olarak değer aldı.

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Printf("%v %T\n", str1, str1)

	a := 65
	b := strconv.Iota(a)
	fmt.Printf("%v %T\n", b, b)

}

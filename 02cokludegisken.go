package main

import "fmt"
a:=4
var packVar = "fonksiyon" // burada main func dışında değişken oluşturabiliriz. Scope'lar nerede olduğu çpk öenmli.
// eğer scop dışında farklı fonksiyonda tanımlanmış değişken var ise kabul etmez. Ama package olarak duruyorsa kabul eder.

func main() {

	/*  	var (
		name string = "fed"  //1. burada var içerisine birden fazla değişken atadık.
		age  int    = 40

		weight float32 = 65.6 // birbirlerine yakın değişkenleri grupladık.Örnek fiziksel özellikler gibi.
		hight  int     = 175
	)  */

	/* var name, age, weight, hight = "fed", 40, 65.6, 175 */ // 2. tek seferde declare(bildirmek) ettik ve assign(atama) yaptık.

	name, age, weight, hight := "fed", 40, 65.6, 175 //3. farklı bildirme

	name = "emre " // burada tekrardan declare etmiyoruz. Değişkene yeniden değer atıyoruz.

	//	var name string  // sadece bunu yazdırırsak bize boş bir string verir. Default value yada zero value si boş string dir.
	// var age int       // burda default value 0 dır. Numeric valuelerin zero valuesi 0 dır.
	// var isMarried boll // boolen ifadelerin default valuesi yani zero valuesi false dur.

	fmt.Println(name, age, weight, hight)

}

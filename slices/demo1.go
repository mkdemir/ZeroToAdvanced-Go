package slices

import "fmt"

func Demo1() {
	// slice geliştirilmiş dizilerdir.
	// Dizilerdeki problem biz eleman sayısını verdikten sonra yeni bir eleman ekleyemiyorduk. Yeni bir eleman ekleyebilmemiz için bizim o dizi ile ilgili eleman sayısını attırmamız gerekiyor. Ams arttırdığımız zaman point ettiği değer değişiyor. Yep yeni bizi dizi oluyor ve elemanlarını kayıp ediyor.

	// make -> make fonk. ile bir çok şey yapılabiliyor. Bunlardan biri slices'dır
	isimler := make([]string, 3) // slice parça demek.
	fmt.Println(isimler)
	isimler[0] = "Mustafa"
	isimler[1] = "Kaan"
	isimler[2] = "Demir"
	isimler = append(isimler, "Test")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}

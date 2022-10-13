package pointers

import "fmt"

/*
Pointer'lar günümüzde yeni nesil programlama dillerinde Go'nun dışında mesela C# ve JAVA'da referans tipler olarak karşımıza çıkar. Benzeridir.
*/
/*
func Demo1(sayi int) {
	sayi = sayi + 1
	fmt.Println("Demodaki sayı :", sayi)
}

*/
func Demo1(sayi *int) { // *int değer ziyade onun bellekteki adresi üzerinden çalışacağınızı söylüyorsunuz.
	*sayi = *sayi + 1 // Değeri değilde ilgili o sayının adresini set etmek için biz *sayi koyuyoruz.
	fmt.Println("Demodaki sayı :", *sayi)
}

func Test() {
	numara := 10
	Demo1(&numara)
	// numara ?? 10
}

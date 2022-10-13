package functions // package fmt
import "fmt"

/*
func Topla(sayi1 int, sayi2 int) { // parametre vereceğiz.
	var toplam = sayi1 + sayi2
	fmt.Println("Sonuç : ", toplam)
}

*/

func Topla(sayi1 int, sayi2 int) int { // parametre vereceğiz.
	var toplam = sayi1 + sayi2
	return toplam
}

func SelamVer(kullaniciAdi string) {
	fmt.Println("Merhaba ", kullaniciAdi)
}

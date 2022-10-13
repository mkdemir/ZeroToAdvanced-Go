package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Çift Sayıdır." // Fonksiyon return gördüğünde hayatını sonlandırıyor.
	}

	if sayi%2 != 0 {
		fmt.Println("Tek Sayı Çalıştı")
		return "Tek Sayıdır." // Return, fonksiyonu bitirmeye yarıyor.
	}

	// defer DeferFunc()
	return "Belli Değil"

}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}

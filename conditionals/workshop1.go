package conditionals

import "fmt"

func Demo3() {
	// Üç adet int değişken tanımlayınız.
	// Ekrana en büyük olanı yazdırınız.
	// 10 dk süre : 04:48 Tamamlandı

	/*
		var degisken1 int = 10
		var degisken2 int = 5
		var degisken3 int = 8

		if degisken1 > degisken2 {
			if degisken1 > degisken3 {
				fmt.Println("En büyük değer degisken1: " + fmt.Sprintf("%v", degisken1))
			} else {
				fmt.Println("En büyük değer degisken3: " + fmt.Sprintf("%v", degisken3))
			}
		} else if degisken2 > degisken1 {
			if degisken2 > degisken3 {
				fmt.Println("En büyük değer degisken1: " + fmt.Sprintf("%v", degisken2))
			} else {
				fmt.Println("En büyük değer degisken3: " + fmt.Sprintf("%v", degisken3))
			}

		}
	*/

	var sayi1, sayi2, sayi3 int = 10, 23235, 18

	// buffer oluşturma
	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayı: %v", enBuyuk)
	fmt.Println("En büyü sayı: " + fmt.Sprintf("%v", enBuyuk))

}

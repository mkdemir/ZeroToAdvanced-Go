package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {

	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.") // Hata fırlatma yapıyor.
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı giriniz", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı giriniz", nil
	}

	return "Bildiniz", nil // Hata dönmek nil yani hata yok demek
}

func Demo2() {

	mesaj, hata := TahminEt(5022)
	fmt.Println(mesaj)
	fmt.Println(hata)
	// fmt.Println(TahminEt(101))
	// fmt.Println(TahminEt(-10))
}

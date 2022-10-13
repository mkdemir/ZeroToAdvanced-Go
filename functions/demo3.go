package functions

func ToplaVariadic(sayilar ...int) int { // Sayıları dizi olarak ele alıyor.
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}

	return toplam
}

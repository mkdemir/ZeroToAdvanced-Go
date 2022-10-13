package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"} // make yerine bir tane array tanımlayıp slices oluşturabiliriz.
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler)) // kopyasını oluşturmak istediğimizde yep yeni bir dizi oluşturma.
	fmt.Println(sehirlerKopya)

	copy(sehirlerKopya, sehirler) // Bunu normal array^ler içinde yapabilirsiniz.

	fmt.Println(sehirlerKopya)

	sehirler = append(sehirler, "Adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3]) // 1'den itibaren 3'e kadar demek bu (Yani bana şehirlerin içerisindeki elemanların 3. indisinden başlayıp 4'e kadar.)
	fmt.Println(sehirler[1:])  // 1. indisten başla sonuna kadar al.

	fmt.Println(sehirler[:2])
}

package for_range

import "fmt"

// Range yapısı ile dizi temelli yapıları biz iterate edebiliyoruz. Elemanları tek tek dolaşbiliyoruz.
// Maps range kullanmak daha mantıklı
// Başka dillerdeki for each döngülerine benzetilebilinir.
func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for _, sehir := range sehirler { // range sehirler, şehirleri benim için gez demek
		fmt.Println(sehir)
	}
	for i, sehir := range sehirler { // range sehirler, şehirleri benim için gez demek
		fmt.Print(i)
		fmt.Println(sehir)
	}

}

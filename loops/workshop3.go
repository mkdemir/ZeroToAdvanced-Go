package loops

import "fmt"

// 04:46
// 220 ve 284 arkadaş sayılardır
// 10 ve 65 arkadaş sayı değildir. -> 220 : 1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110 = 284 284 : 1 + 2 + 4 + 71 + 142 = 220
func Demo5() {

	sayi1 := 17296 //220
	sayi2 := 18416 //284
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 { //scope
			toplam1 = toplam1 + i
		}
	}

	for i := 1; i < sayi2; i++ { //scope
		if sayi2%i == 0 {
			toplam2 = toplam2 + i
		}

	}

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayılardır.", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayılar değildir.", sayi1, sayi2)
	}

}

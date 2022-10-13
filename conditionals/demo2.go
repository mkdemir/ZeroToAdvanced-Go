package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki para yetersiz")
	} else if cekilmekIstenen == hesap { // syntax => Kodlamadaki yazım şeklidir. (else if == ama)
		fmt.Print("Paranız hazırlanıyor")
		fmt.Print("Dikkat, hesapta para kalmadı")
	} else {
		fmt.Print("Paranız hazırlanıyor")
	}
	// İç içe If'lerin kullanımı

	/*
		if sart() {
			if sart1 {

			}
			//kod

		}
	*/

}

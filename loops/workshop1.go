package loops

import "fmt"

/*
	Oyun yazacağız.

	* 1'de 100 kadar 1 sayı söylersiniz. Aklınızda tutarsınız. Ve karşınızdaki kişi o sayıyı tahmin eder.

	* 70 tahmin 20 yukarı dersiniz.

	1. Kaç tahminde bildiniz.
		* Örn: 3 tahmin : Süper
			* 1-3: Süper  4-6: Güzel > 6 Fenal Değil

*/

/*
func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	// 05:14

	fmt.Print("Bir sayı tahmin ediniz: ")
	fmt.Scanln(&tahminEdilenSayi) // Kullanıcıdan bilgi alma

	// fmt.Println(tahminEdilenSayi)

	/*
		if aklimdakiSayi == tahminEdilenSayi {
			fmt.Println("Evet sayıyı buldunuz.")
		} else if aklimdakiSayi < tahminEdilenSayi {
			fmt.Println("Aklımdaki sayı tahmin edilen sayıda küçüktür.")
		} else {
			fmt.Println("Aklımdaki sayı tahmin edilen sayıdan büyüktür.")
		}
*/
/*
func Demo3() {
	aklimdakiSayi := 80 //manuel olarak yazdık bunu rang ile de yazabilirdik.
	tahminEdilenSayi := 0
	sayac := 0
	text := "metin"
	// 08:36

	fmt.Print("Bir sayı tahmin ediniz: ")
	fmt.Scanln(&tahminEdilenSayi) // satırı yara  (Burada bir pointer söz konusu)
	for aklimdakiSayi != tahminEdilenSayi {

		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayı giriniz")
		}
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayı giriniz")
		}
		fmt.Scanln(&tahminEdilenSayi)

		sayac = sayac + 1
	}

	if sayac > 6 {
		text = "Fena Değil"
	} else if sayac >= 4 {
		text = "Güzel"
	} else {
		text = "Süper"
	}

	fmt.Println("Bravo bildiniz " + fmt.Sprintf("%v", text))
*/

func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Print("Bir sayı tahmin ediniz: ")
	fmt.Scanln(&tahminEdilenSayi)

	tahminSayisi = tahminSayisi + 1
	for aklimdakiSayi != tahminEdilenSayi {

		if tahminEdilenSayi > 100 || tahminEdilenSayi < 0 {
			fmt.Println("Lütfen 100 ile 0 arasında bir değer giriniz")
			fmt.Scanln(&tahminEdilenSayi)
		}

		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}

	}

	basariDurumu := "" // empty

	if tahminSayisi > 0 && tahminSayisi <= 3 { // && -> Doğru olması || (pipe-boru) -> Ya da
		basariDurumu = "Süper"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena Değil"
	}

	fmt.Printf("Bravo Bildiniz. %v tahmin : %v", tahminSayisi, basariDurumu)

}

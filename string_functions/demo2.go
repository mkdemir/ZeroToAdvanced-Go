package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "mkdemir"
	fmt.Println(s.HasPrefix(isim, "mkd")) // Ön ek (Sonuç true veya false döndürüyor).
	fmt.Println(s.HasSuffix(isim, "mir")) // Son ek

	fmt.Println(s.Index(isim, "em")) // em arar.
	fmt.Println(s.Index(isim, "a"))  // em arar.

	harfler := []string{"m", "k", "d", "e", "m", "i", "r"}

	fmt.Println(s.Join(harfler, "")) // Elementler string ( Tek bir string yaptı )
	fmt.Println(s.Join(harfler, "-"))

	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	// Replace ile değiştirme
	fmt.Println(s.Replace(sonuc, "*", "+", -1)) // -1 Tüm hepisi değiştir demek

	// TC: 20943865315,02022022,50000
	fmt.Println(s.Split(sonuc, "*"))

	fmt.Println(s.Repeat(sonuc, 5)) // 5 tane kopyasını gönderir.
}

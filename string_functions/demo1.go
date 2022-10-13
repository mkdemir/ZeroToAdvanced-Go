package string_functions

// importlarda alias dediğimiz yapıları kullanabiliriz.
import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Mustafa"
	fmt.Println(s.Count(isim, "a")) // ismin içinde kaç tane "a" var say bana demek. (strings paketi ile geliyor). (Case Sensitive) // ascii
	fmt.Println(s.Contains(isim, "a"))
	sonuc := s.Index(isim, "a") // Bulamazsa -1 dönderecek.
	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim)) // Metni büyük harfe çevir
	fmt.Println(s.ToUpper(isim)) // Metni büyük harfe çevir
}

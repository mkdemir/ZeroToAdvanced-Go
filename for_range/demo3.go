package for_range

import "fmt"

func Demo3() {
	//map
	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for k := range sozluk {
		fmt.Println(k) // İlk yazılan şey her zaman key
	}

	for k, v := range sozluk {
		fmt.Println(k) // İlk yazılan şey her zaman key'dir.
		fmt.Println(v)
	}
}

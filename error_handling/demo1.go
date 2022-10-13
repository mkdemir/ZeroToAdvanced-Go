package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo12.txt") // Dosya açma (2 tane dönüşü vardır. 1 Tanesi dosyanın kendisini verir.) Yoksa bize bir hata dönüyor.
	// nil -> Dosya bulunursa err o anki durumu nil'dir. Yani error'un oldmadğı 0 durumu

	// if err != nil { // Hata vardır. (Hata varsa)
	// 	fmt.Println("Dosya Bulunamadı")
	// } else {
	// 	fmt.Println(f.Name())
	// }

	if err != nil { // nil -> Hata olmadığı anlamına geliyor. Bir hata yoksa denir. Burada ise hata varsa diyoruz.

		if pErr, ok := err.(*os.PathError); ok { // err türü PathError mı
			fmt.Println("Dosya Bulunmadı : ", pErr.Path)
			return
			//err. -> Bir interface metodu, (*os.PathError) -> Type assertion ()
		} else {
			fmt.Println("Dosya Bulunamadı")
			return
		}

	} else {
		fmt.Println(f.Name()) // Hata yoksa burası çalışacak.
	}
}

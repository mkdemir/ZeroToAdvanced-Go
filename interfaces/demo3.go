package interfaces

import "fmt"

func dogrula(i interface{}) { // Bizim gönderdiğimiz herşeyi karşılıyor.
	sayi, ok := i.(int) // Tip dönüşü yapma | Type assertion
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = "mkdemir" // Yukarıdaki sayi ile burada sayi ayrı şeyler. (int ve string bir interface'dir)
	dogrula(5)
	dogrula(sayi1)

	var sayi2 interface{} = 5 // Yukarıdaki sayi ile burada sayi ayrı şeyler. (int ve string bir interface'dir)
	dogrula(sayi2)

}

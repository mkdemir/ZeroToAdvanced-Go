package main1 // Go packagelar ile çalışıyor yani dosyada package olmak zorunda.

import (
	"fmt"
)

func main1() { // Küçük harfle başlayan private, büyük harfle başlayan public (Dolayısıyla private olanı sadece main package kullanabilirken, public olanı diğer packagelar da kullanabilecekt) (encapsulation)

	// fmt.Println("Hello Security!")

	// go run deyincede derliyor sadece temp'in altında geçici bir yere derliyor. Orada çalıştıryor.
	// Primitive Type -> int, float, string vs.
	// var i int
	// i = 42
	// println(i)

	// var f float32 = 3.14
	// println(f)
	// name := "mkdemir"
	// println(name)

	// secondname := "Demir"
	// number := 4.32

	// fmt.Println(name, secondname, number)

	// var lastName *string = new(string) // pointer bellekteki bir adresi söylüyor.
	// *lastName = "Demir"
	// fmt.Println(lastName)
	// fmt.Println(*lastName)

	// var lastName *string = &name
	// // *lastName = "Demir"
	// fmt.Println(*lastName, lastName, &name)

	// test_name := "mustafa"
	// var test_name1 *string = new(string)
	// fmt.Println(test_name, test_name1)

	// Örnek 1: -------------------------
	// var name string
	// name = "Mustafa"
	// secondname := "Kaan"
	// age := 3
	// fmt.Println(name, secondname, age)
	// ----------------------------------

	// name := "mkdemir"

	// var pname *string = new(string) // pointer name pointer adresi görüntülendi

	// *pname = "mustafa"

	// fmt.Println(name, *pname)

	// name := "mkdemir"

	// fmt.Println(name)

	// var pname *string = &name

	// fmt.Println(*pname)

	// *pname = "demir"

	// fmt.Println(name, *pname)

	// const c int = 3 // Derleme zamanında bir kez kullanılıyor.

	// //c = 4 // Kullanamazsınız.
	// println(c)

	// println(monday, tday, wday)

	var arr [3]int // value type default 0 değeri alır.

	arr[0] = 1
	arr[1] = 2
	arr[2] = 4 // fix size, dinamik değil

	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr, arr2)

	slice := arr2[:] // array yok slice

	slice2 := arr2[1:]
	fmt.Println(slice) // Ekledikçe kendini büyüten bilgi
	fmt.Println(slice2)

	m := map[string]int{"mkdemir": 42}
	fmt.Println(m["mkdemir"])

	// Örnek 2

	// Fixed Lenght Slice
	var isimler [3]string

	isimler[0] = "mustafa"
	isimler[1] = "kaan"
	isimler[2] = "demir"

	fmt.Println(isimler)
	fmt.Println(isimler[1 : len(isimler)-1])

	// Slice

	var gercekler []string

	gercekler = append(gercekler, "Çok")

	gercekler = append(gercekler, "acidir") // append slice kopyalıyor.

	fmt.Println(gercekler)

	// Map

	harita := map[string]int{}

	harita["Ankara"] = 14
	harita["Ankara"] = 15

	fmt.Println(harita)

	type person struct {
		Id   int
		Name string
	}

	p := person{
		Id:   1,
		Name: "mkdemir",
	}

	fmt.Println(p)

}

const (
	monday = iota + 1
	tday   = 2 << monday // left shift
	wday
)

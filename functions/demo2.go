package functions

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) { // sayi1 int: Parametre
	toplama := sayi1 + sayi2
	cikarma := sayi1 - sayi2
	carpma := sayi1 * sayi2
	bolme := float32(sayi1 / sayi2)

	return toplama, cikarma, carpma, bolme
}

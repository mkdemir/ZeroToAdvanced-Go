/*
* Bu dosyayı hangi paketin içersine yerleştireceğimizi söylüyor olmamız. * package ve func main isimlerinin aynı olması zorunlu değildir. (Burada main isminde bir paket oluşturmuşsuz.)
* Bildiğimiz paket ifadesiyle benzetebilirsiniz.
* Neden bir şeyi paketleriz içerisine birbiri ile ilişkili şeyleri koyarız. Hediye paketi gibi düşünebilirsiniz. Bir kaç tane aldığımız hediyeyi o paketin içerisine koyarız. Ve o bizim için bir doğum günü paketidir.
* Birbiri ile ilişikili olan değerleri biz bu paketlerin içerisine koyarız. Ve onu daha sonra istediğimiz yerde kullanırız.

* fmt paketi bizim için format paketi bizim için aslında bir paket onun içerisine ilgili dosyaları koymuşlar ve onu çalıştırıyoruz.

* Bir paketi başka bir yerde kullanmak için import ifadesiyle onu sisteme eklemeniz gerekli

* spaghetti code: İlgisiz şeylerin aynı dosya içinde bulunması

* Module -> İçinde packa<ge bulunduran yapıdır.
 */
package variables

import "fmt" // Go'nun içerisinde olan google tarafından yazılmış bir paket var. (Paket içerisine ilgili dosyalarımızı koyacağız).

// Bir go kodunun çalışabilmesi için öncelikle main dediğimiz bir bloğa ihtiyacımız vardır. (Main bloğu vasıtasıyla go komutlarını çalıştırabiliyoruz.)
// camelCase -> ilkDemo
// PascalCase

// Bir go kodunun çalışabilmesi için öncelikle main dediğimiz bir bloğa ihtiyacımız vardır. (Main bloğu vasıtasıyla go komutlarını çalıştırabiliyoruz.)
func Demo1() { // go dilinde fonksiyonlar Pascal Case'dir. Kelimelerin ilk harfi büyüktür.
	fmt.Print("Hello Friend ") // Ekrana yazdırma
	// Kullanmadığınız bir şeyi bir çok dilin aksine bir değişken veya bir import yaptığınızda onu kullanmasanızda sıkıntı çıkarmaz. Ama go sıkıntı çıkarıyor.
	// Aslında main bir fonksiyon. Şuan yazmak zorundayız. Bu blok go kodumuzun çalışacağını bloğun ta kendisi
	fmt.Print("Hello Friend")
	fmt.Println("!")

	/*
		Değişken ve Veri tipleri:
	*/
	// String tipinde bir değişken tanımladık. String, metinsel ifadeleri tutmak için kullanılır.
	// var metin adlı keyword'tür. Keyword anahtar kelimedir. Go programlama dilinde ve birçok programlama dilinde lullanılan bir anahtar kelimedir.
	var metin string = "Merhaba Dünya! " // Birçok programlama dilinde ' veya " tırnak fark etmiyebiliyor ama go programlama dilinde biz metinleri " tırnak içine alıyoruz.
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Println(metin)

	// Integer -> Tam Sayı, overflow sınırları aşmakdır.
	// Sayısal bir veri tipi sayı anlamına geliyor. (Tamsayı Integer)
	var kdv int = -15
	fmt.Println(kdv) // ctrl + space: inbtellisense
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	// Float -> Ondalık
	// Tam sayı olmayan (numeric odalıklı tutabileceğimiz).
	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	// var zorunludur. Ama onun yerine kullanabileceğimiz diğer ifade ise
	kdv3 := 25.2 // := hem tanımlıyorum hemde atama yapıyorum demek.
	//kdv3 := "Mustafa" // Sen float64 tipindeki bir değere string değer atayamazsın diyor. Go güvenli bir dil olduğu içindir.
	fmt.Println(kdv3)
	// Değişkenin veri tipini öğrenmek için
	fmt.Printf("Veri Tipi: %T", kdv3) // Formatlı yazdırma işlemi gerçekleştirmek için % gördüğünüzde ilk % ,'den sonra kdv3 içine koymaya yarar. %T demek Type demekdir.

	/*
		Mantıksal Veri Tipleri:
			Mantıksal Durum: Programlamada belirli koşullar vardır. (False veya True'dur)
	*/

	// var durum bool
	var durum bool = true // false
	// durum := bool, var durum bool
	var metin1 string = "Mustafa"
	var metin2 string = "Mustafa"
	durum = metin1 == metin2        // false
	fmt.Println("\nDurum: ", durum) // Sonuç false = değer atama == eşit eşit mi (== eşit eşit operatörü diyoruz).

	durum = metin1 != metin2 // != eşit değil (farklı mı)
	fmt.Println("\nDurum: ", durum)
	fmt.Println("\nDurum: ", !durum) // ! ise mantıksal bir değeri önüne ! koyduğunuzda durumun tersini yazar. !'mi sadece bool'ın önüne koyabilirsiniz.s
}

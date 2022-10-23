# Go ile Programlamaya Giriş

![Gopher](./images/gopher.png)

Go bir programlama dilidir. (Mobil, masaüstü ve web geliştirme için altyapılar kurabilirsiniz).

Go basit, güvenilir ve verimli yazılımlar tasarlamamızı sağlayan açık kaynak bir yazılım dilidir. Google'daki bir takım ve birçok açık kaynak destekçisinin katılımıyla ortaya çıkmıştır.

Go'yu kurmak için [Go Kurulum](https://go.dev/dl/) tıkladığınızda uygun olan seçeneği seçtiğinizde otomatik olarak kurulum gerçekleşecektir.

## Kurulum

### Ubuntu

Go'yu ubuntuya kurmak için.

1. `curl -OL https://go.dev/dl/go1.19.2.linux-amd64.tar.gz` -> `Curl` kullanın ve vurgulanan URL'yi az önce kopyaladığınız URL ile değiştirdiğinizden emin olun. `-O` bayrağı bunun bir dosyaya çıktı vermesini sağlar ve `-L` bayrağı HTTPS yönlendirmelerini yönlendirir, çünkü bu bağlantı Go web sitesinden alınmıştır ve dosya indirilmeden önce buraya yönlendirilecektir.

2. `sha256sum go1.19.2.linux-amd64.tar.gz`İndirdiğiniz dosyanın bütünlüğünü doğrulamak için sha256sum komutunu çalıştırın ve dosya adını argüman olarak iletin.

3. `` Ardından, tar'lı dosyayı çıkarmak için `tar`'ı kullanın. Bu komut, tar'a başka herhangi bir işlem yapmadan önce verilen dizine geçmesini söyleyen -C bayrağını içerir. Bu, çıkarılan dosyaların Go'yu yüklemek için önerilen konum olan /usr/local/ dizinine yazılacağı anlamına gelir... x bayrağı tar'a çıkarmasını, v verbose çıktı (çıkarılan dosyaların bir listesi) istediğimizi ve f bir dosya adı belirteceğimizi söyler.
    * Go'yu yüklemek için önerilen konum /usr/local/go olmasına rağmen, bazı kullanıcılar farklı yolları tercih edebilir veya gerektirebilir.
4. `sudo nano ~/.profile` 

5. `export PATH=$PATH:/usr/local/go/bin`

6. `source ~/.profile`

7. `go version`

`go version`: Go versiyonunu öğrenme (Ayrıca go kurlumu değil mi kontrol edebilirsiniz).

`go run .\main.go`: Bu klasörün içerisindeki main.go dosyasını çalıştır.

`go mod init module_ismi`: Modül oluşturma, go module initialisation  => Module oluşturma

`import "golesson/variables"` -> Modül ismi, paket ismi => Anlamı ise ben main paketi içerisinde golesson modül altındaki variables paketini kullanmak istiyorum.

## JSON API'ler ile yani `Restful` servislerle çalışma

{JSON} Placeholder

Restful servislerden önce SOAP dediğimiz sistemler üzerinde yaplıyordu. Ondan öncesinde remoting üzerinden yapılıyordu

Restful -> Dağıtık sistemleri birbiri ile görüştürmek için restful servislerdir.
['jsonplaceholder'](https://jsonplaceholder.typicode.com/)

* JSON'da biz buna obje notasyonu diyoruz. Bir kayıtdır.

```json
{
    "userId": 1,
    "id": 1,
    "title": "delectus aut autem",
    "completed": false
  }
```

### ROUTES

* GET: Ben okumak istiyorum demek. (Sende bir de3ğişiklik yapmayacağım bana lütfen bir tane data ver). (Okumak)
* POST: Silme, güncelleme hemde data eklemek için kullanılır çoğunlukla
* PUT
* PATCH
* DELETE

[json-server](https://github.com/typicode/json-server)

* `npm install -g json-server` => Paket yönetiminden yükleme yap. `json-server` kur. -g (global) sen buna istediğin yolda yazarak çalıştırabilirsin. (C'den istediğin yerde çalışabilir).

* `json-server --watch db.json` eğer çalışmıyorsa `npx json-server --watch db.json`


# Go ile Programlamaya Giriş

![Gopher](./images/gopher.png)

Go bir programlama dilidir. (Mobil, masaüstü ve web geliştirme için altyapılar kurabilirsiniz).

Go'yu kurmak için [Go Kurulum](https://go.dev/dl/) tıkladığınızda uygun olan seçeneği seçtiğinizde otomatik olarak kurulum gerçekleşecektir.

`go version`: Go versiyonunu öğrenme (Ayrıca go kurlumu değil mi kontrol edebilirsiniz).

`go run .\main.go`: Bu klasörün içerisindeki main.go dosyasını çalıştır.

`go mod init module_ismi`: Modül oluşturma, go module initialisation  => Module oluşturma

`import "golesson/variables"` -> Modül ismi, paket ismi => Anlamı ise ben main paketi içerisinde golesson modül altındaki variables paketini kullanmak istiyorum.

## JSON API'ler ile yani Restful servislerle çalışma

{JSON} Placeholder

Restful servislerden önce SOAP dediğimiz sistemler üzerinde yaplıyordu. Ondan öncesinde remoting üzerinden yapılıyordu

Restful -> Dağıtık sistemleri birbiri ile görüştürmek için restful servislerdir.
!['jsonplaceholder'](https://jsonplaceholder.typicode.com/)

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

!['json-server'](https://github.com/typicode/json-server)

* `npm install -g json-server` => Paket yönetiminden yükleme yap. `json-server` kur. -g (global) sen buna istediğin yolda yazarak çalıştırabilirsin. (C'den istediğin yerde çalışabilir).

* `json-server --watch db.json` eğer çalışmıyorsa `npx json-server --watch db.json`


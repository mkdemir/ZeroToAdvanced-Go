package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct { // https://mholt.github.io/json-to-go/ api'nizdeki rest yapısını JSON'ını yani buraya yazarsınız go'daki karşılığını verecktir.
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	repsonse, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // Bize error dönüyor.

	if err != nil {
		fmt.Println(err)
	}
	defer repsonse.Body.Close() // Kanal açıklığı olursa performansı düşük uygulamalar olur.

	bodyBytes, _ := ioutil.ReadAll(repsonse.Body) // http response

	bodyString := string(bodyBytes) // Bize dönen data'nın string hali ama o data formatlı olarak geliyor.
	fmt.Println(bodyString)

	// String formatını json dönüştürmek için bir paket var.
	var todo Todo // Todo değişkinden oluşturma
	json.Unmarshal(bodyBytes, &todo)

	fmt.Println(todo) // {1 1 delectus aut autem false} todo struct'ına atadı
}

func Demo2() { // get operasyonu bir data okumak için kullanılıyordu post opersyonu ise aslında bir datayı post etme üzerine kurulu yani yollama üzerine kullanma
	todo := Todo{1, 2, "Alış veriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo)) // content type, ben sana şu formatta yolluyorum demek. Gönderilen data'nın hangi formatta olduğunu söylediği bilgidir.
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) // http response

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)

}

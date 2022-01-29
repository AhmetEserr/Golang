package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"UserId"` //alt gr + virgül(,) 2 kere basılacak // neyi netledik hangi alanı neye eşleştirecek
	Id        int    `json:"Id"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

// Json to go yazarak internete hazır struct yapılarını transfor eden web siteler ile de yapılabilir..!!
func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err) //hatayı kontrol et
	}
	defer response.Body.Close() //kapatma

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyStrings := string(bodyBytes)
	fmt.Println(bodyStrings)
	//stringe çevirdik.
	//json formatında stringe dönüştürmek için
	var todo Todo
	json.Unmarshal(bodyBytes, &todo) //nesne de görmek istersek
	fmt.Println(todo)
}

//post gönderi olarak yollamak için kullanılır.. veri kaynagına müdale etme yeni kayıt etme veriyi güncelleme veriye silme işlemleri için kullanılır.
func Demo2() {
	//önce bilgiyi toparlamak lazım..
	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	//jsona çevirilmesi gerekir..
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8", bytes.NewBuffer(jsonTodo)) //adrese kayıt ekleme olayına girer.
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyStrings := string(bodyBytes)
	fmt.Println(bodyStrings)
	//stringe çevirdik.
	//json formatında stringe dönüştürmek için
	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse) //nesne de görmek istersek
	fmt.Println(todoResponse)
	//{1 201 Alışveriş yapılacak false} şeklinde dönecektir.
}

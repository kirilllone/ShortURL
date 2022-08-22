package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

// Функция shorting генерирует короткую ссылку случайным образом
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Функция isValidUrl проверяет URL
func isValidUrl(token string) bool {
	_, err := url.ParseRequestURI(token)
	if err != nil {
		return false
	}
	u, err := url.Parse(token)
	if err != nil || u.Host == "" {
		return false
	}
	return true
}

// Map для хранения ссылок (ключ-сокращенная, значение-исходная)
var LinkList = make(map[string]string)

func main() {
	// две ветки по тз: первая - добавляет новую ссылку, вторая - по короткой выдает оригинал
	http.HandleFunc("/post", NewLinkHandler)
	http.HandleFunc("/", FindLinkHandler)
	log.Fatal(http.ListenAndServe(":8085", nil))

}

// поиск короткой ссылки в памяти
func FindLinkHandler(w http.ResponseWriter, req *http.Request) {
	// считываем тело запроса
	body, err := io.ReadAll(req.Body)
	if err != nil {
		// при ошибке - 500
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer req.Body.Close()
	// поиск ссылки в Map
	if OriginalLink, found := LinkList[string(body)]; found {
		w.Write([]byte(OriginalLink))
		// редирект на оригинальную страницу
		http.Redirect(w, req, OriginalLink, http.StatusTemporaryRedirect)

		w.WriteHeader(http.StatusTemporaryRedirect)

	} else {
		// если такой  ссылки  нет - 404
		w.WriteHeader(http.StatusNotFound)
	}
}

// добавление новой ссылки
func NewLinkHandler(w http.ResponseWriter, req *http.Request) {

	// считываем тело запроса
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer req.Body.Close()
	// проверка на URL
	if !isValidUrl(string(body)) {
		// при ошибке - 500
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		//генерируем короткую ссылку
		var Shortlink = shorting()
		//добавляем в Map
		LinkList[Shortlink] = string(body)
		//статус - 200
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(Shortlink))

	}
}

// примеры запросов через терминал

// curl -v -X POST -H "Content-Type: text/plain" -d "https://yandex.ru" 'localhost:8085/post'
//curl -v -X POST -H "Content-Type: text/plain" -d "https://google.com" 'localhost:8085/post'

// curl -v -H "Content-Type: text/plain" -d "XVlBz" 'localhost:8085'

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var args = os.Args[1:]

func handler(w http.ResponseWriter, r *http.Request) {
	// если наш запрос это запрос GET
	if r.Method == "GET" {
		// установка заголовка для индикации JSON ответа
		w.Header().Set("Content-Type", "application/json")

		// запись ответа
		err := json.NewEncoder(w).Encode(args)
		if err != nil {
			// запись ошибки 500 если кодировка неудачна
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	// если наш запрос это POST или PUT запрос
	if r.Method == "PUT" || r.Method == "POST" {
		// читаем body
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// запись ошибки 500 если ReadAll терпит неудачу)
			w.WriteHeader(http.StatusInternalServerError)
		}

		// раздели строки body на субстроки с помощю
		// одной или более whitespace(\s) как ограничитель
		pattern := regexp.MustCompile(`\s+`)
		newArgs := pattern.Split(string(reqBody), -1)

		// присобачить новые аргументы к уже существующим
		args = append(args, newArgs...)
		// возвращает 204
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	// handler для вивода gets, posts, puts
	http.HandleFunc("/", handler)
	// listen, server and fatally отключаются когда нужно
	log.Fatal(http.ListenAndServe(":8099", nil))
}






























//"код ваще не украденый а общедоступный!"
//В.И. Ленин)))
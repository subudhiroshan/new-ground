package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	storedValue = 0
)

func get(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, string(strconv.Itoa(storedValue)))
}

func increase(w http.ResponseWriter, _ *http.Request) {
	storedValue += 1
	io.WriteString(w, "increased\n")
}

func decrease(w http.ResponseWriter, _ *http.Request) {
	storedValue -= 1
	io.WriteString(w, "decreased\n")
}

func double(w http.ResponseWriter, _ *http.Request) {
	storedValue *= 2
	io.WriteString(w, "doubled\n")
}

func zero(w http.ResponseWriter, _ *http.Request) {
	storedValue = 0
	io.WriteString(w, "zeroed\n")
}

func main() {
	http.HandleFunc("/", get)
	http.HandleFunc("/increase", increase)
	http.HandleFunc("/decrease", decrease)
	http.HandleFunc("/double", double)
	http.HandleFunc("/zero", zero)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("page/"))
	http.Handle("/", http.StripPrefix("/", fs))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"net/http"
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/css"))))
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}

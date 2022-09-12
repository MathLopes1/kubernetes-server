package main

import "net/http"

func main() {
	http.HandleFunc("/", Helllo)
	http.ListenAndServe(":8080", nil)
}

func Helllo (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle</h1>"))
}
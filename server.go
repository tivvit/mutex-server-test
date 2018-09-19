package main

import (
	"net/http"
)


func simpleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", simpleHello)
	http.HandleFunc("/cnt", simpleHello)
	http.HandleFunc("/goroutines", simpleHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}



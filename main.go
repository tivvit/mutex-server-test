package main

import (
	"net/http"
	"sync"
)

var s sync.Mutex

func sayHello(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	w.Write([]byte("ok"))
	s.Unlock()
}

func main() {
	s = sync.Mutex{}
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}



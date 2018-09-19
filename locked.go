package main

import (
	"net/http"
	"sync"
	//"runtime"
	"fmt"
	"sync/atomic"
	"runtime"
)

var s sync.Mutex
var cnt int32

func lockHello(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	w.Write([]byte("ok"))
	s.Unlock()
}

func lockCntHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got request (goroutines %d)\n", runtime.NumGoroutine())
	s.Lock()
	fmt.Printf("Locked (goroutines %d)\n", runtime.NumGoroutine())
	w.Write([]byte("ok"))
	s.Unlock()
	fmt.Printf("Unocked\n")
}

func lockGoroutinesHello(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&cnt, 1)
	fmt.Printf("Waiting %d\n", atomic.LoadInt32(&cnt))
	s.Lock()
	w.Write([]byte("ok"))
	s.Unlock()
	atomic.AddInt32(&cnt, -1)
}

func main() {
	s = sync.Mutex{}
	cnt = 0
	http.HandleFunc("/", lockHello)
	http.HandleFunc("/cnt", lockCntHello)
	http.HandleFunc("/goroutines", lockGoroutinesHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}



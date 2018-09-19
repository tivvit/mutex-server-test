package main

import (
	"net/http"
	"sync"
	"time"
	//"fmt"
)

var l sync.Mutex
var queue []string

func lockQHello(w http.ResponseWriter, r *http.Request) {
	l.Lock()
	w.Write([]byte("ok"))
	queue = append(queue, "")
	//fmt.Printf("Queue len: %d cap: %d\n", len(queue), cap(queue))
	l.Unlock()
}


func main() {
	l = sync.Mutex{}

	go func() {
		for {
			//fmt.Printf("XX %d\n", len(queue))
			for len(queue) > 0 {
				queue = queue[1:]
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	http.HandleFunc("/", lockQHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}



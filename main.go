package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {

		now := strconv.Itoa(int(time.Now().Unix()))
		w.Write([]byte(now))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

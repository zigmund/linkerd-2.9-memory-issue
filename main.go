package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	addr       = ":8080"
	minDelayMS = 25
	maxDelayMS = 100
)

func main() {
	http.HandleFunc("/fast", fast)
	http.HandleFunc("/slow", slow)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	log.Printf("server started on %s", addr)
	select {}
}

func fast(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Write([]byte("ok"))
}

func slow(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 128Kb reply
	b := randBytes(128 * 1024)

	// Delay request for for minDelay...maxDelay ms
	delay := time.Duration(minDelayMS+rand.Intn(maxDelayMS-minDelayMS)) * time.Millisecond
	time.Sleep(delay)

	_, err := w.Write(b)

	if err != nil {
		log.Print(err)
	}
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}

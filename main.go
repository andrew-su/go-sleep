package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	durStr, ok := r.URL.Query()["duration"]
	if ok && len(durStr) > 0 {
		dur, err := strconv.Atoi(durStr[0])
		if err != nil {
			log.Printf("Received invalid duration: %v", durStr)
			return
		}
		log.Printf("Sleeping for %v", time.Duration(dur) * time.Second)
		time.Sleep(time.Duration(dur) * time.Second)
	}
	log.Print("Waking up")
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

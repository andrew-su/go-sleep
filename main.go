package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	durStr, ok := r.URL.Query()["duration"]
	if !ok {
		log.Printf("Did not find parameter duration")
		w.Write([]byte("Duration not specified"))
	}

	var sleepDuration time.Duration
	dur, err := strconv.Atoi(durStr[0])
	if err != nil {
		log.Printf("Received invalid duration: %v", durStr)
		return
	}
	sleepDuration = time.Duration(dur) * time.Second
	log.Printf("Sleeping for %v", sleepDuration)
	time.Sleep(sleepDuration)

	w.Write([]byte(fmt.Sprintf("Slept for %v", sleepDuration)))
	log.Print("Waking up")
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

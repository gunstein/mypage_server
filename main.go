package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./gunstein_vatnar_no"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

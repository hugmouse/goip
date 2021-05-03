package main

import (
	"log"
	"net/http"
)

func main() {
  // this code is broken right now
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
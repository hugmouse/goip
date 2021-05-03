package main

import (
	"github.com/hugmouse/goip/api/getInfo"
	"log"
	"net/http"
)

func main() {
  // this code is broken right now
	http.HandleFunc("/", handler.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
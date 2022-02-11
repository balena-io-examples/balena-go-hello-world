package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./views")))
	addr := ":80"
	fmt.Println("Listening on port", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("p", 8080, "Web server port")

func main() {
	flag.Parse()
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Printf("Listening on :%d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

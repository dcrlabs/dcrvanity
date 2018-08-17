package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8000", "port")
	flag.Parse()
	listenHost := "127.0.0.1:" + *port
	log.Printf("HTTP server listening on http://%s\n", listenHost)
	log.Fatal(http.ListenAndServe(listenHost, http.FileServer(http.Dir("."))))
}

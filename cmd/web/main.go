package main 

import (
	"net/http"
	"log"
	"flag"
)

func main() {
	mux := http.NewServeMux()
	
	fileserver := http.FileServer(http.Dir("./ui/static"))
	addr := flag.String("addr", ":4000", "Port to run serve mux on")
	flag.Parse()
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("GET /", home)
	log.Printf("Starting server on port %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

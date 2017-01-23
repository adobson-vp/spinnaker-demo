package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	bufstream, err := ioutil.ReadFile("app/index.html")

	if err != nil {
		log.Fatal("Error during application running: %v", err)
		os.Exit(1)
	}

	io.WriteString(w, string(bufstream[:]))
	log.Println("Req: ", r.Host, r.URL.Path, r.RemoteAddr)
}

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

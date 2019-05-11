package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/cmoorebytes/wordoftheday/pkg/router"
)

func main(){
	log.Println("Initializing the service...")
	handler := http.NewServeMux()
	log.Println("Registering API endpoints...")
	handler.HandleFunc("/api/word/wotd", router.GetWordHandler)
	handler.HandleFunc("/api/word/scramble", router.ScrambleHandler)
	handler.HandleFunc("/api/word/puzzle",router.PuzzleHandler)
	handler.HandleFunc("/test", router.TestHandler)
	log.Println("Registering static site file handler...")
	handler.HandleFunc("/", router.DefaultHandler)
	log.Println("Getting HTTP server port number...")
	port_number := router.PortNumber()
	log.Println(fmt.Sprintf("Starting server. Listening on port %s", port_number))
	log.Fatal(http.ListenAndServe(port_number, handler))
}
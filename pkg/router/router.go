package router

import (
	"net/http"
	"fmt"
	"os"
)

const default_port_number = 8000

func PortNumber() string{
	port, ok := os.LookupEnv("PORT")
	if(!ok){
		return fmt.Sprintf(":%d", default_port_number)
	} else {
		return fmt.Sprintf(":%s", port)
	}
	
}

func Register(handler *http.ServeMux){
	handler.HandleFunc("/api/word/wotd", GetWordHandler)
	handler.HandleFunc("/api/word/scramble", ScrambleHandler)
	handler.HandleFunc("/api/word/puzzle", PuzzleHandler)
	handler.HandleFunc("/test", TestHandler)
	handler.HandleFunc("/", DefaultHandler)
}
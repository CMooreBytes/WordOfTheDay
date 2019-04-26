package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var svc WordServiceInterface

func testHandler(w http.ResponseWriter, r *http.Request) {
	result := new(Result)
	result.Word = "word"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func handler(w http.ResponseWriter, r *http.Request) {
	svc := new (WordService)
	result := svc.GetWord()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main(){
	log.Println("Starting up")
	http.HandleFunc("/api/word", handler)
	http.HandleFunc("/test/word", testHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
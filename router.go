package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"html/template"
	"log"
)

const default_port_number = 8000

type AppRouter struct { }

func (AppRouter) PortNumber() string{
	return fmt.Sprintf(":%d", default_port_number);
}

func (router AppRouter) Register(handler *http.ServeMux){
	handler.HandleFunc("/", router.AppHandler)
	handler.HandleFunc("/api/word/wotd", router.GetWordHandler)
	handler.HandleFunc("/api/word/scramble", router.ScrambleHandler)
	handler.HandleFunc("/api/word/puzzle", router.PuzzleHandler)
	handler.HandleFunc("/test", router.TestHandler)
}

func (router AppRouter) AppHandler(w http.ResponseWriter, r *http.Request) {
	t,err := template.New("index.html").ParseFiles("index.html")
	if(err != nil){
		log.Fatal(err)
	}
	svc := new (WordService)
	result := new(Result)
	result.Word, result.ScrambledWord = svc.GetScrambledWord()
	t.Execute(w, result)
}

func (router AppRouter) TestHandler(w http.ResponseWriter, r *http.Request) {
	svc := new (WordService)
	word := "word"
	scrambled_word := svc.Scramble(word)
	processResponse(w, word, scrambled_word)
}

func (router AppRouter) GetWordHandler(w http.ResponseWriter, r *http.Request) {
	svc := new (WordService)
	word, scrambled_word := svc.GetScrambledWord()
	processResponse(w, word, scrambled_word)
}

func (router AppRouter) ScrambleHandler(w http.ResponseWriter, r *http.Request) {
	query_string := r.URL.Query()
	word := query_string.Get("word")
	svc := new (WordService)
	scrambled_word := svc.Scramble(word)
	processResponse(w, word, scrambled_word)
}

func (router AppRouter) PuzzleHandler(w http.ResponseWriter, r *http.Request) {
	svc := new (WordService)
	_, scrambled_word := svc.GetScrambledWord()
	processResponse(w, "", scrambled_word)
}

func processResponse(w http.ResponseWriter, word string, scrambled_word string) {
	result := new(Result)
	result.Word = word;
	result.ScrambledWord = scrambled_word
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
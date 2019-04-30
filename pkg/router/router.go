package router

import (
	"encoding/json"
	"net/http"
	"fmt"
	"html/template"
	"log"
	"github.com/cmoorebytes/wordoftheday/pkg/wordservice"
)

const default_port_number = 8000

func PortNumber() string{
	return fmt.Sprintf(":%d", default_port_number);
}

func Register(handler *http.ServeMux){
	handler.HandleFunc("/", appHandler)
	handler.HandleFunc("/api/word/wotd", getWordHandler)
	handler.HandleFunc("/api/word/scramble", scrambleHandler)
	handler.HandleFunc("/api/word/puzzle", puzzleHandler)
	handler.HandleFunc("/test", testHandler)
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	t,err := template.New("index.html").ParseFiles("index.html")
	if(err != nil){
		log.Fatal(err)
	}
	
	result := new(Result)
	result.Word, result.ScrambledWord = wordservice.GetScrambledWord()
	t.Execute(w, result)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	word := "word"
	scrambled_word := wordservice.Scramble(word)
	result := new(Result)
	result.Word = word;
	result.ScrambledWord = scrambled_word
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getWordHandler(w http.ResponseWriter, r *http.Request) {
	word, scrambled_word := wordservice.GetScrambledWord()
	result := new(Result)
	result.Word = word;
	result.ScrambledWord = scrambled_word
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func scrambleHandler(w http.ResponseWriter, r *http.Request) {
	query_string := r.URL.Query()
	word := query_string.Get("word")
	scrambled_word := wordservice.Scramble(word)
	result := new(Result)
	result.Word = word;
	result.ScrambledWord = scrambled_word
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func puzzleHandler(w http.ResponseWriter, r *http.Request) {
	_, scrambled_word := wordservice.GetScrambledWord()
	result := new(PuzzleResult)
	result.ScrambledWord = scrambled_word
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
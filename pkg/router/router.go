package router

import (
	"net/http"
	"fmt"
	"os"
	"html/template"
	"log"
	"path"
	"encoding/json"
	"github.com/cmoorebytes/wordoftheday/pkg/wordservice"
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

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	base := path.Base(r.URL.Path)
	if(base == "/") {
		base = "index.html"
	}
	switch path := path.Ext(r.URL.Path); path {
		case ".js":
			http.ServeFile(w, r, "wwwroot/js/" + base)
		case ".css":
			http.ServeFile(w, r, "wwwroot/css/" + base)
		default:
			t,err := template.New(base).ParseFiles("wwwroot/" + base)
			if(err != nil){
				log.Fatal(err)
			}
			
			result := new(Result)
			result.Word, result.ScrambledWord = wordservice.GetScrambledWord()
			t.Execute(w, result)
	}
}

func GetWordHandler (w http.ResponseWriter, r *http.Request) {
	result := new(Result)
	result.Word, result.ScrambledWord = wordservice.GetScrambledWord()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func PuzzleHandler(w http.ResponseWriter, r *http.Request) {
	_, scrambled_word := wordservice.GetScrambledWord()
	result := PuzzleResult { ScrambledWord: scrambled_word }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ScrambleHandler(w http.ResponseWriter, r *http.Request) {
	query_string := r.URL.Query()
	word := query_string.Get("word")
	result := Result {Word: word, ScrambledWord: wordservice.Scramble(word) }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}


func TestHandler(w http.ResponseWriter, r *http.Request) {
	word := "word"
	result := Result { Word: word, ScrambledWord: wordservice.Scramble(word) }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

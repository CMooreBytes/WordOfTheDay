package router

import (
	"net/http"
	"encoding/json"
	"github.com/cmoorebytes/wordoftheday/pkg/wordservice"
)

func ScrambleHandler(w http.ResponseWriter, r *http.Request) {
	query_string := r.URL.Query()
	word := query_string.Get("word")
	result := Result {Word: word, ScrambledWord: wordservice.Scramble(word) }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
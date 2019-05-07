package router

import (
	"net/http"
	"encoding/json"
	"github.com/cmoorebytes/wordoftheday/pkg/wordservice"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	word := "word"
	result := Result { Word: word, ScrambledWord: wordservice.Scramble(word) }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

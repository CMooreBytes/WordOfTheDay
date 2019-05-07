package router

import (
	"net/http"
	"encoding/json"
	"github.com/cmoorebytes/wordoftheday/pkg/wordservice"
)

func PuzzleHandler(w http.ResponseWriter, r *http.Request) {
	_, scrambled_word := wordservice.GetScrambledWord()
	result := PuzzleResult { ScrambledWord: scrambled_word }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

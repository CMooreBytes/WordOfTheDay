package router

import (
	"net/http"
	"encoding/json"
	"github.com/cmoorebytes/wordoftheday/pkg/wordservice"
)

func GetWordHandler(w http.ResponseWriter, r *http.Request) {
	result := new(Result)
	result.Word, result.ScrambledWord = wordservice.GetScrambledWord()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
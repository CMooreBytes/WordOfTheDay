package router

import (
	"net/http"
	"html/template"
	"log"
	"path"
	"github.com/cmoorebytes/wordoftheday/pkg/wordservice"
)

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
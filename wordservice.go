package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"math/rand"
	"fmt"
)

const svc_endpoint_base = "https://en.wiktionary.org/w/api.php?action=parse&format=json&prop=links&page=Wiktionary:Word_of_the_day/"

type WordService struct {
	ServiceEndpointBase string
}

type WordServiceInterface interface{
	GetUrl(t time.Time) string
	GetWord() string
	GetScrambledWord() (string,string)
	Scramble(word string) string
}

func (w WordService) GetUrl(t time.Time) string {
	tf := t.Format("January__2")
	return fmt.Sprintf("%s%s", svc_endpoint_base, tf)
}

func (w WordService) GetWord() string {
	svc_endpoint := w.GetUrl(time.Now())
	resp, err := http.Get(svc_endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseBody := new(ServiceResponse)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseBody)
	if err != nil {
		log.Fatal(err)
	}
	
	return responseBody.Parse.Links[0].Value
}

func  (w WordService) GetScrambledWord() (string,string) {
	word := w.GetWord();
	scrambled_word := w.Scramble(word)
	return word, scrambled_word
}

func (w WordService) Scramble(word string) string{
	char_arr := scramble([]rune(word), 0)
	return string(char_arr)
}

func scramble(word []rune, index int) []rune {
	if index >= len(word) {	
		return word
	} else {	
		rand.Seed(time.Now().UnixNano())
		cursor := rand.Intn(len(word)-index)
		letter := word[cursor + index]		
		word[cursor + index] = word[index]		
		word[index] = letter		
		return scramble(word, index + 1)
	}
}
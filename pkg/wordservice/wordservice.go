package wordservice

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"math/rand"
	"fmt"
)

const svc_endpoint_base = "https://en.wiktionary.org/w/api.php?action=parse&format=json&prop=links&page=Wiktionary:Word_of_the_day/"

func GetUrl(t time.Time) string {
	month := t.Format("January")
	date := t.Format("2")
	return fmt.Sprintf("%s%s_%s", svc_endpoint_base, month, date)
}

func GetWord() string {
	svc_endpoint := GetUrl(time.Now())
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

func  GetScrambledWord() (string,string) {
	word := GetWord();
	scrambled_word := Scramble(word)
	return word, scrambled_word
}

func Scramble(word string) string{
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
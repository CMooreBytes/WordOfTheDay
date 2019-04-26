package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const svc_endpoint_base = "https://en.wiktionary.org/w/api.php?action=parse&format=json&prop=links&page=Wiktionary:Word_of_the_day/"

type ServiceResponse struct {
	Parse ParseResult  `json:"parse"`
}

type ParseResult struct {
	Title  string `json:"title"`
	Pageid int `json:"pageid"`
	Links  []LinkResult `json:"links"`
}

type LinkResult struct {
	Ns     int `json:"ns"`
	Exists string `json:"exists"`
	Value string `json:"*"`
}

type Result struct{
	Word string
}

type WordServiceInterface interface{
	GetWord() string
}

type WordService struct {

}

func (w WordService) GetWord() *Result {
	tf := time.Now().Format("January__2")
	svc_endpoint := svc_endpoint_base + tf
	
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
	result := new(Result)
	result.Word = responseBody.Parse.Links[0].Value
	return result
}
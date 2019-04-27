package main

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
	ScrambledWord string
}
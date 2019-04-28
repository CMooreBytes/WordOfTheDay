package main

import (
	"log"
	"net/http"
	"fmt"
)

func main(){
	log.Println("Initializing the router...")
	router := new(AppRouter)
	handler := http.NewServeMux()
	log.Println("Registering API endpoints...")
	router.Register(handler)
	port_number := router.PortNumber()
	log.Println(fmt.Sprintf("Starting server. Listening on port %s", port_number))
	log.Fatal(http.ListenAndServe(port_number, handler))
}